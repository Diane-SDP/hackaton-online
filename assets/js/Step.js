function PlusStep(button){
    console.log(button.value)
    url = '/BO'
    var data = {
        type:"plus",
        mess: "",
        id : button.value,
    }
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => {
        location.reload()
    })
    .catch(error => {
        console.error('Erreur lors de l\'envoi de la requête:', error);
    });
    
}
function MoinsStep(button){
    console.log(button.value)
    url = '/BO'
    var data = {
        type:"moins",
        mess: "",
        id : button.value,
    }
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .then(response => {
        location.reload()
    })
    .catch(error => {
        console.error('Erreur lors de l\'envoi de la requête:', error);
    });
    
}