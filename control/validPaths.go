package control

import (
	"lem-in/types"
)

func ValidPaths(paths types.Paths) types.Paths {
	sort(paths)
	var result [][][]string
	for i, path := range paths.AllPaths {
		var ValidPaths [][]string
		ValidPaths = append(ValidPaths, path)
		result = append(result, ValidPaths)
		for j := i + 1; j < len(paths.AllPaths); j++ {
			if !checkIntercept(ValidPaths, paths.AllPaths[j]) {
				ValidPaths = append(ValidPaths, paths.AllPaths[j])
				result = append(result, ValidPaths)
			}
		}
	}
	return types.Paths{ValidPaths: result}
}

func checkIntercept(ValidPaths [][]string, path []string) bool {
	for _, paths := range ValidPaths {
		for _, room1 := range paths[:len(paths)-1] {
			for _, room2 := range path[:len(path)-1] {
				if room1 == room2 {
					return true
				}
			}
		}
	}
	return false
}

func sort(paths types.Paths) {
	for i := 0; i < len(paths.AllPaths); i++ {
		for j := i + 1; j < len(paths.AllPaths)-1; j++ {
			if len(paths.AllPaths[i]) > len(paths.AllPaths[j+1]) {
				paths.AllPaths[j+1], paths.AllPaths[i] = paths.AllPaths[i], paths.AllPaths[j+1]
			}
		}
	}
}
