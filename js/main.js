document.addEventListener("DOMContentLoaded", function(_event) {
    updateUserBackend()
});

var curretntUser = null

async function updateUserBackend() {
    const responce = await fetch("/user")
    const data = await responce.text()

    var root = document.getElementById('root')
    root.innerHTML = data
}

async function updateUserFrontend() {
    const data = await fetch("https://randomuser.me/api").then(function(response) { return response.json(); })
     
    var name = document.getElementById("name-string")
    var email = document.getElementById("email-string")
    var phone = document.getElementById("phone-string")
    var picture = document.getElementById("picture-url")
    var picture_box = document.getElementById("img-box")
    var info_seed = document.getElementById("seed-string")
    var api_version = document.getElementById("api-version-string")

    name.innerHTML = `${data.results[0].name.first} ${data.results[0].name.last}`
    email.innerText = data.results[0].email
    phone.innerText = data.results[0].cell
    picture.alt = name.innerHTML
    picture.src = data.results[0].picture.large
    picture_box.title = name.innerHTML
    info_seed.innerText = data.info.seed
    api_version.innerText = data.info.version

    var gender = "male"
    var opposite_gender = "female"

    if (picture_box.classList.contains(opposite_gender)) {
        [gender, opposite_gender] = [opposite_gender, gender]
    }

    if (gender != data.results[0].gender) {
        var nodes = document.querySelectorAll(`.${gender}`)
        
        nodes.forEach(function(currentValue, _currentIndex, _listObj) {
            currentValue.classList.remove(gender)
            currentValue.classList.add(opposite_gender)
        })    
    }
}
