require 'SecureRandom'

class Edge
  attr_accessor :content,:vertices,:id
  def initialize
    self.content = {}
    self.vertices = Array.new(2)
  end
  
  def inspect
    "Edge:@id=#{self.id} @vertices=#{vertices.inspect} @content=#{content.inspect}"
  end
end

class Vertex
  attr_accessor :content,:edges,:id
  def initialize
    self.content = {}
    self.edges = {}
  end
  
  def inspect
    "Vertex:@id=#{self.id} @edges.count=#{edges.count} @content=#{content.inspect}"
  end
end

class Graph
  attr_accessor :vertices, :edges
  def initialize
    self.vertices = {}
    self.edges = {}
  end
  
  def add_edge!(edge,from_vertex,to_vertex)
    edge.id = SecureRandom.random_number(16**16)
    self.edges[edge.id] = edge
    
    edge.vertices[0] = from_vertex
    edge.vertices[1] = to_vertex

    from_vertex.edges[edge.id] = edge
    to_vertex.edges[edge.id] = edge
    
    self
  end
  
  def add_vertex!(vertex)
    vertex.id = SecureRandom.random_number(16**16)
    self.vertices[vertex.id] = vertex
    
    self
  end
  
  # When you remove a vertex it removes all associated edges
  def remove_vertex!(vertex)
        
    vertex.edges.each do |edge_id,edge|
      self.remove_edge!(edge)
    end
    
    vertices.delete(vertex.id)

    self
  end
      
  def remove_edge!(edge)
    edge.vertices.each do |vertex|
      vertex.edges.delete(edge.id)
    end
    edges.delete(edge.id)
    
    self
  end

  def inspect
    "Graph: vertices: #{vertices.count} edges:#{edges.count}"
  end
  
end

g = Graph.new
v1 = Vertex.new
v2 = Vertex.new
e1 = Edge.new
e2 = Edge.new

g.add_vertex!(v1).add_vertex!(v2).add_edge!(e1,v1,v2).add_edge!(e2,v2,v1).remove_vertex!(v1)