package main

import (
	"fmt"
	"lem-in/control"
	"lem-in/parser"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error: missing filename")
		fmt.Println("USAGE: go run main.go example00.txt")
		return
	}

	filePath := os.Args[1]
	data := parser.ReadFile(filePath)
	graph, ants := parser.ParseData(data)

	paths := control.FindPaths(graph)
	ValidPaths := control.ValidPaths(paths)
	sortedCombPaths := control.SortCombPaths(ValidPaths)
	bestCombPaths := control.BestCombPaths(ants, sortedCombPaths)
	sendAnts := control.SendAnts(ants, data, bestCombPaths)

	turns := 0
	for _, v := range sendAnts {
		fmt.Println(v)
		turns++
	}
	fmt.Println()
	fmt.Printf("Turns number: %v\n", turns)
}
