$("#new-publication").on("submit", createPublication);
$(".like-publication").on("click", likePublication);

function likePublication(event) {
    event.preventDefault();

    const clickedElement = $(event.target);

    const publicationID = clickedElement.closest("div").data("publication-id");

    clickedElement.prop("disabled", true);

    $.ajax({
        url: `/publications/${publicationID}/like`,
        method: "POST"
    }).done(function() {
        const likeCounter = clickedElement.next("span");

        const likes = parseInt(likeCounter.text())

        likeCounter.text(likes + 1);
    }).fail(function() {
        alert("Problem when liking publication")
    }).always(function() {
        clickedElement.prop("disabled", false);
    });
}

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