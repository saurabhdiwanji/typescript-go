package main

import (
	"fmt"
	"io"
	"time"
)

type osSys struct {
	writer             io.Writer
	defaultLibraryPath string
	newLine            string
	cwd                string
}

func (s *osSys) Now() time.Time {
	return time.Now()
}

// implement FS function

// close FS function

func main() {
	fmt.Println("Hello, World!")
}
