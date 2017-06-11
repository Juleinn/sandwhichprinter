function updateSandwiches(){
	$.get("/sandwiches", function(data, status){
		if(status == "success")
		{
			$("#sandwich_list").empty();
			data2 = JSON.parse(data);
			$.each(data2, function(i, e){
				var textSandwich = "";

				if(e.Slices != null){
					textSandwich += e.Slices[0];
				}

				for(var index in e.Garnishes){
					textSandwich += ", " + e.Garnishes[index] ;
				}

				if(e.Slices != null && e.Slices.length == 2){
					textSandwich += ", " + e.Slices[1];
				}

				$("#sandwich_list").append("<li>" + textSandwich + "</li>");

			})
		} else {
			alert("Update failed please try again later");
		}
	});


}

$(document).ready(function(){
		updateSandwiches();

		$("#reload_button").click(function(){
				updateSandwiches();
				});
		});
