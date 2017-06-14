function updateSandwiches(){
  $.get("test.json", function(data, status){
    if(status == "success")
    {
      $("#sandwich_list").empty();
      $.each(data, function(i, e){
        $("#sandwich_list").append("<li>" + e.Name + " : " + e.content + "</li>");
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