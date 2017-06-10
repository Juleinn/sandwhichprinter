$(document).ready(function(){	
	
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
    	}
    });
});
