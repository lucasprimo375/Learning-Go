$("#unfollow").on("click", unFollow);
$("#follow").on("click", follow);
$("#edit-user").on("submit", editUser);
$("#update-password").on("submit", updatePassword);
$("#delete-profile").on("click", DeleteProfile);

function DeleteProfile() {
    Swal.fire({
        title: "Atention!",
        text: "Are you sure you want to delete your profile permanently? This is an irreversible action!",
        showCancelButton: true,
        cancelButtonText: "Cancel",
        icon: "warning"
    }).then(function(confirmation){
        if (confirmation.value) {
            $.ajax({
                url: "/delete-profile",
                method: "DELETE"
            }).done(function(){
                Swal.fire("Success!", "Your profile was deleted successfully", "success")
                    .then(function(){
                        window.location = "/logout"
                    });
            }).fail(function(){
                Swal.fire("Ops...", "Error when deleting your profile", "error")
            });
        }
    });
}

function updatePassword(event) {
    event.preventDefault();

    if ($("#new-password").val() != $("#confirm-new-password").val()){
        Swal.fire("Ops...", "New passwords are not the same", "warning");
        return;
    }

    $.ajax({
        url: "/update-password",
        method: "POST",
        data: {
            current: $("#current-password").val(),
            new: $("#new-password").val()
        }
    }).done(function(){
        Swal.fire("Success!", "Password updated successfully!", "success")
            .then(function(){
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Ops...", "Error when updating password!", "error");
    });
}

function editUser(event) {
    event.preventDefault();

    $.ajax({
        url: "/edit-profile",
        method: "PUT",
        data: {
            name: $("#name").val(),
            email: $("#email").val(),
            nick: $("#nick").val()
        }
    }).done(function(){
        Swal.fire("Success!", "Profile updated successfully!", "success")
            .then(function(){
                window.location = "/profile";
            });
    }).fail(function() {
        Swal.fire("Ops...", "Error when updating profile!", "error");
    });
}

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