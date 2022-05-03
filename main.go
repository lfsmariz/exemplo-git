package main

import "fmt"

func main() {

	fmt.Println(chuvasol(true))

}

func chuvasol(b bool) string {
	if b {
		return "chuva"
	}

	return "sol"
}
