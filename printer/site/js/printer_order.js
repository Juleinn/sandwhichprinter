$(document).ready(function(){	
	
	$("#print_button").click(function(){
		topSlice = $("#topslice_selector").children("li.uk-active").text()
		topGarnish = $("#topgarnish_selector").children("li.uk-active").text()
		bottomGarnish = $("#bottomgarnish_selector").children("li.uk-active").text()
		bottomSlice	 = $("#bottomslice_selector").children("li.uk-active").text()
		if(confirm("Confirm sandwich : " + topSlice + "," + topGarnish + "," + bottomGarnish + "," + bottomSlice  + "?"))
		{
    		// generate data
    		var xmlDocument = $.parseXML("<sandwich></sandwich>")
    		
    		if(topSlice != "Nothing"){
    			var topSliceElem = xmlDocument.createElement('slice');
	    		topSliceElem.appendChild(document.createTextNode(topSlice));
    			xmlDocument.documentElement.appendChild(topSliceElem);
    		}
    		if(topGarnish != "Nothing"){
    			var topGarnishElem = xmlDocument.createElement('garnish');
	    		topGarnishElem.appendChild(document.createTextNode(topGarnish));
    			xmlDocument.documentElement.appendChild(topGarnishElem);
    		}
    		if(bottomGarnish != "Nothing"){
    			var bottomGarnishElem = xmlDocument.createElement('garnish');
	    		bottomGarnishElem.appendChild(document.createTextNode(bottomGarnish));
    			xmlDocument.documentElement.appendChild(bottomGarnishElem);
    		}
    		if(bottomSlice != "Nothing"){
    			var bottomSliceElem = xmlDocument.createElement('slice');
	    		bottomSliceElem.appendChild(document.createTextNode(bottomSlice));
    			xmlDocument.documentElement.appendChild(bottomSliceElem);
    		}

    		$.post("/print", (new XMLSerializer().serializeToString(xmlDocument.documentElement)));

    	}
    });
});