package main

import (
	"fmt"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		go fmt.Println("makeThumbnails", f)
	}
}

func makeThumbnails2(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			fmt.Println("makeThumbnails2", f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

func makeThumbnails3(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			if f == "1.go" {
				errors <- fmt.Errorf("error: %s\n", f)
			}
			errors <- nil
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}

func main() {
	filenames := []string{
		"1.go",
		"2.go",
		"3.go",
	}

	// makeThumbnails(filenames)
	// makeThumbnails2(filenames)
	err := makeThumbnails3(filenames)
	if err != nil {
		fmt.Println("out:", err)
	}
}
