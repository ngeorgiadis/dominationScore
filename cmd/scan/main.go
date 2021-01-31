package main

import (
	"fmt"

	"github.com/ngeorgiadis/dominationScore/internal/dcounter2"
)

func main() {

	_, d1, d2, d3, d4 := dcounter2.ReadDataset("../../data/nodes_all.csv")

	fmt.Println(len(d1), len(d2), len(d3), len(d4))

}
