package path_test

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPath(t *testing.T) {
	//current path
	path, _ := os.Getwd()
	fmt.Println("Current path:", path)

	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println(p)

	//You shoulkd always use Join instead of concatenating manually.
	//Join can remove superfluous separators
	fmt.Println(filepath.Join("dir1//", "filename"))
	fmt.Println(filepath.Join("dir1/../dir1", "filename"))

	//dir and base split path with filename
	fmt.Println("Dir(p)", filepath.Dir(p))
	fmt.Println("Base(p)", filepath.Base(p))

	//check absoulute
	fmt.Println(filepath.IsAbs("dir/file"))
	fmt.Println(filepath.IsAbs("/dir/file"))

	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)                               //.ext
	fmt.Println(strings.TrimSuffix(filename, ext)) //filname

}
