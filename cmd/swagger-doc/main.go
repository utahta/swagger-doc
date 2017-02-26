package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/pkg/browser"
	"github.com/utahta/swagger-doc"
)

type Options struct {
	Spec string `short:"s" long:"spec" description:"Specify the swagger.json" required:"true"`
	Port int    `short:"p" long:"port" default:"3330" description:"port number"`
}

func _main() error {
	opts := &Options{}
	if _, err := flags.Parse(opts); err != nil {
		return err
	}

	baseURL := fmt.Sprintf("http://localhost:%d", opts.Port)
	specURL := opts.Spec
	var specHandler http.HandlerFunc
	if !strings.HasPrefix(specURL, "http://") && !strings.HasPrefix(specURL, "https://") {
		specHandler = func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, opts.Spec)
		}
		specURL = fmt.Sprintf("%s/swagger.json", baseURL)
	}

	var errChan chan error
	go func() {
		if specHandler != nil {
			http.Handle("/swagger.json", specHandler)
		}
		http.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", doc.NewSwaggerUIHandler()))
		http.Handle("/redoc", doc.NewRedocHandler(specURL, "redoc"))

		bind := fmt.Sprintf(":%d", opts.Port)
		errChan <- http.ListenAndServe(bind, nil)
	}()

	browser.OpenURL(fmt.Sprintf("%s/swagger-ui/?url=%s", baseURL, specURL))
	return <-errChan
}

func main() {
	if err := _main(); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}
}
