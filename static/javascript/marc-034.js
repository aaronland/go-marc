/* marc_034.js */

window.addEventListener("load", function load(event){

    var map;
    
    if (Tangram) {

	var api_key = document.body.getAttribute("data-nextzen-api-key");
	var style_url = document.body.getAttribute("data-nextzen-style-url");
	var tile_url = document.body.getAttribute("data-nextzen-tile-url");
	
	var map = L.map("map");    

	var tangramOptions = {
            scene: {
		import: [
                    style_url
		],
		sources: {
                    mapzen: {
			url: tile_url,
			// url_subdomains: ['a', 'b', 'c', 'd'],
			url_params: {api_key: api_key},
			tile_size: 512,
			max_zoom: 16
                    }
		}
            }
	};
	
	var tangramLayer = Tangram.leafletLayer(tangramOptions);
	tangramLayer.addTo(map);
		
	var sw = [ -55, -180 ];
	var ne = [ 55, 180 ];
	var bounds = [ sw, ne ];
	
	map.fitBounds(bounds);
    }

    else {

	var el = document.getElementById("map");
	el.innerText = "Maps are disabled because tangram.js is not present.";
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

	/*
	if (map){
	    var sw = [ -55, -180 ];
	    var ne = [ 55, 180 ];
	    var bounds = [ sw, ne ];

	    map.fitBounds(bounds);
	}
	*/
	
	var on_success = function(rsp){

	    var str = JSON.stringify(rsp, null, 2);
	    var pre = document.createElement("pre");
	    pre.appendChild(document.createTextNode(url));
	    pre.appendChild(document.createTextNode("\n\n"));				
	    pre.appendChild(document.createTextNode(str));
	    
	    var raw = document.getElementById("raw");
	    raw.appendChild(pre);
	    
	    var bbox = rsp["bbox"];
	    var minx = bbox[0].toFixed(6);
	    var miny = bbox[1].toFixed(6);
	    var maxx = bbox[2].toFixed(6);
	    var maxy = bbox[3].toFixed(6);

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
		var opts = { padding: [20, 20] };

		console.log("FIT BOUNDS", bounds)
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

	// console.log("FETCH", url);

	req.open("get", url, true);
	req.send();
	
	return false;
    }
});
