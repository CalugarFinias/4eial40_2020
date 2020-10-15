package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	re io.Reader
}

func main() {
	s := strings.NewReader(" S a l u t l e s a m i s !") // creation of Reader , read a string, creation object
	r := spaceEraser{s}                                  // create us Space Eraser who take a member of reader type
	io.Copy(os.Stdout, &r)
}

func (r spaceEraser) Read(p []byte) (int, error) { // the int for the number  of readed charcaters
	//the error is returned when we arrived at the end of file
	n, err := r.re.Read(p)
	j := 0
	for i := 0; i < n; i++ {
		if p[i] != 32 { // if ellement of buffer different of space
			p[j] = p[i] // if not a space i copy the caracter in my byte array
			j++         // counter for the non space caracters
		}
	}
	return j, err // i return the number of space caracters and err
}
