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