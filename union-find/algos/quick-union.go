package unionfind

import "fmt"

/*
	Quick-Union is implemented using lazy approach.
	A forest is made with many trees each having a root. There are not cyclic nature.

	Data Structure:
	* Integer array []id of length n
	* Interpretation: id[i] is the parent of i
	* Root of i is id[id[id[...id[i]...]]] 		<----	keep going until it doesn't change(id[i]==i)

	index	0	1	2	3	4	5	6	7	8	9
	id[] 	0	1	9	4	9	6	6	7	8	9

			0	1		9		6	7	8
					   / \		|
					  2   4     5
						  |
						  3

	Find: For connected(p,q), check if pand q have the same root

	Union: To connect p and q, we merge components copntaining p and q by setting
	the id of p's root to the id of q's root.

*/

//QU is called for running the QuickUnion algorithm
func QU() {

	var n int
	//Reading the length of array from console
	fmt.Println("Please enter the lenght of the array:")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error in reading the data:", err)
	}

	//Creating the []id array
	id := QuickUnionUF(n)

	//Applying the Union
	fmt.Println("The initial array is: ", id)
	id = UnionQU(id, 1, 5)
	id = UnionQU(id, 2, 4)
	id = UnionQU(id, 9, 8)
	id = UnionQU(id, 3, 4)
	id = UnionQU(id, 1, 2)
	fmt.Println("Final array: ", id)

	//Checking if two elements are connected or not(in same component).
	fmt.Println("Connected(2,4): ", ConnectedQU(id, 2, 4))

}

//QuickUnionUF initaites the list of array
func QuickUnionUF(n int) []int {
	id := []int{}
	for i := 0; i < n; i++ {
		id = append(id, i)
	}
	return id
}

//Root finds the root of a given value
func Root(id []int, i int) int {
	//We find the parent of the value recursively until the id and i are same
	for i != id[i] {
		i = id[i]
	}
	return i
}

//ConnectedQU checks if the elements are connected by checking if their roots are same
func ConnectedQU(id []int, p int, q int) bool {
	return Root(id, p) == Root(id, q)
}

//UnionQU connects the two element by making the root of the first as the root of the second
func UnionQU(id []int, p int, q int) []int {
	/*
		First we store the roots of p and q.
		Then we change the root of p to root of q
	*/
	i := Root(id, p)
	j := Root(id, q)
	id[i] = j
	return id
}
