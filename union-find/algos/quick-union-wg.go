package unionfind

import "fmt"

/*
	Weighted Quick-Union.
	Here the depth of the tree is reduced to max lg N which speeds up the process of finding Root

	Data Structure:
	* Integer array []id of length n
	* Interpretation: id[i] is the parent of i
	* Root of i is id[id[id[...id[i]...]]] 		<----	keep going until it doesn't change(id[i]==i)

	* An extra array having sz[i] having the size of the tree. This is used to decide which
	root should be a parent of which

	index	0	1	2	3	4	5	6	7	8	9
	id[] 	0	1	9	4	9	6	6	7	8	9
	sz[]	1	1	1	1	2	1	2	1	1	4

			0	1		9		6	7	8
					   / \		|
					  2   4     5
						  |
						  3

	Find: For connected(p,q), check if pand q have the same root

	Union: To connect p and q, Link the smaller tree to the root of larger tree
	and update the sz[] array of the root of the larger tree

	The root and connected func will be same as quick union.

*/

//QUWG is called for running the Weighted QuickUnion algorithm
func QUWG() {

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
	id, sz = UnionQUWG(id, sz, 1, 5)
	id, sz = UnionQUWG(id, sz, 2, 4)
	id, sz = UnionQUWG(id, sz, 9, 8)
	id, sz = UnionQUWG(id, sz, 3, 4)
	id, sz = UnionQUWG(id, sz, 1, 2)
	fmt.Println("Final array: ", id, "and its sizes: ", sz)

	//Checking if two elements are connected or not(in same component).
	fmt.Println("Connected(2,4): ", ConnectedQU(id, 2, 4))

}

//QuickUnionUFWG initaites the list of id and sz array
func QuickUnionUFWG(n int) ([]int, []int) {
	id := []int{}
	sz := []int{}
	for i := 0; i < n; i++ {
		id = append(id, i)
		sz = append(sz, 1)
	}
	return id, sz
}

//UnionQUWG connects the two element by making the root of the larger tree as
//the root of the smaller tree
func UnionQUWG(id []int, sz []int, p int, q int) ([]int, []int) {
	/*
		First we store the roots of p and q.
		Change the root of smaller tree to the root of larger tree and
		also update the size of the larger tree.
	*/
	i := Root(id, p)
	j := Root(id, q)

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
