console.log("category.js loaded");

document.addEventListener("DOMContentLoaded", function () {

    console.log("DOM Ready");

    const dateInput = document.getElementById("category-date");

    console.log(dateInput);

    if (!dateInput) {
        return;
    }

    dateInput.addEventListener("click", function () {

        console.log("clicked");

        openDateModal("category", this);

    });

});

document.addEventListener("DOMContentLoaded", function () {

    const guestField = document.getElementById("guest-field");

    guestField.addEventListener("click", function (e) {

        e.stopPropagation();

        toggleGuestPopup();

    });

});


const destinationInput =
    document.getElementById("destination-input");

let timer = null;

destinationInput.addEventListener("input", function () {

    clearTimeout(timer);

    timer = setTimeout(async () => {

        if (this.value.length < 2) {

            hideSuggestions();

            return;
        }

        const result =
            await getLocation(this.value);

        renderSuggestions(result.Items);

    },300);

});

function renderSuggestions(items){

    const box =
        document.getElementById("destination-suggestions");

    box.innerHTML="";

    items.forEach(item=>{

        const div=document.createElement("div");

        div.className="destination-item";

        div.innerText=item.Display;

        div.onclick=()=>{

            destinationInput.value=item.Display;

            box.style.display="none";

            window.selectedLocation=item;

        };

        box.appendChild(div);

    });

    box.style.display="block";
}

function hideSuggestions(){

    document.getElementById(
        "destination-suggestions"
    ).style.display="none";
}