$("#new-publication").on("submit", createPublication);
$(document).on("click", ".like-publication", likePublication);
$(document).on("click", ".dislike-publication", dislikePublication);

function dislikePublication(event) {
    event.preventDefault();

    const clickedElement = $(event.target);

    const publicationID = clickedElement.closest("div").data("publication-id");

    clickedElement.prop("disabled", true);

    $.ajax({
        url: `/publications/${publicationID}/dislike`,
        method: "POST"
    }).done(function() {
        const likeCounter = clickedElement.next("span");

        const likes = parseInt(likeCounter.text())

        likeCounter.text(likes - 1);

        clickedElement.removeClass("dislike-publication")
        clickedElement.removeClass("text-danger")
        clickedElement.addClass("like-publication")
    }).fail(function() {
        alert("Problem when disliking publication")
    }).always(function() {
        clickedElement.prop("disabled", false);
    });
}

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

        clickedElement.addClass("dislike-publication")
        clickedElement.addClass("text-danger")
        clickedElement.removeClass("like-publication")
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