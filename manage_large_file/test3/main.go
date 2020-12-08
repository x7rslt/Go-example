package main

//so low ,too
import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"time"
)

// uses scanner.Bytes to avoid allocation.
func main() {
	start1 := time.Now()
	f, err := os.Open(`E:\SocialData\person_data\weibo.txt`)

	scanner := bufio.NewScanner(f)
	if err != nil {
		panic(err)
	}
	fmt.Println("Open spend time:", time.Now().Sub(start1))
	defer f.Close()
	start2 := time.Now()
	toFind := []byte("15555001522")

	for scanner.Scan() {
		if bytes.Contains(scanner.Bytes(), toFind) {
			fmt.Println(scanner.Text())
			fmt.Println("Find spend time:", time.Now().Sub(start2))
			fmt.Println("Congratulations!")
			os.Exit(0)

		}
	}
	fmt.Println("Sorry Not Found.")
	fmt.Println("All done spend time:", time.Now().Sub(start1))
}
