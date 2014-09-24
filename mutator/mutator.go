package mutator

import "github.com/criusmq/woodstock/graph"

type mutator interface{
  Mutate(g* graph.SimpleGraph)
}