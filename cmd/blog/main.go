package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/arthurweinmann/blog/internal/config"
	"github.com/arthurweinmann/go-https-hug/pkg/acme"
	"github.com/arthurweinmann/go-https-hug/pkg/router"
	"github.com/arthurweinmann/go-https-hug/pkg/storage/stores/filesystem"
)

func main() {
	boolArgs, stringargs := parseArgs()

	for argname := range boolArgs {
		switch argname {
		case "help":
			fmt.Println(`You may use:
			--home to set the home directory
			--webdomain to set the domain serving the public website
			--contactemail to set the contact email for the creation of let's encrypt certificates
			`)
			return
		default:
			log.Fatalf("unrecognized bool argument %s", argname)
		}
	}

	for argname, argval := range stringargs {
		switch argname {
		default:
			log.Fatalf("unrecognized string argument %s", argname)
		case "home":
			config.HOME = argval
		case "webdomain":
			config.PublicWebsiteDomain = argval
		case "contactemail":
			config.CertificateContactEmail = argval
		}
	}

	if config.HOME == "" || config.PublicWebsiteDomain == "" || config.CertificateContactEmail == "" {
		log.Fatal("command line arguments --home, --contactemail and --webdomain are mandatory")
	}

	router, err := router.NewRouter(&router.RouterConfig{
		ServeHTMLFolder:       filepath.Join(config.HOME, "web"),
		HTMLFolderDomainNames: []string{"arthurweinmann.com", "www.arthurweinmann.com"},
		RedirectHTTP2HTTPS:    true,
		PerDomain:             map[string]func(r *router.Router, spath []string, w http.ResponseWriter, req *http.Request){},
	})
	if err != nil {
		panic(err)
	}

	store, err := filesystem.NewStore(filepath.Join(config.HOME, "storage"))
	if err != nil {
		panic(err)
	}

	err = acme.Init(&acme.InitParameters{
		InMemoryCacheSize:       32 * 1024 * 1024,
		CertificateContactEmail: config.CertificateContactEmail,
		Store:                   store,
		AuthorizedDomains: map[string]map[string]bool{
			config.PublicWebsiteDomain: {
				"www." + config.PublicWebsiteDomain: true,
			},
		},
		DNSProvider: nil,
		LogLevel:    acme.DEBUG,
		Logger:      os.Stdout,
	})
	if err != nil {
		panic(err)
	}

	go acme.ServeHTTP(nil, true)

	err = acme.ToggleCertificate([]string{config.PublicWebsiteDomain, "www." + config.PublicWebsiteDomain})
	if err != nil {
		panic(err)
	}

	err = acme.ServeHTTPS(":443", router, filepath.Join(config.HOME, "https.log"))
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}

func parseArgs() (map[string]bool, map[string]string) {
	boolArgs := make(map[string]bool)
	strArgs := make(map[string]string)

	for i := 1; i < len(os.Args); i++ {
		if strings.HasPrefix(os.Args[i], "--") {
			arg := strings.TrimPrefix(os.Args[i], "--")
			if i+1 < len(os.Args) && !strings.HasPrefix(os.Args[i+1], "--") {
				strArgs[arg] = os.Args[i+1]
				i++ // skip next arg
			} else {
				boolArgs[arg] = true
			}
		}
	}

	return boolArgs, strArgs
}
