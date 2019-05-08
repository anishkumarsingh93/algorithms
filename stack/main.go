package main

import (
	"fmt"

	stack "github.com/anishkumarsingh93/algorithms/stack/algos"
)

/*
	This is the entrypoint of the algorithms which implements stack.
	Please run this file and choose the options to run the corresponding algorithms.

*/

func main() {

	//Giving the user option to choose which algo to run.
	fmt.Println("Please select the algorithm to test:")
	fmt.Println("1. Stack linked-list")
	//fmt.Println("2. Quick-Union")
	//fmt.Println("3. Weighted Quick-Union")
	//fmt.Println("4. Weighted Quick-Union Path Compression")
	var a int
	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Error in reading the data:", err)
	}

	//Calling appropriate algos
	switch a {
	case 1:
		fmt.Println("Running Stack Linked-list algo")
		stack.Client()
	/*case 2:
		fmt.Println("Running QuickUnion algo")
		unionfind.QU()
	case 3:
		fmt.Println("Running Weighted QuickUnion algo")
		unionfind.QUWG()
	case 4:
		fmt.Println("Running Weighted QuickUnion algo")
		unionfind.QUWGPC()
	*/
	default:
		fmt.Println("Please enter 1, 2 3 or 4")
	}
}
