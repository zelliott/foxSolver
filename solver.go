package main

import "fmt"

// Each tree is every permutation of steps for
// a given starting point.
type Tree struct {
  left  *Tree
  hole  int
  right *Tree
}

func generateTree(start, steps, holes)  {
  var t *Tree

  // For each step
  for i := 0; i < steps; i++ {

    // If we are on a hole at the end
    if start == 1 || start == holes {
      
    }
  }
}

func main()  {
  holes := 3
  solution := []int{1, 2, 3}
  steps := len(solution)
  trees := make(map[int]*Tree)

  for i := 1; i <= 3; i++ {
    trees[i] = generateTree(i, steps, holes)
  }
}
