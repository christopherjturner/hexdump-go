package main

import (
	"bufio"
	"bytes"
	"os"
	//	"strings"
)

const hex string = "0123456789abcdef"

func int2hex(i int64, sb *bytes.Buffer) {
	sb.WriteByte(hex[(i>>24&0xff)/16])
	sb.WriteByte(hex[(i>>24&0xff)%16])
	sb.WriteByte(hex[(i>>16&0xff)/16])
	sb.WriteByte(hex[(i>>16&0xff)%16])
	sb.WriteByte(hex[(i>>8&0xff)/16])
	sb.WriteByte(hex[(i>>8&0xff)%16])
	sb.WriteByte(hex[(i&0xff)/16])
	sb.WriteByte(hex[(i&0xff)%16])
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	const BUF_SIZE = 16
	const LINE_MAX = 32

	buf := make([]byte, BUF_SIZE)
	var pos int64
	line := 0

	var sb bytes.Buffer
	sb.Grow(80 * LINE_MAX)
	for {
		n, err := reader.Read(buf)
		if err != nil {
			break
		}

		if n == 0 {
			break
		}

		// print offset
		int2hex(pos, &sb)
		sb.WriteByte(' ')
		sb.WriteByte(' ')

		// print hex value
		for i := 1; i <= n; i++ {
			sb.WriteByte(hex[int(buf[i-1]/16)])
			sb.WriteByte(hex[int(buf[i-1]%16)])
			sb.WriteByte(' ')
			if i%8 == 0 {
				sb.WriteByte(' ')
			}
		}

		// pad line
		if n < BUF_SIZE {
			sb.Write(bytes.Repeat([]byte{' '}, ((BUF_SIZE-n)*3)+1))
		}

		// prints text version
		sb.WriteByte('|')
		for i := 0; i < BUF_SIZE; i++ {
			if i > n {
				sb.WriteByte(' ')
			} else if buf[i] > 32 && buf[i] < 126 {
				sb.WriteByte(buf[i])
			} else {
				sb.WriteByte('.')
			}
		}
		sb.WriteByte('|')
		sb.WriteByte('\n')

		pos += int64(n)
		line += 1
		if line == LINE_MAX-1 {
			sb.WriteTo(os.Stdout)
			line = 0
			sb.Reset()
		}

	}
	sb.WriteTo(os.Stdout)

}
