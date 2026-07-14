function renderFilters() {

    const container = document.getElementById("filters");

    container.innerHTML = `
        <button class="filter-btn" id="date-filter-btn">
            Dates
        </button>

        <button class="filter-btn" id="price-filter-btn">
            Price
        </button>

        <button class="filter-btn" id="guest-filter-btn">
            Guests
        </button>

        <button class="filter-btn" id="more-filter-btn">
            More
        </button>
    `;

    document
        .getElementById("date-filter-btn")
        .addEventListener("click", () => {
            openDateModal();
        });

    document
        .getElementById("price-filter-btn")
        .addEventListener("click", () => {
            openFilterModal("price");
        });

    document
        .getElementById("guest-filter-btn")
        .addEventListener("click", () => {
            openFilterModal("guest");
        });

    document
        .getElementById("more-filter-btn")
        .addEventListener("click", () => {
            openFilterModal();
        });
}

function formatButtonDate(date) {

    if (!date) {
        return "";
    }

    return date.toLocaleDateString("en-US", {
        month: "short",
        day: "numeric"
    });

}

function updateFilterButtons() {

    const state = window.filterState;

    // Date
    const dateBtn = document.getElementById("date-filter-btn");

    if (state.startDate && state.endDate) {

        dateBtn.innerHTML = `
            ${formatButtonDate(state.startDate)} - ${formatButtonDate(state.endDate)}
            <span class="filter-clear" data-filter="date">×</span>
        `;

    } else {

        dateBtn.textContent = "Dates";

    }

    // Price
    const priceBtn = document.getElementById("price-filter-btn");

    if (
        state.minPrice !== window.priceRange.min ||
        state.maxPrice !== window.priceRange.max
    ) {

        priceBtn.innerHTML = `
            BD ৳${state.minPrice} - BD ৳${state.maxPrice}
            <span class="filter-clear" data-filter="price">×</span>
        `;

    } else {

        priceBtn.textContent = "Price";

    }

    // Guests
    const guestBtn = document.getElementById("guest-filter-btn");

    if (state.guests > 0) {

        guestBtn.innerHTML = `
            ${state.guests} Guests
            <span class="filter-clear" data-filter="guest">×</span>
        `;

    } else {

        guestBtn.textContent = "Guests";

    }

    // More
    const moreBtn = document.getElementById("more-filter-btn");

    let count = 0;

    if (state.petFriendly) count++;
    if (state.ecoFriendly) count++;

    if (count > 0) {

        moreBtn.innerHTML = `
            More
            <span class="filter-count">${count}</span>
        `;

    } else {

        moreBtn.textContent = "More";

    }

    bindClearButtons();
}

function bindClearButtons() {

    document
        .querySelectorAll(".filter-clear")
        .forEach(button => {

            button.onclick = function (e) {

                e.stopPropagation();

                const filter =
                    this.dataset.filter;

                switch (filter) {

                    case "date":

                        window.filterState.startDate = null;
                        window.filterState.endDate = null;

                        if (datePicker) {
                            datePicker.clear();
                        }

                        break;

                    case "price":

                        window.filterState.minPrice =
                            window.priceRange.min;

                        window.filterState.maxPrice =
                            window.priceRange.max;

                        break;

                    case "guest":

                        window.filterState.guests = 0;

                        break;

                }

                applyFilters();

                updateFilterButtons();

            };

        });

}