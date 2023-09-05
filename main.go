package main

import (
	"fmt"

	"github.com/lemolatoon/core"
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(core.Omikuji())
	const NTry = 1000000
	omikujiMap := make(map[core.OmikujiResulter]int)
	for i := 0; i < NTry; i++ {
		omikujiMap[core.Omikuji()]++
	}
	fmt.Println(omikujiMap)

	sum := 0
	for _, v := range omikujiMap {
		sum += v
	}

	for k, v := range omikujiMap {
		fmt.Printf("%s: %f (%v)\n", k, float64(v)/float64(sum)*100, k.Rate()*100)
	}
}
