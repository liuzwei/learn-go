package test

import (
	"fmt"
	"testing"
)

func TestDeferNil(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	deferNil()
}

func TestSearch(t *testing.T) {
	err := search()
	fmt.Println(err)
}

func TestHandleErrors(t *testing.T) {
	err := handleErrors()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
}
