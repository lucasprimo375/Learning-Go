$("#create-form").on("submit", CreateUser);

function CreateUser(event) {
    event.preventDefault();
    
    if($("#password").val() != $("#confirm-password").val()){
        Swal.fire("Oops...", "Passwords are different!", "error");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            name: $("#name").val(),
            email: $("#email").val(),
            nick: $("#nick").val(),
            password: $("#password").val()
        }
    }).done(function() {
        Swal.fire("Success!", "User created successfully!", "success")
            .then(function(){
                $.ajax({
                    url: "/login",
                    method: "POST",
                    data: {
                        email: $("#email").val(),
                        password: $("#password").val()
                    }
                }).done(function(){
                    window.location = "/home";
                }).fail(function(){
                    Swal.fire("Oops...", "Error when authenticating user", "error");
                });
            });
    }).fail(function(err) {
        Swal.fire("Oops...", "Error when creating user", "error");
    });
}