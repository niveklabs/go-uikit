//go:build wasm
// +build wasm

package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
    app.Route("/", &page{})

	app.RunWhenOnBrowser()

}
