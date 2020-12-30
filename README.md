**go-uikit** is a package to be used with [go-app](https://github.com/maxence-charriere/go-app) to build [progressive web apps (PWA)](https://developers.google.com/web/progressive-web-apps/) with, [UIkit](https://getuikit.com/) front-end framework, [Go programming language](https://golang.org) and [WebAssembly](https://webassembly.org).

### Components

- Accordion
- Alert
- Container
- Section

### Example

Compile WASM

```bash
cd example
GOARCH=wasm GOOS=js go build -v -o ./web/app.wasm .
```

Compile Local server

```bash
go build -v -o go-uikit-serve .
./go-uikit-serve local
```