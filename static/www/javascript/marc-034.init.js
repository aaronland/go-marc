/* marc_034.js */

window.addEventListener("load", function load(event){

    // Null Island    
    var map = L.map('map').setView([0.0, 0.0], 1);    

    var init = function(cfg){

	var f = document.getElementById("feedback");
	
	var m = document.getElementById("marc-034");
	m.value = "";

	var e = document.getElementById("example-marc");

	e.onclick = function(){
	    m.value = e.innerText;
	    return false;
	};
	
	var u = document.getElementById("upload");

	u.onclick = function(){
	    
	    var f = document.getElementById("file");
	    const files = f.files;

	    if (files.length == 0){
		console.log("No files");
		return false;
	    }

	    f.innerHTML = "";
	    
	    const csv_f = files[0];

	    aaronland.xhr.postFileAsBinaryData("/convert", csv_f).then((rsp) => {

		var link = document.createElement('a');
		link.href = window.URL.createObjectURL(rsp);
		link.download = 'marc034-bbox.csv';
		
		document.body.appendChild(link);
		link.click();

		document.body.removeChild(link);
		
	    }).catch((err) => {
		f.innerText = "Failed to upload file, " + err;
		console.error("Failed to upload file", err);
	    });
	    
	    return false;
	};
	
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

	    f.innerHTML = "";
	    
	    var raw = document.getElementById("raw");
	    raw.innerHTML = "";
	    
	    var bboxes = document.getElementById("bboxes");
	    bboxes.innerHTML = "";
	    
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

		var label = document.createElement("label");
		label.appendChild(document.createTextNode("Bounding boxes"));
		
		var bboxes = document.getElementById("bboxes");
		bboxes.appendChild(label);
		bboxes.appendChild(ul);
		
		if (map){
		    
		    var sw = [ miny, minx ];
		    var ne = [ maxy, maxx ];
		    
		    var bounds = [ sw, ne ];
		    var opts = { padding: [20, 20] };
		    
		    map.fitBounds(bounds, opts);
		    
		    var layer = L.geoJSON(rsp);
		    layer.addTo(map);
		}

		// START OF intersects stuff
		// put me in a function or another package or something...
		
		var geom = rsp.geometry;
		
		var args = {
		    geometry: geom,
		    sort: [
			"placetype://",                                                                                                                                                            
			"name://",
			"inception://",
                    ],
		};

		var intersects_el = document.getElementById("intersects");
		intersects_el.innerHTML = "";
		
		const xhr = new XMLHttpRequest();
		xhr.open("POST", "/intersects", true);
		xhr.setRequestHeader("Content-Type", "application/json");

		xhr.onload = function(){

		    if (xhr.status != 200){

			intersect_el.innerText = "Unable to derive intersecting geometries";
			console.err("Failed to derive intersecting", req.response);
			return false;		    
		    }

		    try {
			var data = JSON.parse(xhr.responseText);

			/*
			var str = JSON.stringify(data, null, 2);
			var pre = document.createElement("pre");
			pre.appendChild(document.createTextNode(str));

			intersects_el.appendChild(pre);
			 */

			var places = data.places;
			var count_places = places.length;

			var table = document.createElement("table");
			table.setAttribute("class", "table");

			var caption = document.createElement("caption");
			caption.appendChild(document.createTextNode("Intersecting Who's On First places"));

			var thead = document.createElement("thead");
			var tbody = document.createElement("tbody");
			
			var tr = document.createElement("tr");

			var th_id = document.createElement("th");
			th_id.appendChild(document.createTextNode("ID"));

			var th_name = document.createElement("th");
			th_name.appendChild(document.createTextNode("Name"));

			var th_pt = document.createElement("th");
			th_pt.appendChild(document.createTextNode("Placetype"));

			tr.appendChild(th_id);
			tr.appendChild(th_name);
			tr.appendChild(th_pt);

			thead.appendChild(tr);
			table.appendChild(thead);
			
			for (var i=0; i < count_places; i++){

			    var pl = places[i];
			    var id = pl["wof:id"];
			    var name = pl["wof:name"];
			    var pt = pl["wof:placetype"];

			    var link = document.createElement("a");
			    link.setAttribute("href", "https://spelunker.whosonfirst.org/id/" + id);
			    link.setAttribute("target", "whosonfirst");
			    link.appendChild(document.createTextNode(id));

			    var tr = document.createElement("tr");

			    var td_id = document.createElement("td");
			    td_id.appendChild(link);

			    var td_name = document.createElement("td");
			    td_name.appendChild(document.createTextNode(name));

			    var td_pt = document.createElement("td");
			    td_pt.appendChild(document.createTextNode(pt));

			    tr.appendChild(td_id);
			    tr.appendChild(td_name);
			    tr.appendChild(td_pt);
			    tbody.appendChild(tr);
			}

			table.appendChild(tbody);

			var label = document.createElement("label");
			label.appendChild(document.createTextNode("Intersecting Who's On First places"));

			intersects_el.appendChild(label);
			intersects_el.appendChild(table);			
			
		    } catch (err) {

			intersect_el.innerText = "Failed to parse intersecting geometries";
			console.err("Failed to parse intersecting", err);
			
		    }
		};
		
		xhr.send(JSON.stringify(args));

		// END OF intersects stuff		
	    };
	    
	    var req = new XMLHttpRequest();
	    
	    req.onload = function(){

		if (req.status != 200){
		    f.innerText = "Failed to parse data, server returned error: " + req.response;
		    return false;		    
		}
		
		try {
		    var data = JSON.parse(this.responseText);
		}
		
		catch (e){
		    f.innerText = "Failed to parse data, " + e;
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
	
    };

    fetch("/map.json")
        .then((rsp) => rsp.json())
        .then((cfg) => {

            switch (cfg.provider) {
                case "leaflet":

                    var tile_url = cfg.tile_url;

                    var tile_layer = L.tileLayer(tile_url, {
                        maxZoom: 19,
                    });

                    tile_layer.addTo(map);
                    break;

                case "protomaps":

                    var tile_url = cfg.tile_url;

                    var tile_layer = protomapsL.leafletLayer({
                        url: tile_url,
                        theme: cfg.protomaps.theme,
                    })

                    tile_layer.addTo(map);
                    break;

                default:
                    console.error("Uknown or unsupported map provider");
                    return;
	    }

	    // To do: Set bounding box from configs (if defined)
	    
	    init(cfg);
	    
        }).catch((err) => {
	    console.error("Failed to derive map config", err);
	    return;
	});    
	    
});
