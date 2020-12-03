package main

import (
	"fmt"

	"github.com/setraj/safe-map/concurrent"
)

func main() {
	fmt.Println("main function")
	cMap := concurrent.NewConcurrentMap(5)
	cMap.Add(1, "apple")

	fmt.Println(cMap)
}
