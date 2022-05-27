$("#create-form").on("submit", CreateUser);

function CreateUser(event) {
    event.preventDefault()
    
    if ($("#password").val() != $("#confirm-password").val()) {
        alert("Passwords are different!");
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
    })
}