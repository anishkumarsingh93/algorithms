package main

import (
	"fmt"
	unionfind "workspace/algorithms/union-find/algos"
)

/*
	This is the entrypoint of the union-find algorithms to be run.
	Please run this file and choose the options to run the corresponding algorithms.

*/

func main() {

	//Giving the user option to choose which algo to run.
	fmt.Println("Please select the algorithm to test:")
	fmt.Println("1. Quick-Find")
	fmt.Println("2. Quick-Union")
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error in reading the data:", err)
	}

	//Calling appropriate algos
	switch a {
	case 1:
		fmt.Println("Running QuickFind algo")
		unionfind.QF()
	case 2:
		fmt.Println("Running QuickUnion algo")
		unionfind.QU()
	case 3:
		fmt.Println("Running Weighted QuickUnion algo")
		unionfind.QUWG()
	default:
		fmt.Println("Please enter 1 or 2")
	}
}
