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