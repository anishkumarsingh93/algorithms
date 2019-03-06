package unionfind

import "fmt"

/*
	QuickFind is implemented using Eager approach.

	Data Structure:
	* Integer array []id of length n
	* Interpretation: p and q are connected iff(if and only if) they have the same id.

	index	0	1	2	3	4	5	6	7	8	9			0,5 and 6 are connected
	id[] 	0	1	1	8	8	0	0	1	8	8			1,2 and 7 are connected
															3,4,8 and 9 are connected

			0		1--2	 3--4
			|		   |	 |  |
			5--6	   7 	 8  9

	Find: pa dn q are connected if the ids of p and q are same.

	Union: For connecting pand q, we merge components containing p and q by changing
	all entries whose id equals id[p] to id[q].

*/

//QF is called for running the QuickFind algorithm
func QF() {

	var n int
	//Reading the length of array from console
	fmt.Println("Please enter the lenght of the array:")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error in reading the data:", err)
	}

	//Creating the []id array
	id := QuickFindUF(n)

	//Applying the Union
	fmt.Println("The initial array is: ", id)
	id = Union(id, 1, 5)
	id = Union(id, 2, 4)
	id = Union(id, 9, 8)
	id = Union(id, 3, 4)
	fmt.Println("Final array: ", id)

	//Checking if two elements are connected or not(in same component).
	fmt.Println("Connected(2,4): ", Connected(id, 2, 4))

}

//QuickFindUF initaites the list of array
func QuickFindUF(n int) []int {
	id := []int{}
	for i := 0; i < n; i++ {
		id = append(id, i)
	}
	return id
}

//Connected checks if the elements are connected(in the same component or not)
func Connected(id []int, p int, q int) bool {
	return id[p] == id[q]
}

//Union connects the two element so that they are in the same component
func Union(id []int, p int, q int) []int {
	/*
		First we store the pid and qid.
		Then we loop throught the ids and replace with qid for all those ids
		which is equal to the first element(pid).
	*/
	pid := id[p]
	qid := id[q]

	for i := 0; i < len(id); i++ {
		if id[i] == pid {
			id[i] = qid
		}
	}
	return id
}
