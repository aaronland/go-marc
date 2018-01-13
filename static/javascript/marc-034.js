/* marc_034.js */

window.addEventListener("load", function load(event){

    var map;
    
    if (L.Mapzen) {

	var api_key = document.body.getAttribute("data-mapzen-api-key");
	L.Mapzen.apiKey = api_key;
	
	var map_opts = { tangramOptions: {
	    scene: L.Mapzen.BasemapStyles.Refill
	}};
	
	map = L.Mapzen.map('map', map_opts);
	
	var sw = [ -55, -180 ];
	var ne = [ 55, 180 ];
	var bounds = [ sw, ne ];
	
	map.fitBounds(bounds);
    }

    else {

	var el = document.getElementById("map");
	el.innerText = "Maps are disabled because mapzen.js is not present.");
    }
    
    var m = document.getElementById("marc-034");
    m.value = "";
    
    var s = document.getElementById("submit");
    
    s.onclick = function(){

	var m = document.getElementById("marc-034");
	m = m.value;
	m = m.trim();
	
	if (m == ""){
	    return false;
	}
	
	var enc = encodeURIComponent(m);
	var url = location.protocol + "//" + location.host + "/bbox?034=" + enc;
	
	var raw = document.getElementById("raw");
	raw.innerHTML = "";
	
	var bboxes = document.getElementById("bboxes");
	bboxes.innerHTML = "";

	if (map){
	    var sw = [ -55, -180 ];
	    var ne = [ 55, 180 ];
	    var bounds = [ sw, ne ];
	    
	    map.fitBounds(bounds);
	}
	
	var on_success = function(rsp){
	    
	    var str = JSON.stringify(rsp, null, 2);
	    var pre = document.createElement("pre");
	    pre.appendChild(document.createTextNode(url));
	    pre.appendChild(document.createTextNode("\n\n"));				
	    pre.appendChild(document.createTextNode(str));
	    
	    var raw = document.getElementById("raw");
	    raw.appendChild(pre);
	    
	    var bbox = rsp["bbox"];
	    var minx = bbox[0];
	    var miny = bbox[1];
	    var maxx = bbox[2];
	    var maxy = bbox[3];
	    
	    var coords = {
		"S, W, N, E": [ miny, minx, maxy, maxx ],
		"W, S, E, N": [ minx, miny, maxx, maxy ],
		"N, E, S, W": [ maxy, maxx, miny, minx ],
	    }
	    
	    var ul = document.createElement("ul");
	    
	    for (var label in coords) {
		
		var bbox = coords[label];
		var str_bbox = bbox.join(",");
		
		var code = document.createElement("code");
		code.appendChild(document.createTextNode(str_bbox));
		
		var li = document.createElement("li");
		li.appendChild(document.createTextNode(label + " "));
		li.appendChild(code);
		
		ul.appendChild(li);
	    }
	    
	    var bboxes = document.getElementById("bboxes");
	    bboxes.appendChild(ul);

	    if (map){
		var sw = [ miny, minx ];
		var ne = [ maxy, maxx ];
		
		var bounds = [ sw, ne ];
		var opts = { padding: [50, 50] };
		
		map.fitBounds(bounds, opts);
		
		var layer = L.geoJSON(rsp);
		layer.addTo(map);
	    }
	};
	    
	var req = new XMLHttpRequest();
	
	req.onload = function(){
	    
	    try {
		var data = JSON.parse(this.responseText);
	    }
	    
	    catch (e){
		console.log("ERROR", e);
		return false;
	    }
	    
	    on_success(data);
	};
	
	req.open("get", url, true);
	req.send();
	
	return false;
    }
});
