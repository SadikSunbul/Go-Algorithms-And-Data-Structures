package main

import "fmt"

// BTreeNode, B-ağacındaki bir düğümü temsil eder.
type BTreeNode struct {
	keys []int        // Anahtar dizisi
	C    []*BTreeNode // Çocuk işaretçilerinin dizisi
	n    int          // Mevcut anahtar sayısı
	leaf bool         // Düğümün yaprak olup olmadığını belirtir
	t    int          // Minimum derece (anahtar sayısı aralığını tanımlar)
}

// Yeni bir BTreeNode örneği oluşturur ve döndürür.
func newBTreeNode(t int, leaf bool) *BTreeNode {
	return &BTreeNode{
		keys: make([]int, 2*t-1),
		C:    make([]*BTreeNode, 2*t),
		n:    0,
		leaf: leaf,
		t:    t,
	}
}

// Bir dolu olmayan düğüme yeni bir anahtar ekler.
func (node *BTreeNode) insertNonFull(k int) {
	i := node.n - 1

	if node.leaf {
		// Anahtar, yaprağa eklensin.
		for i >= 0 && node.keys[i] > k {
			node.keys[i+1] = node.keys[i]
			i--
		}
		node.keys[i+1] = k
		node.n++
	} else {
		// Anahtar, yaprak olmayan bir düğüme eklenir.
		for i >= 0 && node.keys[i] > k {
			i--
		}

		if node.C[i+1].n == 2*node.t-1 {
			// Eğer çocuk dolu ise böl.
			node.splitChild(i+1, node.C[i+1])

			// Eğer yeni ortanca anahtar daha büyükse, ilerle.
			if node.keys[i+1] < k {
				i++
			}
		}
		node.C[i+1].insertNonFull(k)
	}
}

// Bir dolu çocuk düğümünü böler.
func (node *BTreeNode) splitChild(i int, y *BTreeNode) {
	z := newBTreeNode(y.t, y.leaf)
	z.n = y.t - 1

	// Y'nin ikinci yarısını z'ye kopyalar
	copy(z.keys, y.keys[y.t:])

	// Y yaprağı değilse, çocuk işaretçilerinin ikinci yarısını z'ye kopyalar
	if !y.leaf {
		copy(z.C, y.C[y.t:])
	}

	y.n = y.t - 1

	// Anahtarları ve çocuk işaretçilerini üst düğümde yeniden düzenler.
	for j := node.n; j > i; j-- {
		node.C[j+1] = node.C[j]
	}
	node.C[i+1] = z

	for j := node.n - 1; j >= i; j-- {
		node.keys[j+1] = node.keys[j]
	}
	node.keys[i] = y.keys[y.t-1]
	node.n++
}

// Bu düğümle kök olmak üzere alt ağacı tarama işlevi
func (node *BTreeNode) traverse() {
	i := 0
	for i < node.n {
		if !node.leaf {
			node.C[i].traverse()
		}
		fmt.Printf("%d ", node.keys[i])
		i++
	}
	if !node.leaf {
		node.C[i].traverse()
	}
}

// Bu düğümle kök olmak üzere alt ağaçta bir anahtar arama işlevi
func (node *BTreeNode) search(k int) *BTreeNode {
	i := 0
	for i < node.n && k > node.keys[i] {
		i++
	}
	if i < node.n && k == node.keys[i] {
		return node
	}
	if node.leaf {
		return nil
	}
	return node.C[i].search(k)
}

// BTree, B-ağacını temsil eder.
type BTree struct {
	root *BTreeNode // Kök düğüm işaretçisi
	t    int        // Minimum derece
}

// Yeni bir BTree örneği oluşturur ve döndürür.
func newBTree(t int) *BTree {
	return &BTree{
		root: nil,
		t:    t,
	}
}

// Ağacı taramak için bir işlev.
func (tree *BTree) traverse() {
	if tree.root != nil {
		tree.root.traverse()
	}
}

// Bu ağaçta bir anahtar aramak için bir işlev.
func (tree *BTree) search(k int) *BTreeNode {
	if tree.root == nil {
		return nil
	}
	return tree.root.search(k)
}

// Bu B-ağacına yeni bir anahtar eklemek için ana işlev.
func (tree *BTree) insert(k int) {
	if tree.root == nil {
		// Ağaç boşsa, yeni bir kök oluştur.
		tree.root = newBTreeNode(tree.t, true)
		tree.root.keys[0] = k
		tree.root.n = 1
	} else {
		if tree.root.n == 2*tree.t-1 {
			// Kök doluysa, yeni bir kök oluştur ve eski kökü böl.
			s := newBTreeNode(tree.t, false)
			s.C[0] = tree.root
			s.splitChild(0, tree.root)
			i := 0
			if s.keys[0] < k {
				i++
			}
			s.C[i].insertNonFull(k)
			tree.root = s
		} else {
			// Kök dolu değilse, köke ekle.
			tree.root.insertNonFull(k)
		}
	}
}

func main() {
	// Minimum derecesi 3 olan bir B-ağacı oluşturulur.
	t := newBTree(3)

	// Anahtarlar eklenir.
	t.insert(10)
	t.insert(20)
	t.insert(5)
	t.insert(6)
	t.insert(12)
	t.insert(30)
	t.insert(7)
	t.insert(17)

	fmt.Print("Oluşturulan ağacın taraması: ")
	t.traverse()
	fmt.Println()

	// Anahtarlar aranır.
	k := 6
	if t.search(k) != nil {
		fmt.Println("Mevcut")
	} else {
		fmt.Println("Mevcut Değil")
	}

	k = 15
	if t.search(k) != nil {
		fmt.Println("Mevcut")
	} else {
		fmt.Println("Mevcut Değil")
	}
}
