**go-uikit** is a package to be used with [go-app](https://github.com/maxence-charriere/go-app) to build [progressive web apps (PWA)](https://developers.google.com/web/progressive-web-apps/) with, [UIkit](https://getuikit.com/) front-end framework, [Go programming language](https://golang.org) and [WebAssembly](https://webassembly.org).

### Components

- Accordion
- Alert
- Article
- Breadcrumb
- Button
- Card
- Container
- Grid
- Leader
- Lightbox
- Marker
- Section

### Building Docs

Compile WASM

```bash
cd docs/
GOARCH=wasm GOOS=js go build -v -o ./web/app.wasm ./src
```

Compile Local server

```bash
go build -v -o go-uikit-serve ./src
./go-uikit-serve local
```