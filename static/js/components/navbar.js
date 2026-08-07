function initializeCurrencyDropdown() {

    const saved =
        localStorage.getItem("currency") || "US";

    populateCurrencyDropdown(saved);

    const select =
        document.getElementById("currency-select");

    select.addEventListener("change", function () {

        window.currencyCode = this.value;
        localStorage.setItem("currency", this.value);

        // Refine page
        if (typeof applyFilters === "function" && window.allProperties) {
            // recompute range for new currency
            window.priceRange =
                computePriceRange(
                    window.allProperties,
                    window.currencyCode
                );

            // reset price filter
            window.filterState.minPrice = window.priceRange.min;
            window.filterState.maxPrice = window.priceRange.max;

            applyFilters();
        }

        // Category page
        else if (typeof updateCategoryPrices === "function") {
            updateCategoryPrices();
        }

    });

}


function populateCurrencyDropdown(selected = "US") {

    const select =
        document.getElementById("currency-select");

    if (!select) {
        return;
    }

    select.innerHTML = "";

    const addedCodes = new Set();

    Object.entries(CURRENCIES).forEach(([key, currency]) => {

        // Skip duplicate currencies (EUR, GBP, etc.)
        if (addedCodes.has(currency.Code)) {
            return;
        }
        addedCodes.add(currency.Code);

        const option =
            document.createElement("option");

        option.value = key;

        option.textContent =
            `${currency.Code} (${currency.Symbol})`;

        option.selected =
            key === selected;

        select.appendChild(option);

    });

}


