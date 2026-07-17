document.addEventListener("DOMContentLoaded", function () {

    const dateInput = document.getElementById("category-date");
    const guestField = document.getElementById("guest-field");
    const destinationInput = document.getElementById("destination-input");
    const browseButton = document.getElementById("browse-rentals-btn");

    // Not category page
    if (!destinationInput) {
        return;
    }

    if (dateInput) {
        dateInput.addEventListener("click", function () {
            openDateModal("category", this);
        });
    }

    if (guestField) {
        guestField.addEventListener("click", function (e) {
            e.stopPropagation();
            toggleGuestPopup();
        });
    }

    let timer = null;

    destinationInput.addEventListener("input", function () {

        clearTimeout(timer);

        timer = setTimeout(async () => {

            if (this.value.length < 2) {
                hideSuggestions();
                return;
            }

            const result = await getLocation(this.value);

            renderSuggestions(result.Items);

        }, 300);

    });

    browseButton.addEventListener("click", searchCategory);

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

async function searchCategory() {

    const keyword = document
        .getElementById("destination-input")
        .value
        .trim();

    if (!keyword) {
        alert("Please enter a destination.");
        return;
    }

    const result = await getLocation(keyword);

    if (!result.Success) {
        alert("Location not found.");
        return;
    }

    window.location.href =
        "/all/" + result.GeoInfo.LocationSlug;
}