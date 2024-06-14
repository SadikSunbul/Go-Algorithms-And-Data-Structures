package main

import "fmt"

type Node struct {
	key    int
	height int
	left   *Node
	right  *Node
}

func newNode(key int) *Node {
	return &Node{key: key, height: 1}
}

func height(N *Node) int {
	if N == nil {
		return 0
	}
	return N.height
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rightRotate(y *Node) *Node {
	x := y.left
	T2 := x.right

	// Döndürme işlemi
	x.right = y
	y.left = T2

	// Yükseklikleri güncelleme
	y.height = max(height(y.left), height(y.right)) + 1
	x.height = max(height(x.left), height(x.right)) + 1

	// Yeni kökü döndürme
	return x
}

func leftRotate(x *Node) *Node {
	y := x.right
	T2 := y.left

	// Döndürme işlemi
	y.left = x
	x.right = T2

	// Yükseklikleri güncelleme
	x.height = max(height(x.left), height(x.right)) + 1
	y.height = max(height(y.left), height(y.right)) + 1

	// Yeni kökü döndürme
	return y
}

func getBalance(N *Node) int {
	if N == nil {
		return 0
	}
	return height(N.left) - height(N.right)
}

func insert(node *Node, key int) *Node {
	// 1. Normal BST ekleme işlemini gerçekleştir
	if node == nil {
		return newNode(key)
	}

	if key < node.key {
		node.left = insert(node.left, key)
	} else if key > node.key {
		node.right = insert(node.right, key)
	} else {
		// Tekrarlayan anahtarlar izin verilmez
		return node
	}

	// 2. Bu üst düğümün yüksekliğini güncelle
	node.height = 1 + max(height(node.left), height(node.right))

	// 3. Bu üst düğümün dengesini kontrol etmek için bu düğümün denge faktörünü al
	balance := getBalance(node)

	// Eğer bu düğüm dengesiz hale gelirse, 4 durum vardır
	// Sol Sol Durumu
	if balance > 1 && key < node.left.key {
		return rightRotate(node)
	}

	// Sağ Sağ Durumu
	if balance < -1 && key > node.right.key {
		return leftRotate(node)
	}

	// Sol Sağ Durumu
	if balance > 1 && key > node.left.key {
		node.left = leftRotate(node.left)
		return rightRotate(node)
	}

	// Sağ Sol Durumu
	if balance < -1 && key < node.right.key {
		node.right = rightRotate(node.right)
		return leftRotate(node)
	}

	// Değişiklik yapılmayan düğümü döndür
	return node
}

func preOrder(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.key)
		preOrder(node.left)
		preOrder(node.right)
	}
}

func main() {
	var root *Node

	/* Verilen figürdeki ağacı oluşturma */
	root = insert(root, 10)
	root = insert(root, 20)
	root = insert(root, 30)
	root = insert(root, 40)
	root = insert(root, 50)
	root = insert(root, 25)

	/* Oluşturulan AVL Ağacının ön sıralı gezinimi şöyle olacaktır:
	        30
	       /  \
	     20   40
	    / \    \
	   10 25   50
	*/
	fmt.Print("Oluşturulan ağacın ön sıralı gezinimi: ")
	preOrder(root)
}
