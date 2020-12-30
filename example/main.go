// +build !wasm

package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/maxence-charriere/go-app/v7/pkg/app"
	"github.com/maxence-charriere/go-app/v7/pkg/cli"
	"github.com/maxence-charriere/go-app/v7/pkg/errors"
	"github.com/maxence-charriere/go-app/v7/pkg/logs"
)

type localOptions struct {
	Port int `cli:"p" env:"GO_UIKIT_PORT" help:"The port used by the server that serves the PWA."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	localOpts := localOptions{Port: 7777}
	cli.Register("local").
		Help(`Launches a server that serves the app in a local environment.`).
		Options(&localOpts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run the app on GitHub Pages.`).
		Options(&githubOpts)

	backgroundColor := "#ffffff"

	h := app.Handler{
		Author:          "Nivek Labs",
		BackgroundColor: backgroundColor,
		Description:     "A example application.",
		Keywords: []string{
			"go-app",
			"uikit",
			"pwa",
		},
		LoadingLabel: "Loading Go-UIkit...",
		Name:         "GO-UIkit",
		Icon: app.Icon{
			Default:    "/web/images/android-chrome-192x192.png",
			Large:      "/web/images/android-chrome-512x512.png",
			AppleTouch: "/web/images/apple-touch-icon.png",
		},
		Scripts: []string{
			"/web/js/uikit.min.js",
			"/web/js/uikit-icons.min.js",
			"/web/js/prism.js",
		},
		Styles: []string{
			"https://fonts.googleapis.com/css2?family=Fira+Code:wght@300;400;700&family=Montserrat:wght@300;400;700&display=swap",
			"/web/css/uikit.min.css",
			"/web/css/prism.css",
			"/web/css/app.css",
		},
		ThemeColor: backgroundColor,
		Title:      "GO-UIkit",
	}

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, localOpts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}
}

func runLocal(ctx context.Context, h *app.Handler, opts localOptions) {
	app.Log("%s", logs.New("starting GO-UIkit service").
		Tag("port", opts.Port).
		Tag("version", h.Version),
	)

	s := http.Server{
		Addr:    fmt.Sprintf(":%v", opts.Port),
		Handler: h,
	}

	go func() {
		<-ctx.Done()
		s.Shutdown(context.Background())
	}()

	if err := s.ListenAndServe(); err != nil {
		panic(err)
	}
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	pages := pages()
	p := make([]string, 0, len(pages))
	for path := range pages {
		p = append(p, path)
	}

	if err := app.GenerateStaticWebsite(opts.Output, h, p...); err != nil {
		panic(err)
	}
}

func exit() {
	err := recover()
	if err != nil {
		app.Log("command failed: %s", errors.Newf("%v", err))
		os.Exit(-1)
	}
}
