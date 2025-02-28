var aaronland = aaronland || {};

aaronland.xhr = (function(){

    var self = {

	postFileAsBinaryData: function(url, file){

	    return new Promise((resolve, reject) => {

		const xhr = new XMLHttpRequest();
		xhr.open('POST', url, true);
		xhr.setRequestHeader('Content-Type', 'application/octet-stream');
		xhr.responseType = "blob";
		
		const reader = new FileReader();
		
		reader.onload = function(event) {
		    const arrayBuffer = event.target.result;
		    xhr.send(arrayBuffer);
		};
	    
		reader.onerror = function(err) {
		    console.error('Error reading file');
		    reject(err);
		};
		
		reader.readAsArrayBuffer(file);
		
		xhr.onload = function() {
		    
		    if (xhr.status >= 200 && xhr.status < 300) {
			console.debug('File uploaded successfully');
			resolve(xhr.response);
		    } else {
			console.error('Upload failed with status:', xhr.status);
			reject(xhr.status);
		    };
		    
		};
		
		xhr.onerror = function(err) {
		    console.error('Request error');
		    reject(err);
		};
		
	    });
	}
    };

    return self;
    
})();

