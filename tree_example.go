package main

import (
	"fmt"
)

/*
            5
    3               8
 1      4       7       9
    2

5->3(left) and 8(right)
3->1(left) and 4(right)
1->2(right)
8->7(left) and 9(right)
*/

type Node struct {
	Left  *Node
	Value int
	Right *Node
}

func main() {
	fmt.Println("Tree Example")
	root := Node{nil, 5, nil}

	root.Left = &Node{nil, 3, nil}
	root.Right = &Node{nil, 8, nil}

	root.Left.Left = &Node{nil, 1, nil}
	root.Left.Right = &Node{nil, 4, nil}
	root.Right.Left = &Node{nil, 7, nil}
	root.Right.Right = &Node{nil, 9, nil}

	root.Left.Left.Right = &Node{nil, 2, nil}
	fmt.Println(root)

	// another way to setup the tree
	llr := &Node{nil, 2, nil}
	ll := &Node{nil, 1, llr}
	lr := &Node{nil, 4, nil}
	l := &Node{ll, 3, lr}

	rl := &Node{nil, 7, nil}
	rr := &Node{nil, 9, nil}
	r := &Node{rl, 8, rr}

	base := Node{l, 5, r}
	fmt.Println(base)
}
