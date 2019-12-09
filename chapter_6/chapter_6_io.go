package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	data := []byte("this is my go,hello world \r\n I like it!")
	rd := bytes.NewReader(data)
	nrd := bufio.NewReader(rd)
	//Read
	/*var buf [128]byte
	n, err := nrd.Read(buf[:])
	fmt.Println(string(buf[:n]), n, err) */

	//ReadByte
	/*b, err := nrd.ReadByte()
	fmt.Println(string(b), err)*/

	//ReadBytes
	/*	var delim byte = ','
		line, err := nrd.ReadBytes(delim)
		fmt.Println(string(line), err)
	*/
	//ReadLine
	line, prefix, err := nrd.ReadLine()
	fmt.Println(string(line), prefix, err)
}
