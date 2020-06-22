package main

import (
	"fmt"
	api "github.com/tuweiguang/GoTestLib/api"
	context "golang.org/x/net/context"
)

func main() {
	api.MyApi()

	c := context.Background()
	if c == nil {
		fmt.Println("Background returned nil")
	}
	select {
	case x := <-c.Done():
		fmt.Printf("<-c.Done() == %v want nothing (it should block)", x)
	default:
	}

	if got, want := fmt.Sprint(c), "context.Background"; got != want {
		fmt.Printf("Background().String() = %q want %q", got, want)
	} else {
		fmt.Println(got, want)
	}

	fmt.Println("end")

	fmt.Println("hello jenkins")

	fmt
}
