package main

import (
	"fmt"

	checker "github.com/savisitor15/go-dedupe/internal/checker"
)

func main() {
	// app entry
	fmt.Println("checking if sample/blank.png exists")
	if sum, err := checker.FileToSHA256("sample/blank.png"); err != nil {
		fmt.Printf("Error!\n err: %s\n", err)
	} else {
		fmt.Printf("file: sample/blank.png sum: %x\n", sum)
	}
}
