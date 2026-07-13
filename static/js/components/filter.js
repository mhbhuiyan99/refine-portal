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

        <input
            type="date"
            id="date-picker"
            style="display:none">
    `;

    const picker = new Litepicker({
        element: document.getElementById("date-picker"),

        singleMode: false,

        numberOfMonths: 2,

        numberOfColumns: 2,

        autoApply: false,

        format: "MMM DD, YYYY",

        setup: (picker) => {

            document
                .getElementById("date-filter-btn")
                .addEventListener("click", () => picker.show());

        }
    });

    picker.on("selected", (start, end) => {

        document.getElementById("date-filter-btn").textContent =
            `${start.format("MMM DD")} - ${end.format("MMM DD")}`;

    });
}