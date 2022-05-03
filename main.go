package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println(chuvasol(randomSun()))

}

func chuvasol(b bool) string {
	if b {
		return "chuva"
	}

	return "sol"
}

func randomSun() bool {
	i := rand.Intn(100)

	if i > 50 {
		return true
	}

	return false
}
