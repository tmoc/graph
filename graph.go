package graph

import (
	"container/list"
	"log"
)

type edge struct {
	y    int // The connected vertex.
	next *edge
}

// Graph is a graph data structure.
type Graph struct {
	directed  bool
	edges     []*edge
	nEdges    int
	nVertices int
}

// TraversalData holds information about a traversal of a graph.
type TraversalData struct {
	discovered []bool
	processed  []bool
	parent     []int

	// Depth-first traversal only.
	time      int
	entryTime []int
	exitTime  []int
}

type processVertex func(int, *TraversalData)
type processEdge func(int, int, *TraversalData)

// EdgeClass is used to classify an edge.
type EdgeClass int

const (
	// TREE indicates a tree edge.
	TREE EdgeClass = iota
	// BACK indicates a back edge.
	BACK
	// FORWARD indicates a forward edge.
	FORWARD
	// CROSS indicates a cross edge.
	CROSS
)

// Adjust array size for ignored index 0 where necessary.
func adjustSize(size int) int {
	return size + 1
}

// Determines the class of a given edge.
func edgeClassification(x int, y int, data *TraversalData) EdgeClass {
	if data.parent[y] == x {
		return TREE
	}
	if data.discovered[y] && !data.processed[y] {
		return BACK
	}
	if data.processed[y] && (data.entryTime[y] > data.entryTime[x]) {
		return FORWARD
	}
	if data.processed[y] && (data.entryTime[y] < data.entryTime[x]) {
		return CROSS
	}
	log.Fatalf("Cannot classify edge: %d, %d", x, y)
	return -1
}

// Init initializes TraversalData for traversal.
func (data *TraversalData) Init(g *Graph) {
	size := adjustSize(g.nVertices)
	data.discovered = make([]bool, size)
	data.processed = make([]bool, size)
	data.parent = make([]int, size)
	data.entryTime = make([]int, size)
	data.exitTime = make([]int, size)
}

// Insert an edge into the graph. Part of initialization.
func (g *Graph) insertEdge(directed bool, x int, y int) {
	g.edges[x] = &edge{y: y, next: g.edges[x]} // Place node at the head.
	if directed {
		g.nEdges++
	} else {
		g.insertEdge(true, y, x)
	}
}

// Init initializes the graph.
func (g *Graph) Init(directed bool, edgeList []int) {
	g.directed = directed
	length := len(edgeList)
	keyMap := make(map[int]int)
	for i := 0; i < length; i++ {
		keyMap[edgeList[i]] = 1
	}
	g.nVertices = len(keyMap)
	g.edges = make([]*edge, adjustSize(g.nVertices))
	for xIndex := 0; xIndex < length-1; xIndex += 2 {
		yIndex := xIndex + 1
		x := edgeList[xIndex]
		y := edgeList[yIndex]
		g.insertEdge(directed, x, y)
	}
}

// BreadthFirstTraversal processes vertices in breadth-first order.
func (g *Graph) BreadthFirstTraversal(
	start int,
	pve processVertex, // Process vertex early.
	pvl processVertex, // Process vertex late.
	pe processEdge,
	data *TraversalData,
) *TraversalData {
	data.discovered[start] = true
	queue := list.New()
	queue.PushBack(start)

	for queue.Len() != 0 {
		currentVertex := queue.Remove(queue.Front()).(int)
		pve(currentVertex, data)
		data.processed[currentVertex] = true

		edgePointer := g.edges[currentVertex]

		for edgePointer != nil {
			y := edgePointer.y
			if g.directed == true || data.processed[y] == false {
				pe(currentVertex, y, data)
			}
			if data.discovered[y] == false {
				queue.PushBack(y)
				data.discovered[y] = true
				data.parent[y] = currentVertex
			}
			edgePointer = edgePointer.next
		}

		pvl(currentVertex, data)
	}

	return data
}

// Used by DepthFirstTraversal after initialization to recursively traverse
// vertices.
func (g *Graph) depthFirstTraverseVertex(
	v int,
	pve processVertex,
	pvl processVertex,
	pe processEdge,
	data *TraversalData,
) {
	data.discovered[v] = true
	pve(v, data)

	data.time++
	data.entryTime[v] = data.time

	edgePointer := g.edges[v]

	for edgePointer != nil {
		y := edgePointer.y
		if data.discovered[y] == false {
			data.parent[y] = v
			pe(v, y, data)
			g.depthFirstTraverseVertex(y, pve, pvl, pe, data)
			// The boolean expression for undirected graphs below is subtle.
			// y is either an ancestor or a descendant. The processed check rules out
			// descendant, and the parent check rules out parent ancestor. Only an edge to
			// a non-parent ancestor should be processed, which would indicate a back edge.
		} else if g.directed == true || (data.processed[y] == false && data.parent[v] != y) {
			pe(v, y, data)
		}
		edgePointer = edgePointer.next
	}

	data.processed[v] = true
	pvl(v, data)

	data.exitTime[v] = data.time
	data.time++
}

// DepthFirstTraversal processes vertices in depth-first order.
func (g *Graph) DepthFirstTraversal(
	start int,
	pve processVertex, // Process vertex early.
	pvl processVertex, // Process vertex late.
	pe processEdge,
	data *TraversalData,
) *TraversalData {
	g.depthFirstTraverseVertex(start, pve, pvl, pe, data)
	return data
}

// ConnectedComponents returns the number of connected components in the graph.
func (g *Graph) ConnectedComponents() int {
	if g.nVertices == 0 {
		return 0
	}

	data := &TraversalData{}
	data.Init(g)
	pve := func(v int, data *TraversalData) {}
	pvl := func(v int, data *TraversalData) {}
	pe := func(x int, y int, data *TraversalData) {}

	count := 0

	for i := 1; i <= g.nVertices; i++ {
		if data.discovered[i] == false {
			count++
			g.BreadthFirstTraversal(i, pve, pvl, pe, data)
		}
	}

	return count
}

// Bipartite checks if the graph is bipartite.
func (g *Graph) Bipartite() bool {
	if g.nVertices == 0 {
		return true
	}

	data := &TraversalData{}
	data.Init(g)
	pve := func(v int, data *TraversalData) {}
	pvl := func(v int, data *TraversalData) {}

	bipartite := true
	color := make([]int, adjustSize(g.nVertices)) // 0 is uncolored.

	pe := func(x int, y int, data *TraversalData) {
		if color[x] == color[y] {
			bipartite = false
		} else {
			if color[x] == 1 {
				color[y] = 2
			} else {
				color[y] = 1
			}
		}
	}

	for i := 1; i <= g.nVertices; i++ {
		if bipartite == false {
			break
		}
		if data.discovered[i] == false {
			color[i] = 1
			g.BreadthFirstTraversal(i, pve, pvl, pe, data)
		}
	}

	return bipartite
}

// HasCycles checks if the graph has any cycles
func (g *Graph) HasCycles() bool {
	if g.nVertices == 0 {
		return false
	}

	data := &TraversalData{}
	data.Init(g)
	pve := func(v int, data *TraversalData) {}
	pvl := func(v int, data *TraversalData) {}

	hasCycles := false

	pe := func(x int, y int, data *TraversalData) {
		if data.discovered[y] == true && data.parent[x] != y {
			hasCycles = true
		}
	}

	g.DepthFirstTraversal(1, pve, pvl, pe, data)

	return hasCycles
}

// ArticulationVertices returns a map of all cut-nodes.
func (g *Graph) ArticulationVertices(start int) map[int]string {
	cutNodes := make(map[int]string)
	if g.nVertices == 0 {
		return cutNodes
	}

	data := &TraversalData{}
	data.Init(g)

	// ancestor refers to the earliest reachable ancestor.
	ancestor := make([]int, adjustSize(g.nVertices))
	// outDegree is how many tree edges the vertex has to children. This is used
	// when determining if a vertex is a root cut-node or a child bridge cut-node.
	outDegree := make([]int, adjustSize(g.nVertices))

	// Initializes earliest reachable ancestor to self.
	pve := func(v int, data *TraversalData) {
		ancestor[v] = v
	}

	// Marks vertices as cut-nodes based on various properties. Also potentially
	// updates parent's earliest reachable ancestor.
	pvl := func(v int, data *TraversalData) {
		if data.parent[v] == 0 { // Check if root.
			if outDegree[v] > 1 { // Check if vertex has more than one child.
				cutNodes[v] = "root"
				return
			}
		}

		// If the parent is the root, neither of the below cut-node scenarios are
		// possible.
		parentIsRoot := data.parent[data.parent[v]] == 0

		if parentIsRoot == false {
			if ancestor[v] == data.parent[v] {
				cutNodes[data.parent[v]] = "parent"
			} else if ancestor[v] == v {
				cutNodes[data.parent[v]] = "bridge"
				if outDegree[v] > 0 {
					cutNodes[v] = "bridge"
				}
			}
		}

		timeV := data.entryTime[ancestor[v]]
		timeParent := data.entryTime[ancestor[data.parent[v]]]

		if timeV < timeParent {
			ancestor[data.parent[v]] = ancestor[v]
		}
	}

	// Updates either a vertex's out degree or earliest reachable ancestor value.
	// The only case where no update happens for an undirected graph is when y is
	// the parent of x. For directed graphs there is also no update in the case of
	// a forward or cross edge.
	pe := func(x int, y int, data *TraversalData) {
		class := edgeClassification(x, y, data)
		if TREE == class {
			outDegree[x]++
			return
		}
		if BACK == class && data.parent[x] != y {
			// Found back edge to ancestor y.
			if data.entryTime[y] < data.entryTime[ancestor[x]] {
				ancestor[x] = y
			}
		}
	}

	g.DepthFirstTraversal(start, pve, pvl, pe, data)

	return cutNodes
}

// TopologicalSort returns vertices sorted in topological order.
func (g *Graph) TopologicalSort() []int {
	if g.directed == false {
		log.Fatal("Cannot call TopologicalSort on an undirected graph.")
	}

	data := &TraversalData{}
	data.Init(g)

	sorted := make([]int, 0, g.nVertices)

	pve := func(v int, data *TraversalData) {}
	pvl := func(v int, data *TraversalData) {
		sorted = append(sorted, v)
	}
	pe := func(x int, y int, data *TraversalData) {
		if data.discovered[y] && !data.processed[y] {
			log.Fatalf("Back edge found: %d, %d. Cannot perform a topological sort on a graph with cycles.", x, y)
		}
	}

	for i := 1; i <= g.nVertices; i++ {
		if data.discovered[i] == false {
			g.DepthFirstTraversal(i, pve, pvl, pe, data)
		}
	}

	return sorted
}

// StronglyConnectedComponents returns data about strongly connected components
// in a directed graph.
func (g *Graph) StronglyConnectedComponents() (int, []int) {
	if g.directed == false {
		log.Fatal("Cannot call StronglyConnectedComponents on an undirected graph.")
	}

	data := &TraversalData{}
	data.Init(g)

	// The number of strongly connected components.
	componentCount := 0
	// Which strong component each vertex belongs to.
	scc := make([]int, adjustSize(g.nVertices))
	// The oldest vertex in the same component as each vertex. May be an ancestor
	// or a cousin vertex. Think of "low" as meaning "lowest entry time value".
	low := make([]int, adjustSize(g.nVertices))
	for i := 1; i <= g.nVertices; i++ {
		low[i] = i
	}

	// Stack for the current scc being built up.
	active := list.New()

	// Add newly discovered vertex to the current scc.
	pve := func(v int, data *TraversalData) {
		active.PushFront(v)
	}

	// This serves two purposes. The first is to check if we are done processing a
	// component. The second is to update parent's "low" value if necessary.
	pvl := func(v int, data *TraversalData) {
		// If v itself is the oldest reachable vertex in a scc when we are about to
		// back up from it, then this completes a scc.
		if low[v] == v {
			componentCount++
			scc[v] = componentCount
			current := active.Remove(active.Front()).(int)
			// Clear the stack and mark each vertex within as part of component.
			for current != v {
				scc[current] = componentCount
				current = active.Remove(active.Front()).(int)
			}
		}
		// Update parent's "low" value if necessary.
		if data.parent[v] > 0 { // Check if parent is root.
			if data.entryTime[low[v]] < data.entryTime[low[data.parent[v]]] {
				low[data.parent[v]] = low[v]
			}
		}
	}

	// Examine edges in order to update "low" value for x. Only back or cross
	// edges may impact this value.
	pe := func(x int, y int, data *TraversalData) {
		class := edgeClassification(x, y, data)
		if BACK == class {
			if data.entryTime[y] < data.entryTime[low[x]] {
				low[x] = y
			}
			return
		}
		if CROSS == class {
			if scc[y] == 0 {
				if data.entryTime[y] < data.entryTime[low[x]] {
					low[x] = y
				}
			}
		}
	}

	for i := 1; i <= g.nVertices; i++ {
		if data.discovered[i] == false {
			g.DepthFirstTraversal(i, pve, pvl, pe, data)
		}
	}

	return componentCount, scc
}
