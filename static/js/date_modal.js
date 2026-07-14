let datePicker = null;

function openDateModal() {

    let modal = document.getElementById("date-modal");

    if (!modal) {

        document.body.insertAdjacentHTML(
            "beforeend",
            getDateModalHTML()
        );

        modal = document.getElementById("date-modal");

        document
            .getElementById("date-close")
            .onclick = closeDateModal;

        document
            .getElementById("date-cancel")
            .onclick = closeDateModal;

        document
            .getElementById("date-continue")
            .onclick = applyDates;

        modal.onclick = function (e) {

            if (e.target === modal) {
                closeDateModal();
            }

        };

        datePicker = flatpickr("#date-range", {

            inline: true,
            mode: "range",
            dateFormat: "Y-m-d",

            showMonths: 2,
            monthSelectorType: "static",

            prevArrow: "‹",
            nextArrow: "›"
        });

    }

    modal.style.display = "flex";
}

function closeDateModal() {

    document
        .getElementById("date-modal")
        .style.display = "none";

}

function applyDates() {
    const dates = datePicker.selectedDates;

    if (dates.length === 2) {
        const start = datePicker.formatDate(dates[0], "M d");
        const end = datePicker.formatDate(dates[1], "M d");
        const label = `${start} - ${end}`;

        const dateFilterBtn = document.getElementById("date-filter-btn");
        if (dateFilterBtn) dateFilterBtn.textContent = label;

        const modalDateBtn = document.getElementById("modal-date-btn");
        if (modalDateBtn) modalDateBtn.textContent = `📅 ${label}`;
    }

    closeDateModal();
}

function getDateModalHTML() {

    return `
<div
    id="date-modal"
    class="date-modal">

    <div class="date-dialog">

        <div class="date-header">

            <h2>
                When do you want to travel?
            </h2>

            <button
                id="date-close">
                ✕
            </button>

        </div>

        <div class="date-body">

            <input
                id="date-range">

        </div>

        <div class="date-footer">

            <button
                id="date-cancel">

                Cancel

            </button>

            <button
                id="date-continue">

                Continue

            </button>

        </div>

    </div>

</div>
`;
}