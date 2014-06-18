package mutator

import "github.com/criusmq/woodstock/graph"

type mutator interface {
	mutate(g *graph.Graph)
}
