package main

import (
	"fmt"

	"github.com/koyachi/go-nude"
)

func main() {

	pic := "pic.jpg"

	query, _ := nude.IsNude(pic)
	fmt.Println(query, "\t", pic)
}
