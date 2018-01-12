/* marc_034.js */

window.addEventListener("load", function load(event){

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

			var raw = document.getElementById("raw");
			raw.innerHTML = "";

			var bboxes = document.getElementById("bboxes");
			bboxes.innerHTML = "";

			var on_success = function(rsp){
				console.log("OK", rsp);

				var str = JSON.stringify(rsp, null, 2);
				var pre = document.createElement("pre");
				pre.appendChild(document.createTextNode(str));

				var raw = document.getElementById("raw");
				raw.appendChild(pre);

				var bbox = rsp["bbox"];
				var minx = bbox[0];
				var miny = bbox[1];
				var maxx = bbox[2];
				var maxy = bbox[3];
				
				var coords = {
					"SW, NE": [ miny, minx, maxy, maxx ],
					"WS, EN": [ minx, miny, maxx, maxy ],
					"NE, SW": [ maxy, maxx, miny, minx ],
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

			var enc = encodeURIComponent(m);
			var url = location.protocol + "//" + location.host + "/bbox?034=" + enc;
			
			req.open("get", url, true);
			req.send();
			
			return false;
		}
	});