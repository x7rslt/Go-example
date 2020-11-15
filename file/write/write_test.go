package write_test

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

func TestWrite(t *testing.T) {
	path, _ := os.Getwd()
	filepath := path + "/news.txt"
	data, err := ioutil.ReadFile(filepath)
	check(err)
	//ioutil.ReadFile read is []byte,if print for string use string(data)
	fmt.Println(string(data))
	newcontent := []byte("write test")
	for _, i := range newcontent {
		data = append(data, i)
	}
	err = ioutil.WriteFile(filepath, data, 0644)
	check(err)

	newpath := path + "/news2.txt"
	f, err := os.Create(newpath)
	check(err)
	defer f.Close()

	//write []byte
	d2 := []byte{115, 111, 109, 101}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	//write string
	n3, err := f.WriteString("write string\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	//bufio prvides buffered writers in addition to the buffered readers.
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("%d write bytes\n", n4)
	w.Flush() //use flush to ensure all buffered operations have been applied to the underlying writer.

}
