function addUser() {
    let userNameValue = document.getElementById("playerName").value;
    let firstNameValue = document.getElementById("name").value;
    let lastNameValue = document.getElementById("lastName").value;
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8000/player", true);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
        playerName: userNameValue,
        name: firstNameValue,
        lastName: lastNameValue
    }));
}