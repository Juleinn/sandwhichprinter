$(document).ready(function(){	
	
<<<<<<< HEAD
    $("#print_button").click(function(){
    	topSlice = $("#topslice_selector").children("li.uk-active").text()
    	topGarnish = $("#topgarnish_selector").children("li.uk-active").text()
    	bottomGarnish = $("#bottomgarnish_selector").children("li.uk-active").text()
    	bottomSlice	 = $("#bottomslice_selector").children("li.uk-active").text()
    	if(confirm("Confirm sandwich : " + topSlice + "," + topGarnish + "," + bottomGarnish + "," + bottomSlice  + "?"))
    	{
    		$.post("/print", {
    			"test sandwich": true
    		})
=======
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

>>>>>>> d0e3b9a22f34564b57a7ca07bfd3c9efb3ddb8ba
    	}
    });
});
