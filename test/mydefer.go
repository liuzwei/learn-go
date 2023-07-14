package test

import (
	"fmt"
	"net/http"
	"os"
)

func deferNil() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

func search() error {
	resp, err := http.Get("https://google.com")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}

	return nil
}

func handleErrors() error {
	f, err := os.Open("C:/Users/admin/Desktop/开发信息.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f *os.File) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close 开发信息.txt error: %v \n", err)
			}
		}(f)
	}

	f, err = os.Open("C:/Users/admin/Desktop/t1.json")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f *os.File) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close t1.json error: %v \n", err)
			}
		}(f)
	}
	return nil
}
