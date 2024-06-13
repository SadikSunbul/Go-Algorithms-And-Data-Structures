package main

import "fmt"

//binary serch tree

func main() {

	binaryTree := TreeNode{}

	binaryTree.Add(4)
	binaryTree.Add(2)
	binaryTree.Add(1)
	binaryTree.Add(3)
	binaryTree.Add(8)
	binaryTree.Add(6)
	binaryTree.Add(9)
	binaryTree.Add(10)
	binaryTree.Add(5)

	binaryTree.Delete(Root, 4)

	fmt.Println("2 yi ara :", binaryTree.Search(2))
	fmt.Println("11 u ara :", binaryTree.Search(11))

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

func (tn *TreeNode) Delete(root *TreeNode, veri int) *TreeNode {

	if root == nil {
		return root
	}

	if root.Data > veri {
		//veri datadan kucuk sola gir
		root.Left = root.Left.Delete(root.Left, veri)
	} else if root.Data < veri {
		root.Right = root.Right.Delete(root.Right, veri)
	} else {

		//tek cocuk yada cocuksuz ıse
		if root.Left == nil { //kokun solu yok ıse sagı dondur
			return root.Right
		} else if root.Right == nil { //kokun sagı yok ıse solu dondur
			return root.Left
		}
		//silinen elemanın 2 cocugu var ıse
		root.Data = root.Right.FindMin(root.Right).Data
		root.Right = root.Right.Delete(root.Right, root.Data)
		if root.Data == veri { //burayı global olan Root derını degısıtrebılmek ıcın yazıldı
			Root = root
		}
	}
	return root
}

func (tn *TreeNode) FindMin(root *TreeNode) *TreeNode {
	current := root
	for current.Left != nil {
		current = current.Left
	}
	return current
}

func (tn *TreeNode) FindMax(root *TreeNode) *TreeNode {
	current := root
	for current.Right != nil {
		current = current.Right
	}
	return current
}
