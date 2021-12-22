package cather

import "fmt"

func HandlerError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
