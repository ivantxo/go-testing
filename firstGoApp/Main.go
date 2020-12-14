package main

import (
	"bytes"
	"fmt"
)

func main() {
	var wc WriterCloser = NewBufferWriterCloser()
	wc.Write([]byte("Hello YouTube listeners, this is a test"))
	wc.Close()

	// type conversion
	// r, ok := wc.(io.Reader) // conversion fails, but application doesn't crash
	r, ok := wc.(*BufferWriterCloser)
	if ok {
		fmt.Println("Type converted successfully:", r)
	} else {
		fmt.Println("Conversion failed")
	}
}

// Writer Interface
type Writer interface {
	Write([]byte) (int, error)
}

// Closer Interface
type Closer interface {
	Close() error
}

// WriterCloser Interface
type WriterCloser interface {
	Writer
	Closer
}

// BufferWriterCloser struct
type BufferWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)
	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Close function
func (bwc *BufferWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewBufferWriterCloser function
func NewBufferWriterCloser() *BufferWriterCloser {
	return &BufferWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}
