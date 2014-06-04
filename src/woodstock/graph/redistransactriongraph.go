package graph

type RedisTransactionGraph struct {
}

func NewRedisTransactionGraph() *RedisTransactionGraph {
	return &RedisTransactionGraph{}
}

func (g *RedisTransactionGraph) addNode(node *Node)                               {}
func (g *RedisTransactionGraph) addEdge(edge *Edge, fromNode *Node, toNode *Node) {}

func (g *RedisTransactionGraph) removeNode(node *Node) {}
func (g *RedisTransactionGraph) removeEdge(edge *Edge) {}

func (g *RedisTransactionGraph) Node(id int) *Node { return nil }
func (g *RedisTransactionGraph) Edge(id int) *Edge { return nil }
