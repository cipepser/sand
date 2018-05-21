package main

import (
	"strconv"

	"github.com/Arafatk/DataViz/trees/btree"
)

func main() {
	n := 8
	bt := btree.NewWithIntComparator(n)

	for i := 0; i < n; i++ {
		bt.Put(i, i)
		bt.Visualizer("img/btee" + strconv.Itoa(i) + ".png")
	}

}
