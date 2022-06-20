function addMatch() {
    let checkboxes = document.querySelectorAll('input[name="playerId"]:checked');

    let values = [];
    checkboxes.forEach((checkbox) => {
        values.push(checkbox.value);
    });
    
    if (values.length == 4) {
        var xhr = new XMLHttpRequest();
        xhr.open("POST", "http://localhost:8000/match", true);
        xhr.setRequestHeader('Content-Type', 'application/json');
        xhr.send(JSON.stringify({
            "playerOne": parseInt(values[0]),
            "playerTwo": parseInt(values[1]),
            "playerThree": parseInt(values[2]),
            "playerFour": parseInt(values[3]),
            "status": "pending",
            "result": "pending"
        }));
    } else {
        alert("You have to select only 4 players")
    }
}

function updateMatch(match) {
    const url = "http://localhost:8000/match/" + match.matchId;
    fetch(url, {
        method : "PUT",
        body : JSON.stringify({
            "playerOne": match.playerOne,
            "playerTwo": match.playerTwo,
            "playerThree": match.playerThree,
            "playerFour": match.playerFour,
            "status": "DONE",
            "result": "pending"
        })
    }).then(
        response => response.text()
    ).then(
        html => console.log(html)
    );
}