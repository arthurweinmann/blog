package main

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

type Article struct {
	Title   string `json:"Title"`
	Link    string `json:"Link"`
	Content string `json:"Content"`
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
			art := &Article{}

			fname := strings.Replace(filepath.Base(path), filepath.Ext(filepath.Base(path)), "", -1)

			art.Link = fmt.Sprintf("/articles/%s.html", fname)

			art.Title = strings.ReplaceAll(fname, "_", " ")
			art.Title = strings.Title(art.Title)

			b, err = os.ReadFile(path)
			if err != nil {
				return err
			}

			art.Content = string(mdToHTML(b))

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

	b, err = os.ReadFile(filepath.Join(wd, "web/index.html"))
	if err != nil {
		panic(err)
	}

	indexSkeleton, err := template.New("indexSkeleton").Parse(string(b))
	if err != nil {
		panic(err)
	}

	f, err := os.Create(filepath.Join(wd, "build/web/articles/index.html"))
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
	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return markdown.Render(doc, renderer)
}
