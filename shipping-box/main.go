package main

import (
	"fmt"
	"sort"
)

type Product struct {
	Name string
	Len int
	Wid int
	Hei int
}

type Box struct {
	Len int
	Wid int
	Hei int
}

func getBestBox(availableBoxs []Box, product []Product) Box {
	canUseBaxs := []Box{}
	for _, p := range product {
		pc := p.Hei * p.Len * p.Wid
		for _, b := range availableBoxs {
			bc := b.Wid * b.Len * b.Hei
			if pc >= bc {
				canUseBaxs = append(canUseBaxs, b)
			}
		}
	}
	sort.Slice(canUseBaxs, func(i, j int) bool {
		return canUseBaxs[i].Hei * canUseBaxs[i].Len * canUseBaxs[i].Wid > canUseBaxs[j].Hei * canUseBaxs[j].Len * canUseBaxs[j].Wid
	})
	fmt.Println(canUseBaxs)
	return Box{}
}

func main() {
	b1 := Box{1, 1, 1}
	b2 := Box{2, 2, 2}
	b3 := Box{3, 3, 3}

	p1 := Product{"p1", 5, 2, 5}
	p2 := Product{"p2", 1, 1, 1}
	p3 := Product{"p3", 1, 1, 3}
	allBox := []Box{b1, b2, b3}
	allProduct := []Product{p1, p2, p3}
	bestBox := getBestBox(allBox, allProduct)
	fmt.Println(bestBox)
}
