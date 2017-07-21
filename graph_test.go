package graph

import "testing"

func TestUndirectedInit(t *testing.T) {
	graph := &Graph{}

	if graph.nEdges != 0 {
		t.Errorf("graph.nEdges should be 0, got %v", graph.nEdges)
	}
	if graph.nVertices != 0 {
		t.Errorf("graph.nVertices should be 0, got %v", graph.nVertices)
	}

	edgeList := []int{1, 2, 1, 3, 3, 4, 4, 1, 4, 2}

	graph.Init(false, edgeList)

	if graph.nEdges != 5 {
		t.Errorf("graph.nEdges should be 5, got %v", graph.nEdges)
	}
	if graph.nVertices != 4 {
		t.Errorf("graph.nVertices should be 4, got %v", graph.nVertices)
	}
}

func TestDirectedInit(t *testing.T) {
	graph := &Graph{}

	if graph.nEdges != 0 {
		t.Errorf("graph.nEdges should be 0, got %v", graph.nEdges)
	}
	if graph.nVertices != 0 {
		t.Errorf("graph.nVertices should be 0, got %v", graph.nVertices)
	}

	edgeList := []int{1, 2, 1, 3, 3, 4, 4, 1, 4, 2}

	graph.Init(true, edgeList)

	if graph.nEdges != 5 {
		t.Errorf("graph.nEdges should be 5, got %v", graph.nEdges)
	}
	if graph.nVertices != 4 {
		t.Errorf("graph.nVertices should be 4, got %v", graph.nVertices)
	}
}

func TestBreadthFirstTraversal_Undirected(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
		3, 4,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	var order []int
	processVertexEarly := func(v int, data *TraversalData) {
		order = append(order, v)
	}
	processVertexLate := func(v int, data *TraversalData) {}
	processEdge := func(x int, y int, data *TraversalData) {}

	data := &TraversalData{}
	data.Init(graph)
	result := graph.BreadthFirstTraversal(1, processVertexEarly, processVertexLate, processEdge, data)

	for i := 1; i < graph.nVertices+1; i++ {
		if result.processed[i] != true {
			t.Errorf("result.processed["+string(i)+"] should be true, got %v", result.processed[i])
		}
	}

	if order[0] != 1 {
		t.Errorf("order[0] should be 1, got %v", order[0])
	}
	// 3 placed at the head of 1's edge list, since it was added after 2.
	if order[1] != 3 {
		t.Errorf("order[1] should be 3, got %v", order[1])
	}
	if order[2] != 2 {
		t.Errorf("order[2] should be 2, got %v", order[2])
	}
	if order[3] != 4 {
		t.Errorf("order[3] should be 4, got %v", order[3])
	}

	if result.parent[1] != 0 {
		t.Errorf("result.parent[1] should be 0, got %v", result.parent[1])
	}
	if result.parent[2] != 1 {
		t.Errorf("result.parent[2] should be 1, got %v", result.parent[2])
	}
	if result.parent[3] != 1 {
		t.Errorf("result.parent[3] should be 1, got %v", result.parent[3])
	}
	if result.parent[4] != 3 {
		t.Errorf("result.parent[4] should be 3, got %v", result.parent[4])
	}
}

func TestDepthFirstTraversal_Undirected(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
		3, 4,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	var order []int
	processVertexEarly := func(v int, data *TraversalData) {
		order = append(order, v)
	}
	processVertexLate := func(v int, data *TraversalData) {}
	processEdge := func(x int, y int, data *TraversalData) {}

	data := &TraversalData{}
	data.Init(graph)
	result := graph.DepthFirstTraversal(1, processVertexEarly, processVertexLate, processEdge, data)

	for i := 1; i < graph.nVertices+1; i++ {
		if result.processed[i] != true {
			t.Errorf("result.processed["+string(i)+"] should be true, got %v", result.processed[i])
		}
	}

	if order[0] != 1 {
		t.Errorf("order[0] should be 1, got %v", order[0])
	}
	// 3 placed at the head of 1's edge list, since it was added after 2.
	if order[1] != 3 {
		t.Errorf("order[1] should be 3, got %v", order[1])
	}
	if order[2] != 4 {
		t.Errorf("order[2] should be 4, got %v", order[2])
	}
	if order[3] != 2 {
		t.Errorf("order[3] should be 2, got %v", order[3])
	}

	if result.parent[1] != 0 {
		t.Errorf("result.parent[1] should be 0, got %v", result.parent[1])
	}
	if result.parent[2] != 1 {
		t.Errorf("result.parent[2] should be 1, got %v", result.parent[2])
	}
	if result.parent[3] != 1 {
		t.Errorf("result.parent[3] should be 1, got %v", result.parent[3])
	}
	if result.parent[4] != 3 {
		t.Errorf("result.parent[4] should be 3, got %v", result.parent[4])
	}
}

func TestBreadthFirstTraversal_Directed(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
		3, 4,
	}

	graph := &Graph{}
	graph.Init(true, edgeList)

	var order []int
	processVertexEarly := func(v int, data *TraversalData) {
		order = append(order, v)
	}
	processVertexLate := func(v int, data *TraversalData) {}
	processEdge := func(x int, y int, data *TraversalData) {}

	data := &TraversalData{}
	data.Init(graph)
	result := graph.BreadthFirstTraversal(1, processVertexEarly, processVertexLate, processEdge, data)

	for i := 1; i < graph.nVertices+1; i++ {
		if result.processed[i] != true {
			t.Errorf("result.processed["+string(i)+"] should be true, got %v", result.processed[i])
		}
	}

	if order[0] != 1 {
		t.Errorf("order[0] should be 1, got %v", order[0])
	}
	// 3 placed at the head of 1's edge list, since it was added after 2.
	if order[1] != 3 {
		t.Errorf("order[1] should be 3, got %v", order[1])
	}
	if order[2] != 2 {
		t.Errorf("order[2] should be 2, got %v", order[2])
	}
	if order[3] != 4 {
		t.Errorf("order[3] should be 4, got %v", order[3])
	}

	if result.parent[1] != 0 {
		t.Errorf("result.parent[1] should be 0, got %v", result.parent[1])
	}
	if result.parent[2] != 1 {
		t.Errorf("result.parent[2] should be 1, got %v", result.parent[2])
	}
	if result.parent[3] != 1 {
		t.Errorf("result.parent[3] should be 1, got %v", result.parent[3])
	}
	if result.parent[4] != 3 {
		t.Errorf("result.parent[4] should be 3, got %v", result.parent[4])
	}
}

func TestDepthFirstTraversal_Directed(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
		3, 4,
	}

	graph := &Graph{}
	graph.Init(true, edgeList)

	var order []int
	processVertexEarly := func(v int, data *TraversalData) {
		order = append(order, v)
	}
	processVertexLate := func(v int, data *TraversalData) {}
	processEdge := func(x int, y int, data *TraversalData) {}

	data := &TraversalData{}
	data.Init(graph)
	result := graph.DepthFirstTraversal(1, processVertexEarly, processVertexLate, processEdge, data)

	for i := 1; i < graph.nVertices+1; i++ {
		if result.processed[i] != true {
			t.Errorf("result.processed["+string(i)+"] should be true, got %v", result.processed[i])
		}
	}

	if order[0] != 1 {
		t.Errorf("order[0] should be 1, got %v", order[0])
	}
	// 3 placed at the head of 1's edge list, since it was added after 2.
	if order[1] != 3 {
		t.Errorf("order[1] should be 3, got %v", order[1])
	}
	if order[2] != 4 {
		t.Errorf("order[2] should be 4, got %v", order[2])
	}
	if order[3] != 2 {
		t.Errorf("order[3] should be 2, got %v", order[3])
	}

	if result.parent[1] != 0 {
		t.Errorf("result.parent[1] should be 0, got %v", result.parent[1])
	}
	if result.parent[2] != 1 {
		t.Errorf("result.parent[2] should be 1, got %v", result.parent[2])
	}
	if result.parent[3] != 1 {
		t.Errorf("result.parent[3] should be 1, got %v", result.parent[3])
	}
	if result.parent[4] != 3 {
		t.Errorf("result.parent[4] should be 3, got %v", result.parent[4])
	}
}

func TestConnectedComponents_oneComponent(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
		3, 4,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	if graph.ConnectedComponents() != 1 {
		t.Errorf("graph.ConnectedComponents() should be 1, got %v", graph.ConnectedComponents())
	}
}

func TestConnectedComponents_twoComponents(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
		3, 4,
		5, 6, // Component two.
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	if graph.ConnectedComponents() != 2 {
		t.Errorf("graph.ConnectedComponents() should be 2, got %v", graph.ConnectedComponents())
	}
}

func TestBipartite_true(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	if graph.Bipartite() != true {
		t.Errorf("graph.Bipartite() should be true, got %v", graph.Bipartite())
	}
}

func TestBipartite_false(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
		1, 3,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	if graph.Bipartite() != false {
		t.Errorf("graph.Bipartite() should be false, got %v", graph.Bipartite())
	}
}

func TestHasCycles_false(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	if graph.HasCycles() != false {
		t.Errorf("graph.HasCycles() should be false, got %v", graph.HasCycles())
	}
}

func TestHasCycles_true(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
		1, 3,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	if graph.HasCycles() != true {
		t.Errorf("graph.HasCycles() should be true, got %v", graph.HasCycles())
	}
}

func TestArticulationVertices_root(t *testing.T) {
	edgeList := []int{
		1, 2,
		1, 3,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	cutNodes := graph.ArticulationVertices()

	if len(cutNodes) != 1 {
		t.Errorf("len(cutNodes) should be 1, got %v", len(cutNodes))
	}
	if cutNodes[1] != "root" {
		t.Errorf("cutNodes[1] should be root, got %v", cutNodes[1])
	}
}

func TestArticulationVertices_parent(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
		3, 4,
		4, 2,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	cutNodes := graph.ArticulationVertices()

	if len(cutNodes) != 1 {
		t.Errorf("len(cutNodes) should be 1, got %v", len(cutNodes))
	}
	if cutNodes[2] != "parent" {
		t.Errorf("cutNodes[2] should be parent, got %v", cutNodes[2])
	}
}

func TestArticulationVertices_bridge(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
	}

	graph := &Graph{}
	graph.Init(false, edgeList)

	cutNodes := graph.ArticulationVertices()

	if len(cutNodes) != 1 {
		t.Errorf("len(cutNodes) should be 1, got %v", len(cutNodes))
	}
	if cutNodes[2] != "bridge" {
		t.Errorf("cutNodes[2] should be bridge, got %v", cutNodes[2])
	}
}

func TestTopologicalSort(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
		2, 4,
		4, 5,
	}

	graph := &Graph{}
	graph.Init(true, edgeList)

	sorted := graph.TopologicalSort()

	if sorted[0] != 5 {
		t.Errorf("sorted[0] should be 5, got %v", sorted[0])
	}
	if sorted[1] != 4 {
		t.Errorf("sorted[1] should be 4, got %v", sorted[1])
	}
	if sorted[2] != 3 {
		t.Errorf("sorted[2] should be 3, got %v", sorted[2])
	}
	if sorted[3] != 2 {
		t.Errorf("sorted[3] should be 2, got %v", sorted[3])
	}
	if sorted[4] != 1 {
		t.Errorf("sorted[4] should be 1, got %v", sorted[4])
	}
}

func TestStronglyConnectedComponents(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
		3, 4,
		4, 2,
	}

	graph := &Graph{}
	graph.Init(true, edgeList)

	componentCount, strongComponent := graph.StronglyConnectedComponents()

	if componentCount != 2 {
		t.Errorf("componentCount should be 2, got %v", componentCount)
	}

	if strongComponent[1] != 2 {
		t.Errorf("strongComponent[1] should be 2, got %v", strongComponent[1])
	}

	if strongComponent[2] != 1 {
		t.Errorf("strongComponent[2] should be 1, got %v", strongComponent[2])
	}

	if strongComponent[3] != 1 {
		t.Errorf("strongComponent[3] should be 1, got %v", strongComponent[3])
	}

	if strongComponent[4] != 1 {
		t.Errorf("strongComponent[4] should be 1, got %v", strongComponent[4])
	}

}

func TestStronglyConnectedComponents_withCrossEdge(t *testing.T) {
	edgeList := []int{
		1, 2,
		2, 3,
		3, 4,
		4, 2,
		1, 5,
		4, 5, // Cross edge to scc consisting only of vertex 5.
	}

	graph := &Graph{}
	graph.Init(true, edgeList)

	componentCount, strongComponent := graph.StronglyConnectedComponents()

	if componentCount != 3 {
		t.Errorf("componentCount should be 3, got %v", componentCount)
	}

	if strongComponent[1] != 3 {
		t.Errorf("strongComponent[1] should be 3, got %v", strongComponent[1])
	}

	if strongComponent[2] != 2 {
		t.Errorf("strongComponent[2] should be 2, got %v", strongComponent[2])
	}

	if strongComponent[3] != 2 {
		t.Errorf("strongComponent[3] should be 2, got %v", strongComponent[3])
	}

	if strongComponent[4] != 2 {
		t.Errorf("strongComponent[4] should be 2, got %v", strongComponent[4])
	}

	if strongComponent[5] != 1 {
		t.Errorf("strongComponent[5] should be 1, got %v", strongComponent[5])
	}

}
