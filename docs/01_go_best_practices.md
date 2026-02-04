Twelve Go Best Practices

Francesc Campoy Flores
Gopher at Google
@francesc
http://campoy.cat/+

http://golang.org

* Best practices

From Wikipedia:

"A best practice is a method or technique that has consistently shown results superior
to those achieved with other means"

Techniques to write Go code that is

- simple,
- readable,
- maintainable.

http://golang.org/doc/gopher/gopherbw.png

* Some code

```go
package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

// Example of bad code, missing early return. OMIT
func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err == nil {
		size += 4
		var n int
		n, err = w.Write([]byte(g.Name))
		size += int64(n)
		if err == nil {
			err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
			if err == nil {
				size += 4
			}
			return
		}
		return
	}
	return
}

func main() {
	g := &Gopher{
		Name:     "Gophertiti",
		AgeYears: 3382,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}

// Example of bad API, it's better to use an interface.
func (g *Gopher) WriteToFile(f *os.File) (int64, error) {
	return 0, nil
}

// Example of bad API, it's better to use a narrower interface.
func (g *Gopher) WriteToReadWriter(rw io.ReadWriter) (int64, error) {
	return 0, nil
}

// Example of better API.
func (g *Gopher) WriteToWriter(f io.Writer) (int64, error) {
	return 0, nil
}
```

* Avoid nesting by handling errors first

```go
package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

func (g *Gopher) WriteTo(w io.Writer) (size int64, err error) {
	err = binary.Write(w, binary.LittleEndian, int32(len(g.Name)))
	if err != nil {
		return
	}
	size += 4
	n, err := w.Write([]byte(g.Name))
	size += int64(n)
	if err != nil {
		return
	}
	err = binary.Write(w, binary.LittleEndian, int64(g.AgeYears))
	if err == nil {
		size += 4
	}
	return
}

func main() {
	g := &Gopher{
		Name:     "Gophertiti",
		AgeYears: 3382,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}
```

Less nesting means less cognitive load on the reader

* Avoid repetition when possible

Deploy one-off utility types for simpler code

```go
package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

// Write writes a value to the provided writer in little endian form.
func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
		w.size += int64(binary.Size(v))
	}
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(int32(len(g.Name)))
	bw.Write([]byte(g.Name))
	bw.Write(int64(g.AgeYears))
	return bw.size, bw.err
}

func main() {
	g := &Gopher{
		Name:     "Gophertiti",
		AgeYears: 3382,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}
```

* Avoid repetition when possible

Using `binWriter`

.code bestpractices/shortercode3.go /WriteTo/,/^}/

* Type switch to handle special cases

```go
package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

// Write writes a value to the provided writer in little endian form.
func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	switch v.(type) { // HL
	case string:
		s := v.(string)
		w.Write(int32(len(s)))
		w.Write([]byte(s))
	case int:
		i := v.(int)
		w.Write(int64(i))
	default:
		if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
			w.size += int64(binary.Size(v))
		}
	}
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(g.Name) // HL
	bw.Write(g.AgeYears)
	return bw.size, bw.err
}

func main() {
	g := &Gopher{
		Name:     "Gophertiti",
		AgeYears: 3382,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}
```

* Type switch with short variable declaration

```go
package main

import (
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

type binWriter struct {
	w    io.Writer
	size int64
	err  error
}

// Write writes a value to the provided writer in little endian form.
func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	switch x := v.(type) { // HL
	case string:
		w.Write(int32(len(x)))
		w.Write([]byte(x))
	case int:
		w.Write(int64(x))
	default:
		if w.err = binary.Write(w.w, binary.LittleEndian, v); w.err == nil {
			w.size += int64(binary.Size(v))
		}
	}
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
	bw := &binWriter{w: w}
	bw.Write(g.Name)
	bw.Write(g.AgeYears)
	return bw.size, bw.err
}

func main() {
	g := &Gopher{
		Name:     "Gophertiti",
		AgeYears: 3382,
	}

	if _, err := g.WriteTo(os.Stdout); err != nil {
		log.Printf("DumpBinary: %v\n", err)
	}
}
```

* Writing everything or nothing

```go
package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
)

type Gopher struct {
	Name     string
	AgeYears int
}

type binWriter struct {
	w   io.Writer
	buf bytes.Buffer // HL
	err error
}

// Write writes a value to the provided writer in little endian form.
func (w *binWriter) Write(v interface{}) {
	if w.err != nil {
		return
	}
	switch x := v.(type) {
	case string:
		w.Write(int32(len(x)))
		w.Write([]byte(x))
	case int:
		w.Write(int64(x))
	default:
		w.err = binary.Write(&w.buf, binary.LittleEndian, v) // HL
	}
}
```

* Writing everything or nothing

```go
// Flush writes any pending values into the writer if no error has occurred.
// If an error has occurred, earlier or with a write by Flush, the error is
// returned.
func (w *binWriter) Flush() (int64, error) {
if w.err != nil {
return 0, w.err
}
return w.buf.WriteTo(w.w)
}

func (g *Gopher) WriteTo(w io.Writer) (int64, error) {
bw := &binWriter{w: w}
bw.Write(g.Name)
bw.Write(g.AgeYears)
return bw.Flush() // HL
}

func main() {
g := &Gopher{
Name:     "Gophertiti",
AgeYears: 3382,
}

if _, err := g.WriteTo(os.Stdout); err != nil {
log.Printf("DumpBinary: %v\n", err)
}
}
```

* Function adapters

```go
package bestpractices

import (
	"fmt"
	"log"
	"net/http"
)

func doThis() error { return nil }
func doThat() error { return nil }

// HANDLER1 OMIT
func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := doThis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}

	err = doThat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}
}
```

* Function adapters

```go
func init() {
http.HandleFunc("/", errorHandler(betterHandler))
}

func errorHandler(f func (http.ResponseWriter, *http.Request) error) http.HandlerFunc {
return func (w http.ResponseWriter, r *http.Request) {
err := f(w, r)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
log.Printf("handling %q: %v", r.RequestURI, err)
}
}
}

func betterHandler(w http.ResponseWriter, r *http.Request) error {
if err := doThis(); err != nil {
return fmt.Errorf("doing this: %v", err)
}

if err := doThat(); err != nil {
return fmt.Errorf("doing that: %v", err)
}
return nil
}
```

* Organizing your code

* Important code goes first

License information, build tags, package documentation.

Import statements, related groups separated by blank lines.

	import (
		"fmt"
		"io"
		"log"

		"code.google.com/p/go.net/websocket"
	)

The rest of the code starting with the most significant types, and ending
with helper function and types.

* Document your code

Package name, with the associated documentation before.

	// Package playground registers an HTTP handler at "/compile" that
	// proxies requests to the golang.org playground service.
	package playground

Exported identifiers appear in `godoc`, they should be documented correctly.

	// Author represents the person who wrote and/or is presenting the document.
	type Author struct {
		Elem []Elem
	}

	// TextElem returns the first text elements of the author details.
	// This is used to display the author' name, job title, and company
	// without the contact details.
	func (p *Author) TextElem() (elems []Elem) {

[[http://godoc.org/code.google.com/p/go.talks/pkg/present#Author][Generated documentation]]

[[http://blog.golang.org/godoc-documenting-go-code][Gocode: documenting Go code]]

* Shorter is better

or at least _longer_is_not_always_better_.

Try to find the *shortest*name*that*is*self*explanatory*.

- Prefer `MarshalIndent` to `MarshalWithIndentation`.

Don't forget that the package name will appear before the identifier you chose.

- In package `encoding/json` we find the type `Encoder`, not `JSONEncoder`.

- It is referred as `json.Encoder`.

* Packages with multiple files

Should you split a package into multiple files?

- Avoid very long files

The `net/http` package from the standard library contains 15734 lines in 47 files.

- Separate code and tests

`net/http/cookie.go` and `net/http/cookie_test.go` are both part of the `http`
package.

Test code is compiled *only* at test time.

- Separated package documentation

When we have more than one file in a package, it's convention to create a `doc.go`
containing the package documentation.

* Make your packages "go get"-able

Some packages are potentially reusable, some others are not.

A package defining some network protocol might be reused while one defining
an executable command may not.

.image bestpractices/cmd.png

[[https://github.com/bradfitz/camlistore]]

* APIs

* Ask for what you need

Let's use the Gopher type from before

.code bestpractices/shortercode1.go /type Gopher/,/^}/

We could define this method

.code bestpractices/shortercode1.go /WriteToFile/

But using a concrete type makes this code difficult to test, so we use an interface.

.code bestpractices/shortercode1.go /WriteToReadWriter/

And, since we're using an interface, we should ask only for the methods we need.

.code bestpractices/shortercode1.go /WriteToWriter/

* Keep independent packages independent

.code bestpractices/funcdraw/cmd/funcdraw.go /IMPORT/,/ENDIMPORT/

.code bestpractices/funcdraw/cmd/funcdraw.go /START/,/END/

* Parsing

.code bestpractices/funcdraw/parser/parser.go /START/,/END/

* Drawing

.code bestpractices/funcdraw/drawer/dependent.go /START/,/END/

Avoid dependency by using an interface.

.code bestpractices/funcdraw/drawer/drawer.go /START/,/END/

* Testing

Using an interface instead of a concrete type makes testing easier.

.code bestpractices/funcdraw/drawer/drawer_test.go ,/END/

* Avoid concurrency in your API

.play bestpractices/concurrency1.go /START/,/END/

What if we want to use it sequentially?

* Avoid concurrency in your API

.play bestpractices/concurrency2.go /START/,/END/

Expose synchronous APIs, calling them concurrently is easy.

* Best practices for concurrency

* Use goroutines to manage state

Use a chan or a struct with a chan to communicate with a goroutine

.code bestpractices/server.go /START/,/STOP/

* Use goroutines to manage state (continued)

.play bestpractices/server.go /STOP/,

* Avoid goroutine leaks with buffered chans

.code bestpractices/bufchan.go /SEND/,/BROADCAST/

.code bestpractices/bufchan.go /MAIN/,

* Avoid goroutine leaks with buffered chans (continued)

.play bestpractices/bufchan.go /BROADCAST/,/MAIN/

- the goroutine is blocked on the chan write
- the goroutine holds a reference to the chan
- the chan will never be garbage collected

* Avoid goroutines leaks with buffered chans (continued)

.play bestpractices/bufchanfix.go /BROADCAST/,/MAIN/

- what if we can't predict the capacity of the channel?

* Avoid goroutines leaks with quit chan

.play bestpractices/quitchan.go /BROADCAST/,/MAIN/

* Twelve best practices

1. Avoid nesting by handling errors first
2. Avoid repetition when possible
3. Important code goes first
4. Document your code
5. Shorter is better
6. Packages with multiple files
7. Make your packages "go get"-able
8. Ask for what you need
9. Keep independent packages independent
10. Avoid concurrency in your API
11. Use goroutines to manage state
12. Avoid goroutine leaks

* Some links

Resources

- Go homepage [[http://golang.org]]
- Go interactive tour [[http://tour.golang.org]]

Other talks

- Lexical scanning with Go [[http://www.youtube.com/watch?v=HxaD_trXwRE][video]]
- Concurrency is not parallelism [[http://vimeo.com/49718712][video]]
- Go concurrency patterns [[http://www.youtube.com/watch?v=f6kdp27TYZs][video]]
- Advanced Go concurrency patterns [[http://www.youtube.com/watch?v=QDDwwePbDtw][video]]