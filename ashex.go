package main

import "log"
import "io"
import "os"

const SIZE = 64 * 1024

var glyphs = "0123456789ABCDEF"

func main() {
	buf := make([]byte, SIZE)
	out := make([]byte, SIZE*2)
	for {
		n, e := os.Stdin.Read(buf)
		if e != nil {
			if e == io.EOF {
				return
			}
			log.Fatal(e)
		}
		for i := 0; i < n; i++ {
			out[i*2+0] = glyphs[buf[i]>>4]
			out[i*2+1] = glyphs[buf[i]&15]
		}
		_, e = os.Stdout.Write(out[:2*n])
		if e != nil {
			log.Fatal(e)
		}
	}
}
