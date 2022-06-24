$("#new-publication").on("submit", createPublication);

function createPublication(event) {
    event.preventDefault();

    $.ajax({
        url: "/publications",
        method: "POST",
        data: {
            title: $("#title").val(),
            content: $("#content").val()
        }
    }).done(function(){
        window.location = "/home";
    }).fail(function(){
        alert("Error when creating publication!");
    });
}