// WE need a websocket
//
// with snap.svg we create the graph
"use strict";

var globalVars = {}; // variables we want globally accessible can be stored in here

var s = Snap("#graph");

var i=0;

var clearGraph = function(){
  s.clear()
}

var showGraph = function(graphData, options){
  var edgeGroup = s.g(),
      vertexGroup = s.g();

  edgeGroup.attr({name:"edgegroup"});
  vertexGroup.attr({name:"vertexgroup"});
  
  _(graphData.vertices).each(function(vertex){
    if(vertex.attributes!==undefined && vertex.attributes.type!==undefined){

      var graphics = vertex.attributes.graphics,
      radius = 10,
      height = 20,
      width  = 20;
      switch (vertex.attributes.type){
        case "Place":
          var c = s.circle(graphics.X,
                           graphics.Y,radius);
                           c.attr({
                             fill: "#bada55",
                             stroke: "#000",
                             strokeWidth: 1
                           });
          var txt = s.text(vertex.attributes.namePosition.X,
                           vertex.attributes.namePosition.Y,
                           vertex.attributes.name)
             vertexGroup.add(c)
             vertex.shape = c
             break;

         case "Transition":
           var r = s.rect(graphics.X-width/2,
                          graphics.Y-width/2,
                          width,height);
                          r.attr({
                            fill: "#55daba",
                            stroke: "#000",
                            strokeWidth: 1
                          });
            vertexGroup.add(r)
            vertex.shape =r
            break;
      }
    }
    i++;
  });

  _(graphData.edges).each(function(edge){
    if(edge.attributes!==undefined && edge.attributes.graphics!==undefined){
      var graphics = edge.attributes.graphics;
      var pointsList = _.reduce(graphics.Points,function(memo,point){
        return _.flatten([memo,point.X,point.Y]);
      },[]);


      var line = s.polyline(pointsList);
      line.attr({fill:"none",stroke:"#000",
                strokeWidth:2});

                switch (edge.attributes.type){
                  case "Edge":
                    line.attr({class:"edge"})
                  break;
                  case "Read Edge":
                    line.attr({class:"read-edge"})
                  break;
                  default:
                    console.log(edge.attributes.type)
                }
                edgeGroup.add(line)
                edge.shape = line
    }
  });
};



function getNewGraph(){
  var dataReq = new XMLHttpRequest();
  dataReq.onload = function(){
    var graphData = JSON.parse(this.responseText);
    globalVars.data = graphData;
    clearGraph();
    showGraph(graphData);
  };
  dataReq.open("get","graph.json",true);
  dataReq.send();
}

function sendForm(form){
  var formData = new FormData(form);

  formData.append('secret_token', '1234567890'); // Append extra data before send.

  var xhr = new XMLHttpRequest();
  xhr.open('POST', form.action, true);
  xhr.onload = function(e) { getNewGraph() };

  xhr.send(formData);

  return false; // Prevent page from submitting.
}

window.addEventListener("load",function(){
  document.getElementById("graphFile").addEventListener('change',function(e){
    sendForm(document.getElementById("graphForm"))
  },false);

});

getNewGraph();

var g = document.getElementById("graph");
SVGNavigator(g);

globalVars.graph = s;

window.Woodstock = globalVars;
