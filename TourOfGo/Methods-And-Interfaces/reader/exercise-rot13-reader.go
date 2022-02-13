package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot rot13Reader) Read(arr []byte) (int, error) {
	n, err := rot.r.Read(arr)
	//if err == io.EOF {
	//	return 0, err
	//}

	for i := 0; i < n; i++ {
		if arr[i] >= 65 && arr[i] <= 90 {
			arr[i] = 65 + (arr[i]-65+13)%26
		} else if arr[i] >= 97 && arr[i] <= 122 {
			arr[i] = 97 + (arr[i]-97+13)%26
		}
	}

	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
