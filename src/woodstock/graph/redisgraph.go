package graph

type RedisGraph struct {
}

func NewRedisGraph() *RedisGraph {
	return &RedisGraph{}
}

func (g *RedisGraph) addNode(node *Node)                               {}
func (g *RedisGraph) addEdge(edge *Edge, fromNode *Node, toNode *Node) {}

func (g *RedisGraph) removeNode(node *Node) {}
func (g *RedisGraph) removeEdge(edge *Edge) {}

func (g *RedisGraph) Node(id int) *Node { return nil }
func (g *RedisGraph) Edge(id int) *Edge { return nil }
