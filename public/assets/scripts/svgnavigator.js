/*! 
 * SVGNavigator
 * 
 * Copyright (c) 2014 Pierre-Alexandre St-Jean
 * Released under the MIT license
 * https://raw.githubusercontent.com/pastjean/svgnavigator/master/LICENSE
 * 
 */

;(function(){
  "use strict";

  var defaults = {
    enablePan:  true,
    enableZoom: true
  }

  // Utility functions
  var utils = {}
  utils.isObject = function(obj) {
    var type = typeof obj;
    return type === 'function' || type === 'object' && !!obj;
  };
  utils.defaults = function(obj,defaults) {
    if (!utils.isObject(obj) || !utils.isObject(defaults)) return obj;
    for(var property in defaults){
      if(obj[property] === void 0) obj[property] = defaults[property];
    }
    return obj;
  }


  function SVGNavigator(svgElement,options){
    var state='none',
        stateOrigin;
    
    if(!utils.isObject(options)) options={};
    var options = utils.defaults(options,defaults)

    var r = svgElement.getBoundingClientRect()
    svgElement.setAttribute("viewBox","0 0 "+ r.width+ " " + r.height)

    function getSVGEventPoint(evt) {
      var p = svgElement.createSVGPoint();

      p.x = evt.clientX;
      p.y = evt.clientY;

      return p;
    }
    
    // The events
    function mouseWheelHandler(e){
      if(!options.enableZoom){ return };
      if(e.preventDefault){ e.preventDefault() };

      e.returnValue = false

      var delta;
      if(e.wheelDelta){
        delta = e.wheelDelta / 3600; // Chrome/Safari/IE
      }else{
        delta = e.detail / -90; // Mozilla/Opera
      }
      delta = -delta; // we need to go the right way
      var point = getSVGEventPoint(e);

      // var viewBox = svgElement.getAttribute("viewBox")
      var viewBox = svgElement.viewBox.baseVal
      // Zoom-in or Zoom-out
      viewBox.width   = viewBox.width  * (1 + delta)
      viewBox.height  = viewBox.height * (1 + delta)
      // Translate based on the zoom point
      viewBox.x = viewBox.x - (1+delta) * point.x + point.x
      viewBox.y = viewBox.y - (1+delta) * point.y + point.y
    }
    function mouseMoveHandler(e){
      if(e.preventDefault){ e.preventDefault() };
      e.returnValue = false;
      if(state === 'pan'  && options.enablePan){
        var p = getSVGEventPoint(e);

        var viewBox = svgElement.viewBox.baseVal
        viewBox.x = viewBox.x + (stateOrigin.x - p.x) * (viewBox.width/r.width)
        viewBox.y = viewBox.y + (stateOrigin.y - p.y) * (viewBox.height/r.height)
        stateOrigin = p
      }
    }
    function mouseDownHandler(e){
      if(e.preventDefault){ e.preventDefault() };
      e.returnValue = false;
      state = 'pan';
      stateOrigin = getSVGEventPoint(e);
    }
    function mouseUpHandler(e){
      if(e.preventDefault){ e.preventDefault() };
      e.returnValue = false;
      // Stop all actions on mouseup
      if(state === 'pan') state = '';
    }
    function mouseOutHandler(e){
      if(e.preventDefault){ e.preventDefault() };
      e.returnValue = false;
      // Stop all actions on mouseout
      if(state === 'pan') state = '';
    }

    // Setup handlers
    svgElement.addEventListener('mouseup',  mouseUpHandler,  false)
    svgElement.addEventListener('mousedown',mouseDownHandler,false)
    svgElement.addEventListener('mousemove',mouseMoveHandler,false)
    // when hovering a svg inner element the mouse out is triggered which we don't want
    // svgElement.addEventListener('mouseout', mouseOutHandler, false)

    // IE9, Chrome, Safari, Opera  
    svgElement.addEventListener('mousewheel',mouseWheelHandler,false)
    // Firefox
    svgElement.addEventListener('DOMMouseScroll',mouseWheelHandler,false)
  }

  window.SVGNavigator = SVGNavigator;

})();




