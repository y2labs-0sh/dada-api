//go:generate go run generator.go

package box

import (
	"bytes"
	"compress/zlib"
	"io"
)

const SizeLimit = 20 * 1024 * 1024

type embedBox struct {
	storage map[string][]byte
}

// Create new box for embed files
func newEmbedBox() *embedBox {
	return &embedBox{storage: make(map[string][]byte)}
}

// Add a file to box
func (e *embedBox) Add(file string, content []byte) {
	b := new(bytes.Buffer)
	w, err := zlib.NewWriterLevel(b, zlib.BestCompression)
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(content); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	e.storage[file] = b.Bytes()
}

// Get file's content
func (e *embedBox) Get(file string) []byte {
	if f, ok := e.storage[file]; ok {
		r, err := zlib.NewReader(bytes.NewReader(f))
		if err != nil {
			panic(err)
		}
		bc := make([]byte, 0)
		buf := bytes.NewBuffer(bc)
		if _, err := io.Copy(buf, io.LimitReader(r, SizeLimit)); err != nil {
			panic(err)
		}
		if err := r.Close(); err != nil {
			panic(err)
		}
		return buf.Bytes()
	}
	return nil
}

// Embed box expose
var box = newEmbedBox()

// Add a file content to box
func Add(file string, content []byte) {
	box.Add(file, content)
}

// Get a file from box
func Get(file string) []byte {
	return box.Get(file)
}
