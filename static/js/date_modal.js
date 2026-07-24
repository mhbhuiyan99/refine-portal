let datePicker = null;

function openDateModal(mode = "refine", input = null) {

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
            .onclick = function () {
                applyDates(mode, input);
            };

        modal.onclick = function (e) {

            if (e.target === modal) {
                closeDateModal();
            }

        };

        datePicker = flatpickr("#date-range", {
            mode: "range",
            inline: true,
            showMonths: 2,
            monthSelectorType: "static",
            dateFormat: "Y-m-d",
            static: true,
            minDate: "today"
        });
    }

    if (window.filterState.startDate && window.filterState.endDate) {

        datePicker.setDate([
            window.filterState.startDate,
            window.filterState.endDate
        ], false);

    }

    modal.style.display = "flex";

    datePicker.jumpToDate(
        window.filterState.startDate || new Date()
    );

}

function closeDateModal() {

    document
        .getElementById("date-modal")
        .style.display = "none";

}

function applyDates(mode = "refine", input = null) {

    const dates = datePicker.selectedDates;

    if (dates.length !== 2) {
        return;
    }

    if (mode === "category") {
        window.selectedStartDate = dates[0];
        window.selectedEndDate = dates[1];

        window.filterState.startDate = dates[0];
        window.filterState.endDate = dates[1];

        const nights =
            Math.ceil(
                (
                    dates[1] - dates[0]
                ) / (1000 * 60 * 60 * 24)
            );

        window.filterState.nights = nights;

        const start = flatpickr.formatDate(dates[0], "M j");
        const end = flatpickr.formatDate(dates[1], "M j");

        input.value = `${start} - ${end}`;
    } else {
        window.filterState.startDate = dates[0];
        window.filterState.endDate = dates[1];

        updateFilterButtons();
        applyFilters();
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
