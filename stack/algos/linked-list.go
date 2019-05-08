package stack

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

//Client is the client code for the stack implementation
func Client() {
	fmt.Println("Enter the length of stack. This is only for stopping the function.")
	var l int
	fmt.Scan(&l)

	//Declaring the first node
	first := Node{}

	fmt.Println("Enter strings to push to stack and '-' to pop the stack.")

	//Procession the input
	for i := 0; i < l; i++ {

		reader := bufio.NewReader(os.Stdin)
		str, _ := reader.ReadString('\n')
		item := ""
		if strings.TrimRight(str, "\n") == "-" {
			if IsEmpty(first) {
				fmt.Println("Cannot Pop an empty Stack")
				return
			}
			first, item = Pop(first)
			fmt.Println("Popped:", item, "Now the first is:", first)

		} else {
			fmt.Println("To be pushed:", str)
			first = Push(str, first)
			fmt.Println(first)
		}
	}

	//Current Stack is
	n := first
	fmt.Println("********The Stack is *******", first)
	for n.next != nil {
		fmt.Println(n.item)
		n = *n.next
	}
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
