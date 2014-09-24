package mutator

import( 
  "github.com/criusmq/woodstock/graph"
  "fmt"
)

type SparseMatrix struct{
  m map[int]map[int]int
}
func newSparseMatrix() *SparseMatrix{return &SparseMatrix{m:make(map[int]map[int]int)}}
func (s* SparseMatrix) inner() map[int]map[int]int {return s.m}
func (s* SparseMatrix) put(col int, line int, value int){
  if(s.m[col] == nil){s.m[col] = make(map[int]int)}
  
  s.m[col][line] = value
}

func (s* SparseMatrix) get(col int,line int) int{
  return s.m[col][line]
}
type PInvariantMutator struct{}

func (m* PInvariantMutator) Mutate(s *graph.SimpleGraph){
  // Using the books matrix method so we need to build an adjacency matrix
  

  vertices := s.Vertices()
  //edges := s.Edges()

  matrix := newSparseMatrix()
  
  transitions := make([]*graph.SimpleGraphVertex,0)
  
  for _,vertex := range vertices{
    var attributes = vertex.Attributes()
    var attrType = attributes["type"]
    if(attrType == "Place"){
      matrix.put(vertex.Id(),vertex.Id(),1)
      // La matrice identité de places créée en ajoutant les places
    }
    if(attrType == "Transition"){
      // La matrice des places/transitions est ajoutée en ajoutant les transitions
      // à la fin donc on les ajoutes dans une liste d'attente
      transitions = append(transitions,vertex)
    }  
  }
  
  /* for _,transition:= range transitions{ */
  /*   edges := transition.Edges() */
  /*   for _,edge := range edges { */
  /*     fmt.Println("%v",edge) */
  /*   } */
  /* } */
  
  
  
  fmt.Println(matrix)
}
