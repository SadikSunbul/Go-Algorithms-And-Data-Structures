package main

import "fmt"

//binary serch tree

func main() {

	binaryTree := TreeNode{Data: 5, Left: nil, Right: nil}

	binaryTree.Add(4)
	binaryTree.Add(2)
	binaryTree.Add(1)
	binaryTree.Add(3)
	binaryTree.Add(8)
	binaryTree.Add(6)
	binaryTree.Add(9)

	fmt.Println("2 yi ara :", binaryTree.Search(2))
	fmt.Println("10 u ara :", binaryTree.Search(10))

	fmt.Print()
}

var Root *TreeNode

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func (tn *TreeNode) Add(veri int) {
	newData := TreeNode{Data: veri, Left: nil, Right: nil}

	if Root == nil {
		Root = &newData
		return
	}

	tempData := Root

	for tempData != nil {
		if tempData.Data <= veri {
			if tempData.Right == nil {
				tempData.Right = &newData
				return
			}
			tempData = tempData.Right
		} else {
			if tempData.Left == nil {
				tempData.Left = &newData
				return
			}
			tempData = tempData.Left
		}
	}

}

// minValue, belirli bir düğüm alt ağacında en küçük değeri bulur.
func (tn *TreeNode) minValue() int {
	min := tn.Data
	for tn.Left != nil {
		min = tn.Left.Data
		tn = tn.Left
	}
	return min
}

func (tn *TreeNode) Search(veri int) *TreeNode {

	temp := Root

	for temp != nil {
		if temp.Data < veri {
			temp = temp.Right
		} else if temp.Data > veri {
			temp = temp.Left
		} else {
			return temp
		}
	}
	return nil
}

//ToDo Delete kodu yazılıcak
