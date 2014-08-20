// WE need a websocket
//
// with snap.svg we create the graph

var s = Snap("#graph");
window.snap = s

// var exampleSocket = new WebSocket("ws://www.example.com/socketserver");


// exampleSocket.onerror = function(event){
// }
// exampleSocket.onopen = function(event){
//   exampleSocket.onmessage = function(event){
//     var msg = JSON.parse(event.data)
//   }
//   exampleSocket.onclose = function(event){
//   }
// }

window.graph ={}

var i=0
var showGraph = function(graphData){
  _(graphData.vertices).each(function(vertex){
    if(vertex.attributes!==undefined && vertex.attributes.type!==undefined){
      switch (vertex.attributes.type){
        case "Place":
        s.circle(100+i*70,100+i*70,50).attr({
    fill: "#bada55",
    stroke: "#000",
    strokeWidth: 5
});
        break;
        case "Transition":
          console.log(vertex)
        s.rect(50+i*70,50+i*70,100,100)
        break;
      }
    }
    i++
  })
}

var dataReq = new XMLHttpRequest();
dataReq.onload = function(){
  graphData = JSON.parse(this.responseText)
  window.graph = graphData
  showGraph(graphData);
}
dataReq.open("get","/graph.json",true)
dataReq.send()
