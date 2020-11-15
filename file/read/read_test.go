package read_test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestReadFile(t *testing.T) {
	//Read a file entire contents into memory.
	path, _ := os.Getwd()
	filepath := path + "/news.txt"
	data, err := ioutil.ReadFile(filepath)
	check(err)
	fmt.Println(string(data))
}

func openfile() *os.File {
	path, _ := os.Getwd()
	filepath := path + "/news.txt"
	f, err := os.Open(filepath)
	check(err)
	return f
}

func TestOpenFile(t *testing.T) {
	f := openfile()
	defer f.Close()
	//Read some bytes from the begining of the file.
	b1 := make([]byte, 15)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes :%s\n", n1, string(b1[:n1]))
}

func TestBufio(t *testing.T) {
	//bufio implements a buffered reader that may be useful both for its efficiency with many small reads.
	f := openfile()
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes:%s\n", string(b4))
}
