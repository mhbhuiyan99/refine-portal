function initializeCurrencyDropdown() {

    const saved =
        localStorage.getItem("currency") || "US";

    populateCurrencyDropdown(saved);

    const select =
        document.getElementById("currency-select");

    select.addEventListener("change", function () {

        localStorage.setItem("currency", this.value);

        // Refine page
        if (typeof renderTiles === "function" && window.allProperties) {
            renderTiles(window.allProperties, this.value);
        }

        // Category page
        if (typeof updateCategoryPrices === "function") {
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


