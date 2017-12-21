package main

import "fmt"

/*
	http://www.geeksforgeeks.org/red-black-tree-set-1-introduction-2/
 */

const RED = true
const BLACK = false

type RBTree struct {
	left  *RBTree
	right *RBTree
	key   int
	color bool
}

/*
Every node has a color either red or black
 */

//satisfy the requirement by representing the color with a bool

/*
Root of tree is always black
*/
func RootIsBlack(root *RBTree) bool { return root.color == BLACK }

/*
There are no two adjacent red nodes (A red node cannot have a red parent or red child
 */
func NoAdjacentRedNodes(root *RBTree) bool {
	if root == nil {
		return true
	}
	if root.color == RED {
		if root.left != nil {
			if root.left.color == RED {
				return false
			}
			if !NoAdjacentRedNodes(root.left) {
				return false
			}
		}
		if root.right != nil {
			if root.left.color == RED {
				return false
			}
			if !NoAdjacentRedNodes(root.right) {
				return false
			}
		}
	}
	return true
}

/*

Every path from root to a Null node has same number of black nodes
 */

func _EveryPathHasSameNumberOfBlackNodes(root *RBTree) int {
	if root == nil {
		return 0
	}
	left := _EveryPathHasSameNumberOfBlackNodes(root.left)
	right := _EveryPathHasSameNumberOfBlackNodes(root.right)
	if left != right {
		panic("every path from root has different number of black nodes")
	}
	if root.color == BLACK {
		return left + right + 1
	}
	return left + right
}

func EveryPathHasSameNumberOfBlackNodes(root *RBTree) (b bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered from panic!")
			b = false
		}
	}()
	_EveryPathHasSameNumberOfBlackNodes(root)
	return true
}

func IsRedBlackTree(root *RBTree) bool {
	return RootIsBlack(root) &&
		NoAdjacentRedNodes(root) &&
		EveryPathHasSameNumberOfBlackNodes(root)
}

func main() {
	tree := &RBTree{&RBTree{nil, nil, 1, BLACK},
		&RBTree{nil, nil, 7, RED}, 5, BLACK}

	fmt.Println(RootIsBlack(tree))
	fmt.Println(NoAdjacentRedNodes(tree))
	fmt.Println(EveryPathHasSameNumberOfBlackNodes(tree))
	fmt.Println(IsRedBlackTree(tree))
}
