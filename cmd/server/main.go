package main

import "fmt"

func Run() error {
	fmt.Println("starting app")
	return nil
}

func main() {
	fmt.Println("Go rest")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
