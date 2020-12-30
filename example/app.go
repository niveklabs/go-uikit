// +build wasm

package main

import (
	"github.com/maxence-charriere/go-app/v7/pkg/app"
)

func main() {
	app.Route("/", &page{})

	app.Run()
}
