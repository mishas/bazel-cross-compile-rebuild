package main

import (
	"fmt"

	"example.com/pkg/pb"
)

func main() {
	msg := &pb.MyMessage{Id: 49}
	fmt.Println("MyMessage: ", msg)
}
