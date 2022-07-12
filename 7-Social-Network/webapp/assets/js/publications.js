$("#new-publication").on("submit", createPublication);
$(document).on("click", ".like-publication", likePublication);
$(document).on("click", ".dislike-publication", dislikePublication);
$("#update-publication").on("click", updatePublication);
$(".delete-publication").on("click", deletePublication);

function deletePublication(event) {
    event.preventDefault();

    Swal.fire({
        title: "Alert!",
        text: "Are you sure you want to delete this publication? This action is irreversible.",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmed){
        if (!confirmed.value) return;
    
        const clickedElement = $(event.target);
        const publication = clickedElement.closest("div");
        const publicationID = publication.data("publication-id");

        clickedElement.prop("disabled", true);

        $.ajax({
            url: `/publications/${publicationID}`,
            method: "DELETE"
        }).done(function(){
            publication.fadeOut("slow", function(){
                $(this).remove();
            });
        }).fail(function(){
            Swal.fire("Oops...", "Error when deleting publication", "error");
        });
    });
}

function updatePublication() {
    $(this).prop("disabled", true);

    const publicationID = $(this).data("publication-id");
    
    $.ajax({
        url: `/publications/${publicationID}`,
        method: "PUT",
        data: {
            title: $("#title").val(),
            content: $("#content").val()
        }
    }).done(function(){
        Swal.fire(
            'Success!',
            'Publication updated successfully!',
            'success'
        ).then(function(){
            window.location = "/home";
        });
    }).fail(function(){
        Swal.fire("Oops...", "Error when updating publication", "error");
    }).always(function(){
        $("#update-publication").prop("disabled", false);
    });
}

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
        Swal.fire("Oops...", "Error when disliking publication", "error");
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
        Swal.fire("Oops...", "Error when liking publication", "error");
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
        Swal.fire("Oops...", "Error when creating publication", "error");
    });
}