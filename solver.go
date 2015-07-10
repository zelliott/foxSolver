package main

import "fmt"

// Each tree is every permutation of steps for
// a given starting point.
type Tree struct {
	left  *Tree
	hole  int
	right *Tree
}

// This function actually generates our tree.
func generateTree(t *Tree, start int, steps int, holes int) {

	// Set hole value
	t.hole = start

	// If we still have steps
	if steps > 1 {
		steps--

		// Depending on which hole we are currently on, decide
		// which leaves to build and holes to traverse.
		if start == 1 {
			t.right = &Tree{nil, 0, nil}
			generateTree(t.right, start+1, steps, holes)
		} else if start == holes {
			t.left = &Tree{nil, 0, nil}
			generateTree(t.left, start-1, steps, holes)
		} else {
			t.right = &Tree{nil, 0, nil}
			t.left = &Tree{nil, 0, nil}
			generateTree(t.right, start+1, steps, holes)
			generateTree(t.left, start-1, steps, holes)
		}
	}
}

// This was just a helper function (now unused)
// that I was using to see if the tree was built properly.
func printTree(level int, t *Tree) {
	fmt.Printf("Level: %x, value: %x\n", level, t.hole)

	if t.left != nil {
		printTree(level+1, t.left)
	}

	if t.right != nil {
		printTree(level+1, t.right)
	}
}

// This function takes in a pointer to our tree, a pointer to
// our 2D array, and the current path we are building.  It
// traverses the given tree and builds paths as it does so.
// When each path is finished, it is added to our set of paths for that
// tree.
func generatePaths(t *Tree, tPaths *[][]int, path []int) {

	// Add the hole to our path
	path = append(path, t.hole)

	// If no more leaves in this tree, the path is done,
	// and we can add it to our tPaths.
	if t.left == nil && t.right == nil {
		*tPaths = append(*tPaths, path)
	}

	// If left or right leaves exist, continue traversing the tree.
	if t.left != nil {
		generatePaths(t.left, tPaths, path)
	}

	if t.right != nil {
		generatePaths(t.right, tPaths, path)
	}
}

func main() {

	// Set the number of holes
	holes := 5

	// Input your solution
	solution := []int{2, 2, 3, 3, 2, 2, 4, 4, 3, 3, 2, 2, 5, 5, 4, 4, 3, 3, 2, 2}

	// Here we store: (1) the number of steps in our solution
	// and (2) our eventual 2D array of path permutations.
	steps := len(solution)
	paths := make([][]int, 0)

	// For each starting position
	for i := 1; i <= holes; i++ {

		// Generate a tree representing all possible paths.
		var t Tree = Tree{nil, 0, nil}
		tPaths := make([][]int, 0)
		generateTree(&t, i, steps, holes)

		// Generate all possible paths from that tree and add them
		// to our master paths 2D array.
		generatePaths(&t, &tPaths, make([]int, 0))
		paths = append(paths, tPaths...)
	}

	// Boolean variable that tracks whether or not our solution is valid.
	isSolution := true

	// For each of the paths...
	fmt.Println(paths)
	for i := 0; i < len(paths); i++ {
		pathIsSolution := false

		// Travel along them with the solution and see if it is correct.
		for j := 0; j < len(paths[i]); j++ {
			if paths[i][j] == solution[j] {
				pathIsSolution = true
			}
		}

		isSolution = pathIsSolution

		// If the solution does not work for this path, print a failure message
		// and break out of the loop.
		if !pathIsSolution {
			fmt.Printf("The solution %v does not work for the path %v", solution, paths[i])
			break
		}
	}

	// Success!
	if isSolution {
		fmt.Printf("The solution %v works!  I only had to brute force in through %x permutations!", solution, len(paths))
	}
}
