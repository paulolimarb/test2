package test2

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome! test2", name)
	return message
}
