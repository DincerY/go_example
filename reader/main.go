package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type MyReader struct{}

func (reader MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

type rot13Reader struct {
	r io.Reader
}

func rot13(b byte) byte {
	switch {
	case 'A' <= b && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case 'a' <= b && b <= 'z':
		return 'a' + (b-'a'+13)%26
	default:
		return b
	}
}

func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot13(b[i])
	}
	return len(b), err
}

func main() {
	/*r := strings.NewReader("Hello, Reader!")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}*/

	var ioReaderInterface io.Reader
	ioReaderInterface = &strings.Reader{}
	ioReaderInterface = MyReader{}
	ioReaderInterface = &rot13Reader{}

	res, err := http.Get("https://localhost:7301")
	if err == nil {
		return
	}
	ioReaderInterface = res.Body

	opened, err := os.Open("test.txt")
	if err == nil {
		return
	}
	ioReaderInterface = opened

	fmt.Println(ioReaderInterface)

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)

}
