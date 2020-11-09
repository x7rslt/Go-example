package go_example_test

import (
	"fmt"
	"os"
	"testing"
)

//Quick test file.

func TestArgs(t *testing.T) {
	content := os.Args[1:]
	fmt.Println(content)
}
