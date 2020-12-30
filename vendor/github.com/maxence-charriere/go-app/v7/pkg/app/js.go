package app

import "net/url"

// Type represents the JavaScript type of a Value.
type Type int

// Constants that enumerates the JavaScript types.
const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

// Wrapper is implemented by types that are backed by a JavaScript value.
type Wrapper interface {
	JSValue() Value
}

// Value is the interface that represents a JavaScript value. On wasm
// architecture, it wraps the Value from https://golang.org/pkg/syscall/js/
// package.
type Value interface {
	// Bool returns the value v as a bool. It panics if v is not a JavaScript
	// boolean.
	Bool() bool

	// Call does a JavaScript call to the method m of value v with the given
	// arguments. It panics if v has no method m. The arguments get mapped to
	// JavaScript values according to the ValueOf function.
	Call(m string, args ...interface{}) Value

	// Float returns the value v as a float64. It panics if v is not a
	// JavaScript number.
	Float() float64

	// Get returns the JavaScript property p of value v. It panics if v is not a
	// JavaScript object.
	Get(p string) Value

	// Index returns JavaScript index i of value v. It panics if v is not a
	// JavaScript object.
	Index(i int) Value

	// InstanceOf reports whether v is an instance of type t according to
	// JavaScript's instanceof operator.
	InstanceOf(t Value) bool

	// Int returns the value v truncated to an int. It panics if v is not a
	// JavaScript number.
	Int() int

	// Invoke does a JavaScript call of the value v with the given arguments. It
	// panics if v is not a JavaScript function. The arguments get mapped to
	// JavaScript values according to the ValueOf function.
	Invoke(args ...interface{}) Value

	// IsNaN reports whether v is the JavaScript value "NaN".
	IsNaN() bool

	// IsNull reports whether v is the JavaScript value "null".
	IsNull() bool

	// IsUndefined reports whether v is the JavaScript value "undefined".
	IsUndefined() bool

	// JSValue implements Wrapper interface.
	JSValue() Value

	// Length returns the JavaScript property "length" of v. It panics if v is
	// not a JavaScript object.
	Length() int

	// New uses JavaScript's "new" operator with value v as constructor and the
	// given arguments. It panics if v is not a JavaScript function. The
	// arguments get mapped to JavaScript values according to the ValueOf
	// function.
	New(args ...interface{}) Value

	// Set sets the JavaScript property p of value v to ValueOf(x). It panics if
	// v is not a JavaScript object.
	Set(p string, x interface{})

	// SetIndex sets the JavaScript index i of value v to ValueOf(x). It panics
	// if v is not a JavaScript object.
	SetIndex(i int, x interface{})

	// String returns the value v as a string. String is a special case because
	// of Go's String method convention. Unlike the other getters, it does not
	// panic if v's Type is not TypeString. Instead, it returns a string of the
	// form "<T>" or "<T: V>" where T is v's type and V is a string
	// representation of v's value.
	String() string

	// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
	// false, 0, "", null, undefined, and NaN are "falsy", and everything else
	// is "truthy". See
	// https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
	Truthy() bool

	// Type returns the JavaScript type of the value v. It is similar to
	// JavaScript's typeof operator, except that it returns TypeNull instead of
	// TypeObject for null.
	Type() Type
}

// Null returns the JavaScript value "null".
func Null() Value {
	return null()
}

// Undefined returns the JavaScript value "undefined".
func Undefined() Value {
	return undefined()
}

// ValueOf returns x as a JavaScript value:
//
//  | Go                     | JavaScript             |
//  | ---------------------- | ---------------------- |
//  | js.Value               | [its value]            |
//  | js.Func                | function               |
//  | nil                    | null                   |
//  | bool                   | boolean                |
//  | integers and floats    | number                 |
//  | string                 | string                 |
//  | []interface{}          | new array              |
//  | map[string]interface{} | new object             |
//
// Panics if x is not one of the expected types.
func ValueOf(x interface{}) Value {
	return valueOf(x)
}

// Func is the interface that describes a wrapped Go function to be called by
// JavaScript.
type Func interface {
	Value

	// Release frees up resources allocated for the function. The function must
	// not be invoked after calling Release.
	Release()
}

// FuncOf returns a wrapped function.
//
// Invoking the JavaScript function will synchronously call the Go function fn
// with the value of JavaScript's "this" keyword and the arguments of the
// invocation. The return value of the invocation is the result of the Go
// function mapped back to JavaScript according to ValueOf.
//
// A wrapped function triggered during a call from Go to JavaScript gets
// executed on the same goroutine. A wrapped function triggered by JavaScript's
// event loop gets executed on an extra goroutine. Blocking operations in the
// wrapped function will block the event loop. As a consequence, if one wrapped
// function blocks, other wrapped funcs will not be processed. A blocking
// function should therefore explicitly start a new goroutine.
//
// Func.Release must be called to free up resources when the function will not
// be used any more.
func FuncOf(fn func(this Value, args []Value) interface{}) Func {
	return funcOf(fn)
}

// BrowserWindow is the interface that describes the browser window.
type BrowserWindow interface {
	Value

	// The window current url (window.location.href).
	URL() *url.URL

	// The window size.
	Size() (w, h int)

	// The position of the cursor (mouse or touch).
	CursorPosition() (x, y int)

	setCursorPosition(x, y int)

	// Returns the HTML element with the id property that matches the given id.
	GetElementByID(id string) Value

	// Scrolls to the HTML element with the given id.
	ScrollToID(id string)

	// AddEventListener subscribes a given handler to the specified event. It
	// returns a function that must be called to unsubscribe the handler and
	// release allocated resources.
	AddEventListener(event string, h EventHandler) func()
}

// Event is the interface that describes a javascript event.
type Event struct {
	Value
}

// PreventDefault cancels the event if it is cancelable. The default action that
// belongs to the event will not occur.
func (e Event) PreventDefault() {
	e.Call("preventDefault")
}

// CopyBytesToGo copies bytes from the Uint8Array src to dst. It returns the
// number of bytes copied, which will be the minimum of the lengths of src and
// dst.
//
// CopyBytesToGo panics if src is not an Uint8Array.
func CopyBytesToGo(dst []byte, src Value) int {
	return copyBytesToGo(dst, src)
}

// CopyBytesToJS copies bytes from src to the Uint8Array dst. It returns the
// number of bytes copied, which will be the minimum of the lengths of src and
// dst.
//
// CopyBytesToJS panics if dst is not an Uint8Array.
func CopyBytesToJS(dst Value, src []byte) int {
	return copyBytesToJS(dst, src)
}
