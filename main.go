package main

import "fmt"

func main() {
	LoadResource()
	if err := Run(); err != nil {
		fmt.Printf("some error occurs: %v\n", err)
	}
}

func LoadResource() {

}

func Run() error {
	return nil
}
