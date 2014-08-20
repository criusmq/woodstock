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
          var c = s.circle(vertex.attributes.graphics.X-vertex.attributes.graphics.Xoff+10,
                           vertex.attributes.graphics.Y-vertex.attributes.graphics.Yoff+10,10);
         c.attr({
           fill: "#bada55",
           stroke: "#000",
           strokeWidth: 5
         });
        break;
        case "Transition":
        var r = s.rect(vertex.attributes.graphics.X-vertex.attributes.graphics.Xoff,
                         vertex.attributes.graphics.Y-vertex.attributes.graphics.Yoff,
                         20,20);
         r.attr({
           fill: "#55daba",
           stroke: "#000",
           strokeWidth: 1
         });

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
