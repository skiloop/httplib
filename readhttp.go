package httplib

import (
	"net"
	"bytes"
	"fmt"
)

var newLine = []byte("\r\n")
var slashRBytes = []byte("\r")

func HttpReadHeader(c net.Conn) (data []byte, err error) {
	buf := bytes.Buffer{}
	buff := make([]byte, 1)
	pre := byte(0)

	lineHasData := false
	for {
		_, err := c.Read(buff)
		if nil != err {
			return nil, err
		}
		if buff[0] == 10 && pre == 13 {
			if lineHasData {
				fmt.Print(pre)
				fmt.Print(buff[0])
				lineHasData = false
				pre = 10
				buf.Write(newLine)
				continue
			} else {
				fmt.Print("<>")
				break
			}
		}
		fmt.Printf("%c", buff[0])

		lineHasData = true
		if pre != 13 {
			buf.Write(buff)
		} else {
			buf.Write(slashRBytes)
			buf.Write(buff)
		}
		pre = buff[0]
	}
	return buf.Bytes(), nil
}
