package main

import (
	"github.com/mintance/go-uniqid/uniqid"
	"fmt"
)

func main()  {
	fmt.Println("id: ", uniqid.New(uniqid.Params{"test", true}))
}