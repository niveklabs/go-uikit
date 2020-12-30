package app

import (
	"context"
	"io"
	"net/url"
	"strings"

	"github.com/maxence-charriere/go-app/v7/pkg/errors"
)

// Raw returns a ui element from the given raw value. HTML raw value must have a
// single root.
//
// It is not recommended to use this kind of node since there is no check on the
// raw string content.
func Raw(v string) UI {
	v = strings.TrimSpace(v)

	tag := rawRootTagName(v)
	if tag == "" {
		v = "<div></div>"
	}

	return &raw{
		value: v,
		tag:   tag,
	}
}

type raw struct {
	jsvalue    Value
	parentElem UI
	tag        string
	value      string
}

func (r *raw) Kind() Kind {
	return RawHTML
}

func (r *raw) JSValue() Value {
	return r.jsvalue
}

func (r *raw) Mounted() bool {
	return r.jsvalue != nil
}

func (r *raw) name() string {
	return "raw." + r.tag
}

func (r *raw) self() UI {
	return r
}

func (r *raw) setSelf(UI) {
}

func (r *raw) context() context.Context {
	return nil
}

func (r *raw) attributes() map[string]string {
	return nil
}

func (r *raw) eventHandlers() map[string]eventHandler {
	return nil
}

func (r *raw) parent() UI {
	return r.parentElem
}

func (r *raw) setParent(p UI) {
	r.parentElem = p
}

func (r *raw) children() []UI {
	return nil
}

func (r *raw) mount() error {
	if r.Mounted() {
		return errors.New("mounting raw html element failed").
			Tag("reason", "already mounted").
			Tag("name", r.name()).
			Tag("kind", r.Kind())
	}

	wrapper := Window().Get("document").Call("createElement", "div")
	wrapper.Set("innerHTML", r.value)

	value := wrapper.Get("firstChild")
	if !value.Truthy() {
		return errors.New("mounting raw html element failed").
			Tag("reason", "converting raw html to html elements returned nil").
			Tag("name", r.name()).
			Tag("kind", r.Kind()).
			Tag("raw-html", r.value)
	}

	wrapper.Call("removeChild", value)
	r.jsvalue = value
	return nil
}

func (r *raw) dismount() {
	r.jsvalue = nil
}

func (r *raw) update(n UI) error {
	if !r.Mounted() {
		return nil
	}

	if n.Kind() != r.Kind() || r.name() != r.name() {
		return errors.New("updating raw html element failed").
			Tag("replace", true).
			Tag("reason", "different element types").
			Tag("current-kind", r.Kind()).
			Tag("current-name", r.name()).
			Tag("updated-kind", n.Kind()).
			Tag("updated-name", n.name())
	}

	if v := n.(*raw).value; r.value != v {
		return errors.New("updating raw html element failed").
			Tag("replace", true).
			Tag("reason", "different raw values").
			Tag("current-value", r.value).
			Tag("new-value", v)
	}

	return nil
}

func (r *raw) onNav(*url.URL) {
}

func (r *raw) html(w io.Writer) {
	r.htmlWithIndent(w, 0)
}

func (r *raw) htmlWithIndent(w io.Writer, indent int) {
	writeIndent(w, indent)
	w.Write(stob(r.value))
	w.Write(ln())
}

func rawRootTagName(raw string) string {
	raw = strings.TrimSpace(raw)

	if strings.HasPrefix(raw, "</") || !strings.HasPrefix(raw, "<") {
		return ""
	}

	end := -1
	for i := 1; i < len(raw); i++ {
		if raw[i] == ' ' ||
			raw[i] == '\t' ||
			raw[i] == '\n' ||
			raw[i] == '>' {
			end = i
			break
		}
	}

	if end <= 0 {
		return ""
	}

	return raw[1:end]
}
