require "pp"
require "matrix"
require "digest/sha1"

# TODO: Update the index on each modification of C & B 

class Matrix
  def add_column(v=0)

  end
end

def findInvariants(row_mat)
  mat = Matrix.rows(row_mat)

  m = phase1(mat)
  phase2(m)
end

def phase1(mat)
  # Line 1 :
  # Initialization
  n = nplaces = mat.row_count
  m = ntransitions = mat.column_count

  # Stats
  pCount = nplaces
  ppCount = 0
  mpCount = 0

  a = mat.clone
  c = mat.clone
  aPrime = mat.clone

  d = Matrix.I(nplaces)
  r = {}
  u = Matrix[]
  cRowIndex = Array.new(c.row_count) {|idx| 
    {:t=>:p,:i=>idx}
  }
  aPrimeRowIndex = cRowIndex.clone

  # Set of parallel Places
  qp = []
  # Line 2 :
  ntransitions.times do |i|
    
    # Detect and replace each parallel place with a representative parallel place

    # Make an index of the row hash
    parallelSetIndex = rowReverseIndex(c)
    # Line 3 : For each set of parallel places
    parallelSetIndex.each do |key,parallelPlaces|
      next if parallelPlaces.length <= 1

      # Line 4 :
      # if set contains a parallel place already
      parallelPlace = parallelPlaces.find{ |pIdx| cRowIndex[pIdx][:t] == :pp} 
      notParallelPlace = parallelPlaces.find_all{ |pIdx| cRowIndex[pIdx][:t] != :pp}
      if parallelPlace
        # Line 5 : Add Sj{ppk} \ ppk 
        r[parallelPlace] = r[parallelplace].concat[notParallelPlace]
      else
        # Line 6 :
        qp[ppCount] = {:t=>:pp,:i=>ppCount}
        parallelPlaces.each{|idx| r[ppCount] = [cRowIndex[idx]].concat(r[ppCount].to_a)}
        tempC = c.to_a
        tempC << tempC[parallelPlaces.first].clone
        c = Matrix.rows(tempC)
        cRowIndex << qp[ppCount]
        
        # Line 7 : add to D the new column to 0 and the ppk entry to 1
        tempD = d.to_a.map{|row| row.concat([0])}
        tempD[-1][-1] = 1
        d = Matrix.rows(tempD)

        # Line 8 : Append row to A'
        tempAPrime = aPrime.to_a
        tempAPrime << tempC[-1]
        aPrime = Matrix.rows(tempAPrime)
        aPrimeRowIndex << cRowIndex[-1]
        ppCount = ppCount +1
      end

      #TODO: Line 9 :
      
      # Line 10 : Delete Rows of B that correspond to {Sj \ ppk}
      tempC = c.to_a
      notParallelPlace.each{|placeIndex| tempC.delete_at(placeIndex); cRowIndex.delete_at(placeIndex)}
      c = Matrix.rows(tempC)
    end
  end
  {:aprime=>aPrime,:c=>c,:d=>d,:r=>r,:u=>u,:ci=>cRowIndex,:ai=>aPrimeRowIndex}  
end

def phase2(mat)
  mat
end

# Build a reverse index of the matrix based on each row.hash
def rowReverseIndex(mat)
  idx={}
  rows = mat.row_vectors
  rows.each_index do |i|
    hash = rows[i].hash
    idx[hash] = [i].concat(idx[hash].to_a)
  end
  return idx
end

simple_macro_to_parallel_places = [
  [1,-1,0,0],
  [0,1,0,-1],
  [1,0,-1,0],
  [0,0,1,-1]
]

simple_parallel_to_macro = [
  [1,-1,0],
  [1,-1,0],
  [0,1,-1],
]

simple_with_loop = [
  [1,-1,0,0],
  [0,1,0,-1],
  [1,0,-1,0],
  [0,0,1,-1]
]

# Passe 1
puts "Passe 1"
pp findInvariants(simple_parallel_to_macro)
# Passe 2
puts "Passe 2"
pp findInvariants(simple_macro_to_parallel_places)
