require "pp"

require 'digest/sha1'


def copy(o)
  return Marshal.load(Marshal.dump(o))
end

def identity(n)
  Array.new(n){|i| Array.new(n){|j| i==j ? 1 : 0 }}
end

def hashline(line)
  Marshal.dump(line)
end

def identify(mat)
  mat.reduce([1,{}]){|m,n|
    i,h = m
    h.store("p#{i}",n)
    [i+1,h]
  }.last
end

def addcolumn(idmat,v=0)
  idmat.each{|key,value|
    value << v
  }
end

def findPinvariants(mat)

  n = nplaces = mat.length
  m = ntransitions = mat.first.length

  a = identify(mat)
  c = copy(a)
  d = identify(identity(n))
  aprime = copy(a)

  # parallel place look-up table is an hash of arrays
  r = {}
  # listing of parallel place numbers (faster operations)
  q = 0
  qm = 0

  # macro place look-up table is a matrix (array of arrays)
  u = []


  m.times do |i|
    # Detect and replace each parallel place set with a representative parallel place

    # First detection : build an index to get all parallel places in c
    # index is a dictionnary of {index:[linehash],lines:[]} 

    parallelsetindex = {}

    c.each do |key,value|
      hash = Digest::SHA1.hexdigest(value.join('--'))

      if !parallelsetindex[hash].nil?
        parallelsetindex[hash] << key 
      else
        parallelsetindex[hash] = [key]
      end
    end
    addedrepplace = false

    parallelsetindex.each do |key,values|
      addedrepplace = false
      # This is not a set of parallel places (more than one)
      if values.length <=1
        next
      end

      # if set contains a parallel place already
      if values.find{|x| /pp[0-9]+/ =~ x }
        # Add the set without the parallel place into R(ppk)
        parallelplace = values.find{|x| /pp[0-9]+/ =~ x }
        notpp = values.find_all{ |x| /^m?p[0-9]+$/ =~ x }
        r[parallelplace] = r[parallelplace].concat(notpp)
      else
        addedrepplace = true
        # New parallel place.
        # add a representative parallel place ppk to Qp and add the set to R(ppk)
        q = q+1
        ppk = "pp#{q}"
        r[ppk] = values

        ## Set line to 0
        # append a column
        # set new column line to 1
        # copy in aprime
        d[ppk] = d[values.first].map{|x| 0}
        addcolumn(d)
        d[ppk][-1] = 1
        c[ppk] = c[values.first]
        aprime[ppk] = c[ppk]
      end

      # TODO: ligne 9
      # Append to U 
      #  a) a zero column if a reprensentative parallel place has been created
      if addedrepplace
        u.each {|mplaces| mplaces << 0 }
      end
      #  b) the rows of D that correspond to the macro places in Sj
      mplacesInSj = values.find_all{|x|  /^mp[0-9]+$/ =~ x}
      mplacesInSj.each{|mplace|
        u << d[mplace]
      }

      # Delete rows of B that correspond to The set without ppk
      values.find_all{ |x| /^m?p[0-9]+$/ =~ x }.each do |key|
        c.delete key
        d.delete key
      end
    end
    # EXECUTE ONE ITERATION OF FM1

    # Count Positive and negative element of each cols
    colcount = []
    c.each do |k,values|
      values.each_index do |index|
        if colcount[index].nil?
          colcount[index] = {:n => 0, :p=> 0}
        end
        if values[index] > 0
          colcount[index][:p] = colcount[index][:p] + 1
        end
        if values[index] < 0
          colcount[index][:n] = colcount[index][:n] + 1
        end
      end
    end
    
    colcount = colcount.map do |v|
      v[:p] * v[:n]
    end
    # Get the first minimum column
    min = 0
    colcount.each_index do |index|
      if colcount[min] == 0 or ((colcount[index] != 0) and (colcount[index] < colcount[min]))
        min = index
      end
    end

    if colcount[min] > 0
      # choisir la colonne "min" faire une combinaison linéaire pour l'éliminer

      # append to B the rows resulting from positive linear combinations of 2 lines
      # row pairs in B that annihilate Colk(c)

      checked_pairs = {}
      linear_pairs = []

      c.each do |k1,v1|
        c.each do |k2,v2|

          next if k1 == k2
          next if checked_pairs[k1] == k2 or checked_pairs[k2] == k1
          checked_pairs[k1] = k2
          checked_pairs[k2] = k1
          
          # Les 0 n'affecte pas une combinaison linéaire
          next if v1[min] == 0 or v2[min] == 0
          # Les chiffres doivet être d'un signe différent
          next if ( v1[min] < 0 and v2[min] < 0) or  (v1[min] > 0 and v2[min] > 0)

          # on viens de trouver une pair trouvons le multiplicateur et ajoutons
          # le résultat dans la liste des nouvelles pair
          if v1[min].abs > v2[min].abs
            linear_pairs << {"k1"=>k1,"k2"=>k2,"a1"=>v1[min].abs / v2[min].abs,"a2"=>1}
          else
            linear_pairs << {"k1"=>k1,"k2"=>k2,"a1"=>1,"a2"=>v2[min].abs / v1[min].abs}
          end

        end
      end
      pp linear_pairs
      # Do the linear combination
      linear_pairs.each{|pair|
        qm = qm+1
        c1 = c[pair["k1"]].map{|e| e* pair["a1"]}
        d1 = d[pair["k1"]].map{|e| e* pair["a1"]}
        c2 = c[pair["k2"]].map{|e| e* pair["a2"]}
        d2 = d[pair["k2"]].map{|e| e* pair["a2"]}
        c["mp#{qm}"] = c1.zip(c2).map{|e| e[0]+e[1]}
        d["mp#{qm}"] = d1.zip(d2).map{|e| e[0]+e[1]}
      }
      # Line 13
      # Delete row which dont have 0 in the min column or that checksupport blabla > 1
      # TODO: Check_support
      keys_to_delete = []
      c.each{|k,v| keys_to_delete << k if v[min] != 0}
      keys_to_delete.each{|k| c.delete(k); d.delete(k)}
    end
    
  end
  #add qm 0 columns to d
  d.each{|k,v| v = v.concat(Array.new(qm,0))}


  {:aprime=>aprime,:c=>c,:d=>d,:r=>r,:u=>u}
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

# Passe 1
puts "Passe 1"
pp findPinvariants(simple_parallel_to_macro)
# Passe 2
puts "Passe 2"
pp findPinvariants(simple_macro_to_parallel_places)
