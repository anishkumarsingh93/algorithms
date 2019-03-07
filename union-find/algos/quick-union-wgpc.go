package unionfind

import "fmt"

/*
	Weighted Quick-Union with Path Compression.

	Simple Path Implementation: Make every other node ih the path, while finding
	the root node, point to its grandparent(thereby halving path length)

	The best apporach will be Two-pass implementation: Add a second loop to Root() func
	to set the id[] of each examined node to the root.

*/

//QUWGPC is called for running the Weighted QuickUnion algorithm
func QUWGPC() {

	var n int
	//Reading the length of array from console
	fmt.Println("Please enter the lenght of the array(>10):")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error in reading the data:", err)
	}

	//Creating the []id array
	id, sz := QuickUnionUFWG(n)

	//Applying the Union
	fmt.Println("The initial array is: ", id)
	id, sz = UnionQUWGPC(id, sz, 1, 5)
	id, sz = UnionQUWGPC(id, sz, 2, 4)
	id, sz = UnionQUWGPC(id, sz, 9, 8)
	id, sz = UnionQUWGPC(id, sz, 3, 4)
	id, sz = UnionQUWGPC(id, sz, 1, 2)
	fmt.Println("Final array: ", id, "and its sizes: ", sz)

	//Checking if two elements are connected or not(in same component).
	fmt.Println("Connected(2,4): ", ConnectedQUPC(id, 2, 4))

}

//RootPC finds the root of a given value nas also compresses the path
func RootPC(id []int, i int) int {
	//We find the parent of the value recursively until the id and i are same
	for i != id[i] {
		//Simple path compression, pointing to grandparent
		id[i] = id[id[i]]

		i = id[i]
	}
	return i
}

//ConnectedQUPC checks if the elements are connected by checking if their roots are same
func ConnectedQUPC(id []int, p int, q int) bool {
	return RootPC(id, p) == RootPC(id, q)
}

//UnionQUWGPC connects the two element by making the root of the larger tree as
//the root of the smaller tree
func UnionQUWGPC(id []int, sz []int, p int, q int) ([]int, []int) {
	/*
		First we store the roots of p and q.
		Change the root of smaller tree to the root of larger tree and
		also update the size of the larger tree.
	*/
	i := RootPC(id, p)
	j := RootPC(id, q)

	//Checking if p and q are already connected.
	if i == j {
		return nil, id
	}

	//Changing the roost according to the size
	if sz[i] < sz[j] {
		id[i] = j
		sz[j] += sz[i]
	} else {
		id[j] = i
		sz[i] += sz[j]
	}

	return id, sz
}
