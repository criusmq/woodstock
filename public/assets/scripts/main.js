// WE need a websocket
//
// with snap.svg we create the graph

var s = Snap("#graph");

var bigCircle = s.circle(150, 150, 100);
bigCircle.attr({
    fill: "#bada55",
    stroke: "#000",
    strokeWidth: 5
});

//var exampleSocket = new WebSocket("ws://www.example.com/socketserver");


exampleSocket.onerror = function(event){
}
exampleSocket.onopen = function(event){
  exampleSocket.onmessage = function(event){
    var msg = JSON.parse(event.data)
  }
  exampleSocket.onclose = function(event){
  }
}
