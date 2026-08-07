document.addEventListener("DOMContentLoaded", function () {

    initializeCurrencyDropdown();
    updateCategoryPrices();

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


    browseButton.addEventListener("click", searchCategory);
    buildPropertyTypeTabs();

    initializeCategorySliders();
});


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

    let url = "/refine?search=" + encodeURIComponent(result.GeoInfo.Display)

    if (
        window.filterState.startDate &&
        window.filterState.endDate
    ) {
        url += 
            "&dateStart=" +
            flatpickr.formatDate(
                window.filterState.startDate,
                "Y-m-d"
            );
        
        url += 
            "&dateEnd=" + 
            flatpickr.formatDate(
                window.filterState.endDate,
                "Y-m-d"
            );
    }

    const pax =
        window.filterState.guests > 0
            ? window.filterState.guests
            : 2;

    url += "&pax=" + pax;
    
    url += "&order=1";

    window.location.href = url;
}

function buildPropertyTypeTabs() {

    const cards =
        document.querySelectorAll(".property-card");

    const tabContainer =
        document.getElementById("property-type-tabs");

    if (!tabContainer || cards.length === 0) {
        return;
    }

    const types = [
        ...new Set(
            [...cards].map(card =>
                card.dataset.propertyType
            )
        )
    ];

    const allButton = document.createElement("button");
    allButton.className = "property-type-tab active";
    allButton.textContent = "All";

    allButton.onclick = () => {

        document
            .querySelectorAll(".category-section")
            .forEach(section => {

                section.style.display = "";

                section
                    .querySelectorAll(".property-card")
                    .forEach(card => {

                        card.style.display = "";

                    });

            });

        document
            .querySelectorAll(".property-type-tab")
            .forEach(btn =>
                btn.classList.remove("active"));

        allButton.classList.add("active");
    };

    tabContainer.appendChild(allButton);

    types.forEach(type => {

        const btn =
            document.createElement("button");

        btn.className =
            "property-type-tab";

        btn.textContent = type;

        btn.onclick = () => {

            document
                .querySelectorAll(".category-section")
                .forEach(section => {

                    let visible = 0;

                    section
                        .querySelectorAll(".property-card")
                        .forEach(card => {

                            if (card.dataset.propertyType === type) {
                                card.style.display = "";
                                visible++;
                            } else {
                                card.style.display = "none";
                            }

                        });

                    // hide empty section
                    section.style.display = visible ? "" : "none";

                });

            document
                .querySelectorAll(".property-type-tab")
                .forEach(button =>
                    button.classList.remove("active"));

            btn.classList.add("active");
        };

        tabContainer.appendChild(btn);

    });
}

function initializeCategorySliders() {
    const sliders = document.querySelectorAll(".property-slider");
    sliders.forEach(initSlider);
}
