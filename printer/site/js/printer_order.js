$(document).ready(function(){	
	/* Add the available elements to the page for selection */
	$.get("/available", function(data, status){
		if(status == "success"){
			available = JSON.parse(data)
			topSlice = $("#topslice_selector")
			topGarnish = $("#topgarnish_selector")
			bottomGarnish = $("#bottomgarnish_selector")
			bottomSlice = $("#bottomslice_selector")
			
			topSlice.empty();
			topGarnish.empty();
			bottomGarnish.empty();
			bottomSlice.empty();
	
					
			available.Slices.forEach(function(elem, index){
				topSlice.append("<li>" + elem.Name + "<br/><img src=\"" + elem.Img + "\"/></li>");
				bottomSlice.append("<li>" + elem.Name + "<br/><img src=\"" + elem.Img + "\"/></li>");
			});

			available.Garnishes.forEach(function(elem, index){
				topGarnish.append("<li>" + elem.Name + "<br/><img src=\"" + elem.Img + "\"/></li>");
				bottomGarnish.append("<li>" + elem.Name + "<br/><img src=\"" + elem.Img + "\"/></li>");
			});
			
			// always add the "Nothing choice"
			topSlice.append("<li>Nothing<br/><img src=\"static/empty_tiny.png\"/></li>");
			topGarnish.append("<li>Nothing<br/><img src=\"static/empty_tiny.png\"/></li>");
			bottomGarnish.append("<li>Nothing<br/><img src=\"static/empty_tiny.png\"/></li>");
			bottomSlice.append("<li>Nothing<br/><img src=\"static/empty_tiny.png\"/></li>");
			
			// add default choice
			topSlice.find(":first-child").addClass("uk-active");
			topGarnish.find(":first-child").addClass("uk-active");
			bottomSlice.find(":first-child").addClass("uk-active");
			bottomGarnish.find(":first-child").addClass("uk-active");


	

		} else {
			alert("Unable to get products available for printing. Please try again later");
		}
	});



	/* The following sends the print request according to the choice user made */
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

    		$.post("/print", (new XMLSerializer().serializeToString(xmlDocument.documentElement)), function(data, status){
				if(status == "success"){
					window.location.href = "/";
				} else {
					alert("Unable to print sandwich");
				}
			});

    	}
    });
});
