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

        const startDate = 
            window.pendingStartDate ||
            window.filterState.startDate;
        const endDate = 
            window.pendingEndDate ||
            window.filterState.endDate;

        if(startDate && endDate) {
            datePicker.setDate([
                startDate, 
                endDate
            ], false);
        }
    }

    modal.style.display = "flex";

    datePicker.jumpToDate(
        startDate || new Date()
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

    window.filterState.startDate = dates[0];
    window.filterState.endDate = dates[1];


    if (mode === "category") {
        window.selectedStartDate = dates[0];
        window.selectedEndDate = dates[1];

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
    }  else {
        
        // Refine page
        // Stage only. Do not apply yet.

        window.pendingStartDate = dates[0];
        window.pendingEndDate = dates[1];

        window.pendingNights = Math.ceil(
            (dates[1] - dates[0]) / (1000 * 60 * 60 * 24)
        );

        updateFilterDateButton();
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

function updateFilterDateButton() {

    const button = document.getElementById("modal-date-btn");

    if (!button) {
        return;
    }

    const startDate = 
        window.pendingStartDate ||
        window.filterState.startDate;
    const endDate = 
        window.pendingEndDate ||
        window.filterState.endDate;
        
    if (startDate && endDate) {
        button.querySelector("span:first-child").textContent =
            `${flatpickr.formatDate(startDate, "M j")} - ${flatpickr.formatDate(endDate, "M j")}`;
    } else {
        button.querySelector("span:first-child").textContent = 
            "Select Date";
    }
}