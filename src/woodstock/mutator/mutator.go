package mutator

import "woodstock/graph"

type mutator interface{
  mutate(g *graph.Graph)
}

