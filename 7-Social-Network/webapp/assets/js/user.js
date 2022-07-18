$("#unfollow").on("click", unFollow);
$("#follow").on("click", follow);

function unFollow(event) {
    event.preventDefault();
    const userID = $(this).data("user-id")
    $(this).prop("disabled", true)

    $.ajax({
        url: `/users/${userID}/unFollow`,
        method: "POST"
    }).done(function(){
        window.location = `/users/${userID}`;
    }).fail(function(){
        Swal.fire("Ops...", "Error when trying to unfollow user!", "error");
        $("#unfollow").prop("disabled", false)
    });
}

function follow(event) {
    event.preventDefault();
    const userID = $(this).data("user-id");
    $(this).prop("disabled", true);

    $.ajax({
        url: `/users/${userID}/Follow`,
        method: "POST"
    }).done(function(){
        window.location = `/users/${userID}`;
    }).fail(function(err){
        console.log(err);
        Swal.fire("Ops...", "Error when trying to follow user!", "error");
        $("#follow").prop("disabled", false);
    });
}