package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/lemolatoon/core"
)

func usage() {
	fmt.Println("Usage: omikuji")
	fmt.Println("	   try [number of try]")
}

func main() {
	if len(os.Args) < 2 {
		usage()
		return
	}
	switch os.Args[1] {
	case "try":
		try(os.Args[2:])

	}
}

func try(args []string) {
	if len(args) >= 1 {
		v, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(v)
		tryInner(v)
	} else {
		tryInner(1)
	}

}

// -1 for default value (1000000)
func tryInner(nTry int) {
	if nTry == 1 {
		fmt.Println(core.Omikuji())
		return
	}
	const defaultNTry = 1000000
	if nTry < 0 {
		nTry = defaultNTry
	}
	omikujiMap := make(map[core.OmikujiResulter]int)
	for i := 0; i < nTry; i++ {
		omikujiMap[core.Omikuji()]++
	}

	sum := 0
	for _, v := range omikujiMap {
		sum += v
	}

	for k, v := range omikujiMap {
		fmt.Printf("%s: %f (%v)\n", k, float64(v)/float64(sum)*100, k.Rate()*100)
	}
	fmt.Println(omikujiMap)
}
