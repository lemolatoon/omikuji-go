package core

import (
	"fmt"
	"math/rand"
)

const daikyoWeight = 1
const kyoWeight = 100
const shokichiWeight = 300
const suekichiWeight = 500
const kichiWeight = 2000
const chuukichiWeight = 3000
const daikichiWeight = 130
const daidaikichiWeight = 10

var weights = [8]int{daikyoWeight, kyoWeight, shokichiWeight, suekichiWeight, kichiWeight, chuukichiWeight, daikichiWeight, daidaikichiWeight}
var accumulatedWeight = func() []int {
	acc := make([]int, len(weights))
	for i, w := range weights {
		if i == 0 {
			acc[i] = w
		} else {
			acc[i] = acc[i-1] + w
		}
	}
	return acc
}()

func NewDaikyo() OmikujiResulter {
	return omikujiResult{name: "大凶", index: 0}
}
func NewKyo() OmikujiResulter {
	return omikujiResult{name: "凶", index: 1}
}
func NewShokichi() OmikujiResulter {
	return omikujiResult{name: "小吉", index: 2}
}
func NewSuekichi() OmikujiResulter {
	return omikujiResult{name: "末吉", index: 3}
}
func NewKichi() OmikujiResulter {
	return omikujiResult{name: "吉", index: 4}
}
func NewChuukichi() OmikujiResulter {
	return omikujiResult{name: "中吉", index: 5}
}
func NewDaikichi() OmikujiResulter {
	return omikujiResult{name: "大吉", index: 6}
}
func NewDaidaikichi() OmikujiResulter {
	return omikujiResult{name: "大大吉", index: 7}
}

func NewOmikujiResultFromIndex(index int) OmikujiResulter {
	switch index {
	case 0:
		return NewDaikyo()
	case 1:
		return NewKyo()
	case 2:
		return NewShokichi()
	case 3:
		return NewSuekichi()
	case 4:
		return NewKichi()
	case 5:
		return NewChuukichi()
	case 6:
		return NewDaikichi()
	case 7:
		return NewDaidaikichi()
	default:
		panic(fmt.Sprintf("invalid index: %d", index))
	}
}

type OmikujiResulter interface {
	fmt.Stringer
	Rate() float64
}

type omikujiResult struct {
	index int
	name  string
}

func (o omikujiResult) String() string {
	return o.name
}

func (o omikujiResult) Rate() float64 {
	return float64(weights[o.index]) / float64(accumulatedWeight[len(accumulatedWeight)-1])
}

func (o *omikujiResult) threshold() int {
	return accumulatedWeight[o.index]
}

func Omikuji() OmikujiResulter {
	random := rand.Intn(accumulatedWeight[len(accumulatedWeight)-1])
	for index, threshold := range accumulatedWeight {
		if random < threshold {
			return NewOmikujiResultFromIndex(index)
		}
	}
	return nil
}
