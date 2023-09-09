package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/formatters/html"
	"github.com/alecthomas/chroma/lexers"
	"github.com/alecthomas/chroma/styles"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	mdhtml "github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

var (
	htmlFormatter  *html.Formatter
	highlightStyle *chroma.Style
	highlightCSS   string
)

func init() {
	htmlFormatter = html.New(html.WithClasses(true), html.TabWidth(2))
	if htmlFormatter == nil {
		panic("couldn't create html formatter")
	}
	styleName := "monokailight"
	highlightStyle = styles.Get(styleName)
	if highlightStyle == nil {
		panic(fmt.Sprintf("didn't find style '%s'", styleName))
	}
	b := bytes.NewBuffer(nil)
	err := htmlFormatter.WriteCSS(b, highlightStyle)
	if err != nil {
		panic(err)
	}
	highlightCSS = b.String()
}

type Article struct {
	HighlightCSS    string    `json:"HighlightCSS"`
	Title           string    `json:"Title"`
	Link            string    `json:"Link"`
	Content         string    `json:"Content"`
	LastEdit        time.Time `json:"LastEdit"`
	LastEditPretty  string    `json:"LastEditPretty"`
	MetaDescription string    `json:"MetaDescription"`
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.RemoveAll(filepath.Join(wd, "build/web"))
	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(filepath.Join(wd, "build/web/articles"), 0700)
	if err != nil {
		panic(err)
	}

	b, err := os.ReadFile(filepath.Join(wd, "web/article.html"))
	if err != nil {
		panic(err)
	}

	articleSkeleton, err := template.New("articleskeleton").Parse(string(b))
	if err != nil {
		panic(err)
	}

	var articles []any

	err = filepath.Walk(filepath.Join(wd, "web/articles"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			lastEdit, err := getFirstGitModificationTime(path)
			if err != nil {
				return err
			}

			art := &Article{
				HighlightCSS:   "<style>" + highlightCSS + "</style>",
				LastEdit:       lastEdit,
				LastEditPretty: PrettyDate(lastEdit),
			}

			fname := strings.Replace(filepath.Base(path), filepath.Ext(filepath.Base(path)), "", -1)

			art.Link = fmt.Sprintf("/articles/%s.html", fname)

			art.Title = strings.ReplaceAll(fname, "_", " ")
			art.Title = strings.Title(art.Title)

			b, err = os.ReadFile(path)
			if err != nil {
				return err
			}

			bstr := string(b)

			spl := strings.Split(bstr, "--END METAS--")
			if len(spl) > 1 {
				bstr = spl[1]

				scanner := bufio.NewScanner(strings.NewReader(spl[0]))
				for scanner.Scan() {
					line := scanner.Text()
					if strings.HasPrefix(line, "+++") {
						spl2 := strings.Split(line[3:], " ")
						linename := spl2[0]
						linetext := strings.Join(spl2[1:], " ")
						switch linename {
						case "MetaDescription":
							if len(linetext) > 140 {
								art.MetaDescription = linetext[:140] + "..."
							} else {
								art.MetaDescription = linetext
							}
						}
					}
				}

				if err := scanner.Err(); err != nil && err != io.EOF {
					panic(err)
				}
			}

			art.Content = string(mdToHTML([]byte(bstr)))

			f, err := os.Create(filepath.Join(wd, "build/web/articles/"+fname+".html"))
			if err != nil {
				return err
			}
			defer f.Close()

			b, err = json.Marshal(art)
			if err != nil {
				return err
			}

			data := map[string]any{}

			err = json.Unmarshal(b, &data)
			if err != nil {
				return err
			}

			err = articleSkeleton.Execute(f, data)
			if err != nil {
				return err
			}

			articles = append(articles, art)
		}

		return nil
	})
	if err != nil {
		panic(err)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].(*Article).LastEdit.After(articles[j].(*Article).LastEdit)
	})

	b, err = os.ReadFile(filepath.Join(wd, "web/index.html"))
	if err != nil {
		panic(err)
	}

	indexSkeleton, err := template.New("indexSkeleton").Parse(string(b))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join(wd, "build/web/index.html"))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	b, err = json.Marshal(&struct {
		Articles []any `json:"Articles"`
	}{
		Articles: articles,
	})
	if err != nil {
		panic(err)
	}

	data := map[string]any{}

	err = json.Unmarshal(b, &data)
	if err != nil {
		panic(err)
	}

	err = indexSkeleton.Execute(f, data)
	if err != nil {
		panic(err)
	}
}

func mdToHTML(md []byte) []byte {
	// create markdown parser with extensions
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	// create HTML renderer with extensions
	htmlFlags := mdhtml.CommonFlags | mdhtml.HrefTargetBlank
	opts := mdhtml.RendererOptions{Flags: htmlFlags, RenderNodeHook: myRenderHook}
	renderer := mdhtml.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}

func getFirstGitModificationTime(filePath string) (time.Time, error) {
	// Using a pipeline of git log to retrieve all commits and then tail to get the last (oldest) commit.
	cmd := exec.Command("sh", "-c", fmt.Sprintf("git log --pretty=format:%%cI %s | tail -n 1", filePath))
	output, err := cmd.Output()
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to execute command: %v", err)
	}

	// Remove leading and trailing white space
	outputStr := strings.TrimSpace(string(output))

	// Parse the time in RFC3339 format
	modTime, err := time.Parse(time.RFC3339, outputStr)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse time: %v", err)
	}

	return modTime, nil
}

func PrettyDate(t time.Time) string {
	return t.Format("Monday, 02-January-2006, 03:04 PM")
}

// based on https://github.com/alecthomas/chroma/blob/master/quick/quick.go
func htmlHighlight(w io.Writer, source, lang, defaultLang string) error {
	if lang == "" {
		lang = defaultLang
	}
	l := lexers.Get(lang)
	if l == nil {
		l = lexers.Analyse(source)
	}
	if l == nil {
		l = lexers.Fallback
	}
	l = chroma.Coalesce(l)

	it, err := l.Tokenise(nil, source)
	if err != nil {
		return err
	}
	return htmlFormatter.Format(w, highlightStyle, it)
}

// an actual rendering of Paragraph is more complicated
func renderCode(w io.Writer, codeBlock *ast.CodeBlock, entering bool) {
	defaultLang := ""
	lang := string(codeBlock.Info)
	htmlHighlight(w, string(codeBlock.Literal), lang, defaultLang)
}

func myRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
	if code, ok := node.(*ast.CodeBlock); ok {
		renderCode(w, code, entering)
		return ast.GoToNext, true
	}
	return ast.GoToNext, false
}
