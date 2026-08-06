
function populateCurrencyDropdown(selected = "US") {

    const select =
        document.getElementById("currency-select");

    if (!select) {
        return;
    }

    select.innerHTML = "";

    Object.entries(CURRENCIES).forEach(([key, currency]) => {

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