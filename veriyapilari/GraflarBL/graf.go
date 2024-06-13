package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Grafikteki düğüm sayısı
	vertices := 5

	// Bir grafik oluşturuluyor
	graph := NewGraph(vertices)

	// Grafiğe kenarlar ekleniyor
	graph.AddEdge(0, 1)
	graph.AddEdge(0, 2)
	graph.AddEdge(1, 3)
	graph.AddEdge(1, 4)
	graph.AddEdge(2, 4)

	// Düğüm 0'dan başlayarak BFS gezisi yapılıyor
	fmt.Print("Düğüm 0'dan başlayarak Önce Genişlik Araması: ")
	graph.BFS(0)
	fmt.Println()
	// Düğüm 0'dan başlayarak DFS gezisi yapılıyor
	fmt.Print("Düğüm 0'dan başlayarak Derinlik Öncelikli Arama: ")
	graph.DFS(0)
}

// Graph, bitişiklik listesi kullanılarak bir grafiği temsil eder
type Graph struct {
	vertices int          // Grafikteki düğüm sayısı
	adjList  []*list.List // Düğümlerin bitişiklik listeleri
}

// NewGraph, bir grafiği başlatır ve oluşturur
func NewGraph(vertices int) *Graph {
	g := &Graph{
		vertices: vertices,
		adjList:  make([]*list.List, vertices), // Grafikteki her düğümün bitişiklik listesi oluşturuluyor
	}

	// Her düğümün bitişiklik listesi için boş bir liste oluşturuluyor
	for i := range g.adjList {
		g.adjList[i] = list.New()
	}

	return g
}

// AddEdge, grafiğe bir kenar ekler
func (g *Graph) AddEdge(u, v int) {
	// u'dan v'ye bir kenar ekleniyor
	g.adjList[u].PushBack(v)
}

// BFS, grafikte Önce Genişlik Araması gerçekleştirir
func (g *Graph) BFS(startNode int) {
	visited := make([]bool, g.vertices) // Ziyaret edilmiş düğümleri tutmak için bir dizi oluşturuluyor

	queue := list.New() // BFS için bir kuyruk oluşturuluyor
	visited[startNode] = true
	queue.PushBack(startNode)

	// Kuyruk boş olana kadar BFS devam ediyor
	for queue.Len() != 0 {
		currentNode := queue.Remove(queue.Front()).(int) // Kuyruktan bir düğüm çıkarılıyor ve işleniyor
		fmt.Print(currentNode, " ")                      // Düğüm değeri yazdırılıyor

		// Düğümün komşuları üzerinde dolaşılıyor
		for e := g.adjList[currentNode].Front(); e != nil; e = e.Next() {
			neighbor := e.Value.(int) // Komşu düğüm değeri alınıyor
			if !visited[neighbor] {   // Eğer komşu ziyaret edilmemişse
				visited[neighbor] = true // Komşu ziyaret edildi olarak işaretleniyor
				queue.PushBack(neighbor) // Kuyruğa komşu ekleniyor
			}
		}
	}
}

// DFS, grafikte Derinlik Öncelikli Arama gerçekleştirir
func (g *Graph) DFS(startNode int) {
	visited := make([]bool, g.vertices) // Ziyaret edilmiş düğümleri tutmak için bir dizi oluşturuluyor

	// Derinlik Öncelikli Arama'yı gerçekleştiren yardımcı fonksiyon çağrılıyor
	g.DFSUtil(startNode, visited)
}

// DFSUtil, grafikte Derinlik Öncelikli Arama'nın yardımcı fonksiyonunu gerçekleştirir
func (g *Graph) DFSUtil(v int, visited []bool) {
	// Düğümü ziyaret edilmiş olarak işaretle ve yazdır
	visited[v] = true
	fmt.Print(v, " ")

	// Bu düğüme komşu olan tüm düğümleri tekrarla
	for e := g.adjList[v].Front(); e != nil; e = e.Next() {
		neighbor := e.Value.(int) // Komşu düğüm değeri alınıyor
		if !visited[neighbor] {
			g.DFSUtil(neighbor, visited) // Komşu düğümü ziyaret et
		}
	}
}

//ToDo Prim | kuruskal | Dijikstra | Belman-Ford yazılacak
