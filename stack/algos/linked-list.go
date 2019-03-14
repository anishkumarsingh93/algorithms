package stack

/*
	We use linked-list to implement stack

		to --> or --> not --> to --> be --> null

	We maintain a pointer to the first node ins a linked-list;
	we insert/remove from the front.

	Pop()

		or --> not --> to --> be --> null

	1. Save the item to return
		item := first.item
	2. Delete the first node
		first = first.next
	3. Return saved item
		return item

	Push("that")

		that --> or --> not --> to --> be --> null

	1. Save a link to the list
		oldFirst := first
	2. Create a new Node and set the instance variables in the new node
		return Node{"that", &first}

*/

//Node is the node of the linked list
type Node struct {
	item string //Node is the value to be saved
	next *Node  //This is the pointer to the next Node
}

//IsEmpty checks whether the stack is empty of not
func IsEmpty(first Node) bool {
	return first.next == nil
}

//Push appends a value at the start
func Push(item string, first Node) Node {
	return Node{item, &first}
}

//Pop removes the value for the front
func Pop(first Node) (Node, string) {
	item := first.item
	first = *first.next
	return first, item
}
