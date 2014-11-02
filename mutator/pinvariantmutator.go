package mutator

import (
  "fmt"
  "github.com/criusmq/woodstock/graph"
)

type SparseMatrix struct {
  m map[int]map[int]int
  HighestColIndex int
}

func newSparseMatrix() *SparseMatrix {
  return &SparseMatrix{m: make(map[int]map[int]int)}
}
func (s *SparseMatrix) matrix() map[int]map[int]int { return s.m }
func (s *SparseMatrix) put(col int, line int, value int) {
  if s.m[col] == nil {
    s.m[col] = make(map[int]int)
  }
  
  if value == 0 && s.get(col,line) != 0{
    s.rm(col,line)
    return
  }
  s.m[col][line] = value
}

func (s *SparseMatrix) rm(col int, line int){
  if(s.m[col][line] != 0){
    delete(s.m[col],line)
  }
  if len(s.m[col]) == 0{
    delete(s.m,col)
  }
}

func (s *SparseMatrix) get(col int, line int) int {
  _,ok := s.m[col]
  if(ok){
    n,ok := s.m[col][line]
    if(ok){
      return n
    }
  }
  return 0
}

func (s *SparseMatrix) linearAddColumn(sourceCol int, destinationCol int){
  s.linearAddColumnWithCoefficients(sourceCol,1,destinationCol,1)
}


func (s *SparseMatrix) linearAddColumnWithCoefficients(sourceCol int, sourceCoef int, destinationCol int , destinationCoef int){
  for line,element := range s.m[sourceCol]{
    value := destinationCoef * s.get(destinationCol,line) + sourceCoef * element
    s.put(destinationCol,line,value)
  }
}

type PInvariantMutator struct{}

func (m *PInvariantMutator) Mutate(graph *graph.SimpleGraph) {
  // Using the books matrix method so we need to build an adjacency matrix

  bMatrix := newSparseMatrix()
  cMatrix := newSparseMatrix()

  // Phase 0
  initializeMatricesFromGraph(graph,bMatrix,cMatrix)
  // Phase 1
  for len(cMatrix.matrix()) > 0 {
    // 1.1
    rows := findRowsWithOnlyPosOrNegElements(cMatrix)
    // 1.1a
    deleteNonZeroColsFromRows(rows,bMatrix,cMatrix)
    // 1.1b
    rows = findRowsWithOnlyOnePosOrNegElement(cMatrix)
    // 1.1.b.1

    substituteColumnsWithLinearCombinationInRows(rows,bMatrix,cMatrix)
    // 1.1.b.2 coefficient replacement to eleminate columns
    columnEliminator(bMatrix,cMatrix)
  fmt.Printf("c %v\n",cMatrix)
  fmt.Printf("b %v\n",bMatrix)

  } // (end-for)

  // Phase 2
  minimizationOfInvariants(bMatrix)
  deleteNonMinimalSupportColumns(bMatrix)
  fmt.Printf("b %v\n",bMatrix)
}

func deleteNonMinimalSupportColumns(bMatrix *SparseMatrix)

func minimizationOfInvariants(bMatrix *SparseMatrix){

  for {
    row := rowWithNegativeElements(bMatrix)

    if(row == -1){ break }
    
    fmt.Println("loopinnn")

    posElemCols := make([]int,0)
    negElemCols := make([]int,0)

    for col,_ := range bMatrix.matrix(){
      v := bMatrix.get(col,row)
      if(v<0){negElemCols = append(negElemCols,col)}
      if(v>0){posElemCols = append(posElemCols,col)}
    } 

    if(len(posElemCols)!=0){
      for _,posCol := range posElemCols{
        for _,negCol := range negElemCols{
          newcol := make(map[int]int)
          
          m := bMatrix.matrix()
          
          m[bMatrix.HighestColIndex] = newcol
          newcolIndex := bMatrix.HighestColIndex + 1
          bMatrix.HighestColIndex = newcolIndex
          
          // HELLO --- linear combination
          // row,posCol
          a := bMatrix.get(negCol,row)
          b := bMatrix.get(posCol,row)
          bMatrix.linearAddColumnWithCoefficients(posCol,a,newcolIndex,1)
          bMatrix.linearAddColumnWithCoefficients(negCol,b,newcolIndex,1)
          
          thegcd := gcdOfMapElements( m[newcolIndex] )
          for _,e := range m[newcolIndex] {
            e = e/thegcd
          }
        }
      }
    }
    for _,col := range negElemCols{
      delete(bMatrix.matrix(),col)
    }
  } // (end-for)
  // TODO: DELETE FROM B ALL COLUMNS WITHOUT MINIMAL SUPPORT
}

func gcdOfMapElements(numbers map[int]int) int{

  firstRun := true
  lastGCD := 0

  for _,e := range numbers{
    if firstRun {
      firstRun = false
      lastGCD = e
    }else{
      lastGCD = gcd(lastGCD,e)    
    }
  }

  return lastGCD
}

func gcd(n1 int, n2 int) int{
  if(n2 ==0){
    return n1
  }
  return gcd(n2,n1%n2)
}

func rowWithNegativeElements(matrix *SparseMatrix) int {
  for _,rows:= range matrix.matrix(){
    for row,element:= range rows{
      if element < 0{
        return row
      }
    }
  }
  return -1
}

func abs(x int) int {
 if x < 0 {
 	return -x
 }
 return x
}

func columnEliminator(bMatrix *SparseMatrix, cMatrix *SparseMatrix){
  chosenrow:=-1 // h
  chosencol:=-1 // k
  for col , rows := range cMatrix.matrix(){
    for row, _ := range rows{
      chosencol=col
      chosenrow=row
      break
    }
    break
  }
  if(chosenrow != -1 && chosencol != -1){
    fmt.Println("OHOHOH! we got a winner")

    elementa := cMatrix.get(chosencol,chosenrow)

    for col,_ := range cMatrix.matrix(){
      elementb := cMatrix.get(col,chosenrow)
      if(col != chosencol && elementb != 0){
        fmt.Println("OHOHOH! we got a BIG winner")
        // replace col with with linear combination of col and chosencol
        // with coefficients
        a:=0
        b:=0

        if( (elementa > 0 && elementb > 0) || (elementa<0 && elementb<0) ){
          a = -abs(elementb)
          b = abs(elementa)
        }else{
          a = abs(elementb)
          b = abs(elementa)
        }

        cMatrix.linearAddColumnWithCoefficients(chosencol,a,col,b)
        bMatrix.linearAddColumnWithCoefficients(chosencol,a,col,b)
      }
    }
    delete(cMatrix.matrix(),chosencol)
    delete(bMatrix.matrix(),chosencol)
  }
}

func substituteColumnsWithLinearCombinationInRows(rows []int, bMatrix *SparseMatrix, cMatrix *SparseMatrix){

  positiveCols := make([]int,0)
  negativeCols := make([]int,0)
  // we need a map of the positive and negative cols so we can linear combine them
  counter := posNegCounter{npos:0,nneg:0}
  for _,row := range rows{
    // find the index of the column containing the element
    for col, column := range cMatrix.matrix(){
      element := column[row]
      if (element < 0){
        negativeCols = append(negativeCols,col)
        counter.nneg = counter.nneg +1}
      if (element > 0){
        positiveCols = append(positiveCols,col)
        counter.npos = counter.npos +1
      }
    }
    var originalCol int
    var combineCols []int
    // treatment on the 1 that is number 
    if(counter.npos == 1){
      originalCol = positiveCols[0]
      combineCols = negativeCols
    }else if(counter.nneg ==1){
      originalCol = negativeCols[0]
      combineCols = positiveCols
    }
  
    for _, col := range combineCols{
      mult := cMatrix.get(col,row)
      if(mult<0){mult = mult*-1}
      for i := 0; i < mult ; i++ {
        cMatrix.linearAddColumn(originalCol,col)
        bMatrix.linearAddColumn(originalCol,col)
      }
    }
    // delete the originating column used for the linear combination
    delete(cMatrix.matrix(),originalCol)
    delete(bMatrix.matrix(),originalCol)

  }
}

func  deleteNonZeroColsFromRows(rows []int,bMatrix *SparseMatrix,cMatrix *SparseMatrix){
  for _ , row := range rows {
    coltodel := make([]int,0)
    for col, colMap := range cMatrix.matrix() {
      if(colMap[row] != 0){
        coltodel = append(coltodel,col)
      }
    }
    for _, col := range coltodel{
      delete(cMatrix.matrix(),col)
      delete(bMatrix.matrix(),col)
    }
  }
}

type posNegCounter struct {npos int; nneg int}
func countPosAndNegElementsByRow(matrix *SparseMatrix) map[int]*posNegCounter{
  rowCount := make(map[int]*posNegCounter)

  for col, colMap := range matrix.matrix() {
    for row, _ := range colMap {

      n := matrix.get(col,row)
      _ , ok := rowCount[row]
      if(!ok){rowCount[row] =   &posNegCounter{npos:0,nneg:0}}

      if(n > 0){
        e:=rowCount[row]
        e.npos = e.npos+1
      }
      if(n < 0){
        e:=rowCount[row]
        e.nneg = e.nneg+1
      }
    }
  }
  return rowCount
}

func findRowsWithOnlyOnePosOrNegElement(cMatrix *SparseMatrix) []int {
  type dualcounter struct {npos int;nneg int}

  rowCount := countPosAndNegElementsByRow(cMatrix)

  retainedrows := make([]int,0)
  for row,count := range rowCount{
    if((count.npos == 1 && count.nneg > 0) ||
      (count.npos > 0 && count.nneg == 1)){
      retainedrows = append(retainedrows,row)
    }
  }
  return retainedrows
}

func findRowsWithOnlyPosOrNegElements(cMatrix *SparseMatrix) []int{
  type dualcounter struct {npos int;nneg int}

  rowCount := countPosAndNegElementsByRow(cMatrix)

  retainedrows := make([]int,0)
  for row,count := range rowCount{
    if((count.npos != 0 && count.nneg == 0) ||
      (count.npos == 0 && count.nneg != 0)){
      retainedrows = append(retainedrows,row)
    }
  }
  return retainedrows
}

func initializeMatricesFromGraph(g* graph.SimpleGraph,bMatrix *SparseMatrix, cMatrix *SparseMatrix){

  for _, vertex := range g.Vertices(){

    vertexattributes := vertex.Attributes()
    vertexType := vertexattributes["type"]

    if vertexType == "Place"{
      // Add into b matrix (identity matrix)
      bMatrix.put(vertex.Id(),vertex.Id(),1)
      if(bMatrix.HighestColIndex < vertex.Id()){
        bMatrix.HighestColIndex=vertex.Id()
      }
    }

    if vertexType == "Transition"{

      for _, edge := range vertex.Edges(){
        connectedVertices := edge.Vertices()
        multiplicity := 1
        multiplicity_i, ok := edge.Attributes()["multiplicity"]

        if ok {
          multiplicity = multiplicity_i.(int)
        }

        var connectionID int
        if connectedVertices[0].Id() == vertex.Id(){
          multiplicity = -multiplicity
          connectionID = connectedVertices[1].Id()
        }else{
          multiplicity = multiplicity
          connectionID = connectedVertices[0].Id()
        }
          cMatrix.put(connectionID,vertex.Id(),multiplicity)
          if(cMatrix.HighestColIndex < connectionID){
            cMatrix.HighestColIndex = connectionID
          }
      }
    }
  }
}
