package main

import (
	"testing"
)

func Test_traverseGrid(t *testing.T) {
	//given
	terrane := []string{"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#"}
	travelPattern := [][]int{{3, 1}}
	travelPatternList := [][]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}

	t.Run("given the terrain and one travelPattern", func(t *testing.T) {
		var product int = 1
		for _, val := range travelPattern {
			numTrees := traverseGrid(terrane, val[0], val[1])
			product = product * numTrees
		}

		if product != 8 {
			t.Fatal("test Failed! the sum of trees encountered is incorrect")
		}
	})

	//when
	t.Run("given the terrain and a travelPattern list", func(t *testing.T) {
		var product int = 1
		for _, val := range travelPatternList {
			numTrees := traverseGrid(terrane, val[0], val[1])
			product = product * numTrees
		}

		if product != 336 {
			t.Fatal("test Failed! the product of trees encountered is incorrect")
		}
	})
}
