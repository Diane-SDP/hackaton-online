var ClientCountry = document.getElementById("clientselector")
var ClientCP = document.getElementById("ClientPostalCode")
var ClientVille = document.getElementById("ClientCity")
var ClientPays = document.getElementById("VilleClient")
var ClientVilleHfrance = document.getElementById("PaysClient")

ClientPays.style.display ="none"
ClientVilleHfrance.style.display ="none"


var StartCountry = document.getElementById("startselector")
var StartCP = document.getElementById("StartPostalCode")
var StartVille = document.getElementById("StartCity")
var StartPays = document.getElementById("VilleDepart")
var StartVilleHfrance = document.getElementById("PaysDepart")


StartPays.style.display ="none"
StartVilleHfrance.style.display ="none"

function StartChange(){
    if (StartCountry.value == "france"){
        StartCP.style.display = "inline-block"
        StartVille.style.display = "inline-block"
        StartPays.style.display = "none"
        StartPays.value = ""
        StartVilleHfrance.style.display = "none"
        StartVilleHfrance.value = ""
    }else {
        StartCP.style.display = "none"
        StartCP.value = ""
        StartVille.style.display = "none"
        StartVille.value = ""
        StartPays.style.display = "inline-block"
        StartVilleHfrance.style.display = "inline-block"
    }
}



function CLientChange(){
    if (ClientCountry.value == "france"){
        ClientCP.style.display = "inline-block"
        ClientVille.style.display = "inline-block"
        ClientPays.style.display = "none"
        ClientVilleHfrance.style.display = "none"
    }else {
        ClientCP.style.display = "none"
        ClientVille.style.display = "none"
        ClientPays.style.display = "inline-block"
        ClientVilleHfrance.style.display = "inline-block"
    }
}