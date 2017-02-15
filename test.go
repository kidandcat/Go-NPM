package main

import (
	"fmt"

	"./node_modules/go-locate/"
)

func main() {
	fmt.Println("Locate IP 8.8.8.8")
	body, status := locate.Locate("8.8.8.8")
	fmt.Println(status, body)
}
