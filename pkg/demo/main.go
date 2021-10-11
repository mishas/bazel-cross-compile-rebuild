package main

import (
	"fmt"

	_ "firebase.google.com/go/v4"
	_ "gocloud.dev/blob"

	"example.com/pkg/pb"
)

func main() {
	msg := pb.MyMessage{}
	fmt.Println("Hello world", msg)
}
