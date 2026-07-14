function openFilterModal(active = null) {

    let modal = document.getElementById("filter-modal");

    if (!modal) {
        document.body.insertAdjacentHTML(
            "beforeend",
            getFilterModalHTML()
        );

        modal = document.getElementById("filter-modal");

        document
            .getElementById("filter-close")
            .addEventListener("click", closeFilterModal);

        modal.addEventListener("click", (e) => {
            if (e.target === modal) {
                closeFilterModal();
            }
        });

        bindModalEvents();
    }

    modal.style.display = "flex";

    highlightSection(active);
}

function closeFilterModal() {
    document.getElementById("filter-modal").style.display = "none";
}

function highlightSection(section) {

    document
        .querySelectorAll(".filter-section")
        .forEach(item => item.classList.remove("active"));

    if (!section) {
        return;
    }

    const target =
        document.getElementById(`${section}-section`);

    if (target) {
        target.classList.add("active");
        target.scrollIntoView({
            behavior: "smooth",
            block: "center"
        });
    }
}

function bindModalEvents() {

    const minus = document.getElementById("guest-minus");
    const plus = document.getElementById("guest-plus");
    const count = document.getElementById("guest-count");

    let guest = 0;

    minus.onclick = () => {
        if (guest > 0) {
            guest--;
            count.textContent = guest;
        }
    };

    plus.onclick = () => {
        guest++;
        count.textContent = guest;
    };

    document
        .getElementById("modal-date-btn")
        .onclick = () => {
            openDateModal();
        };

    document.getElementById("filter-clear").onclick = () => {
        guest = 0;
        count.textContent = "0";
        document.querySelectorAll("#filter-modal input[type=checkbox]").forEach(cb => cb.checked = false);

        minSlider.value = MIN_PRICE;
        maxSlider.value = MAX_PRICE;
        updatePriceUI();
    };

    document
        .getElementById("filter-search")
        .onclick = closeFilterModal;

    const MIN_PRICE = 0;
    const MAX_PRICE = 50000;
    const MIN_GAP = 500; // smallest allowed distance between the two handles

    const minSlider = document.getElementById("min-price");
    const maxSlider = document.getElementById("max-price");
    const minInput = document.getElementById("min-price-value");
    const maxInput = document.getElementById("max-price-value");
    const fill = document.getElementById("price-slider-fill");

    function updatePriceUI() {
        const min = Number(minSlider.value);
        const max = Number(maxSlider.value);

        const minPercent = ((min - MIN_PRICE) / (MAX_PRICE - MIN_PRICE)) * 100;
        const maxPercent = ((max - MIN_PRICE) / (MAX_PRICE - MIN_PRICE)) * 100;

        fill.style.left = `${minPercent}%`;
        fill.style.right = `${100 - maxPercent}%`;

        minInput.value = min;
        maxInput.value = max;
    }

        minSlider.oninput = () => {
            if (Number(minSlider.value) > Number(maxSlider.value) - MIN_GAP) {
                minSlider.value = Number(maxSlider.value) - MIN_GAP;
            }
            updatePriceUI();
        };

        maxSlider.oninput = () => {
            if (Number(maxSlider.value) < Number(minSlider.value) + MIN_GAP) {
                maxSlider.value = Number(minSlider.value) + MIN_GAP;
            }
            updatePriceUI();
        };

        minInput.oninput = () => {
            let value = Math.min(Number(minInput.value), Number(maxSlider.value) - MIN_GAP);
            value = Math.max(value, MIN_PRICE);
            minSlider.value = value;
            updatePriceUI();
        };

        maxInput.oninput = () => {
            let value = Math.max(Number(maxInput.value), Number(minSlider.value) + MIN_GAP);
            value = Math.min(value, MAX_PRICE);
            maxSlider.value = value;
            updatePriceUI();
        };

        updatePriceUI(); // draw initial fill on modal open
}

function getFilterModalHTML() {

    return `
    <div id="filter-modal" class="filter-modal">

        <div class="filter-dialog">

            <div class="filter-header">
            <h2>Filters</h2>
            <button id="filter-close">✕</button>
            </div>

            <div class="filter-body">

            <div class="filter-top">

            <label>
            <input type="checkbox">
            Pet-friendly only
            </label>

            <label>
            <input type="checkbox">
            Eco-friendly only
            </label>

        </div>

        <div
            class="filter-section"
            id="date-section">

            <h3>Select a date</h3>

            <button id="modal-date-btn">
            📅 Select Date
            </button>

        </div>

        <div
            class="filter-section"
            id="guest-section">

            <h3>Guests</h3>

            <div class="guest-box">

            <button id="guest-minus">−</button>

            <span id="guest-count">0</span>

            <button id="guest-plus">+</button>

        </div>

        </div>

        <div
            id="price-section"
            class="filter-section">

            <h3>Price range</h3>

            <div class="price-slider">
                <div class="price-slider-track"></div>
                <div class="price-slider-fill" id="price-slider-fill"></div>
                <input id="min-price" type="range" min="0" max="50000" value="0">
                <input id="max-price" type="range" min="0" max="50000" value="50000">
            </div>

            <div class="price-inputs">

                <input
                    id="min-price-value"
                    type="number"
                    value="0">

                <span>—</span>

                <input
                    id="max-price-value"
                    type="number"
                    value="50000">

            </div>

        </div>

        <div class="filter-section">
        <h3>Amenities</h3>
            <div class="amenities-grid">
                <label><input type="checkbox"> Air Conditioner</label>
                <label><input type="checkbox"> Balcony/terrace</label>
                <label><input type="checkbox"> Bedding/linens</label>
                <label><input type="checkbox"> Breakfast</label>
                <label><input type="checkbox"> Child Friendly</label>
                <label><input type="checkbox"> Hot Tub</label>
                <label><input type="checkbox"> Internet/Wifi</label>
                <label><input type="checkbox"> Kitchen</label>
                <label><input type="checkbox"> Laundry</label>
            </div>
        </div>

        </div>

        <div class="filter-footer">

        <button id="filter-clear">
        Clear
        </button>

        <button id="filter-search">
        Search
        </button>

        </div>

        </div>

    </div>
    `;
}