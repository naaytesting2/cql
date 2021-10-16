package main

import (
	"errors"
	"fmt"
)

func b() (string, error) {
	return "hello", errors.New("error B")
}


func a() (n int, err error) {
	if n != 0 {
		return 1, nil
	}

	resolve := func() {
		output, err := b()
		if err != nil {
			err = fmt.Errorf("failed to list modules: %v\n%s", err, output)
			return
		}
	}

	resolve()
	return n, err
}

func main() {
	fmt.Println(a())
}
