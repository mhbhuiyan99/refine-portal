function openFilterModal(active = null) {
    const existing = document.getElementById("filter-modal");
    if (existing) existing.remove(); // always rebuild with latest data

    document.body.insertAdjacentHTML("beforeend", getFilterModalHTML());
    const modal = document.getElementById("filter-modal");

    document.getElementById("filter-close").addEventListener("click", closeFilterModal);
    modal.addEventListener("click", (e) => {
        if (e.target === modal) closeFilterModal();
    });
    bindModalEvents();

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

    const minSlider = document.getElementById("min-price");
    const maxSlider = document.getElementById("max-price");
    const minInput = document.getElementById("min-price-value");
    const maxInput = document.getElementById("max-price-value");
    const fill = document.getElementById("price-slider-fill");

    // Dynamic scale limits based on the actual range configuration
    const range = window.priceRange || { min: 0, max: 50000 };
    const MIN_PRICE = Number(range.min);
    const MAX_PRICE = Number(range.max);
    const MIN_GAP = Math.round((MAX_PRICE - MIN_PRICE) * 0.01) || 100; // 1% of the total range as minimum gap

    minSlider.min = MIN_PRICE;
    minSlider.max = MAX_PRICE;
    maxSlider.min = MIN_PRICE;
    maxSlider.max = MAX_PRICE;

    minSlider.value = Number(minInput.value);
    maxSlider.value = Number(maxInput.value);

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

    function updatePriceUI() {
        const min = Number(minSlider.value);
        const max = Number(maxSlider.value);

        // Fix the calculation to prevent division by zero if MIN and MAX are equal
        const totalRange = MAX_PRICE - MIN_PRICE || 1;
        const minPercent = ((min - MIN_PRICE) / totalRange) * 100;
        const maxPercent = ((max - MIN_PRICE) / totalRange) * 100;

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

    document.querySelector('.price-slider').addEventListener('click', (e) => {
        if (e.target.tagName === 'INPUT') return; // dragging a thumb already handles itself

        const rect = e.currentTarget.getBoundingClientRect();
        const percent = (e.clientX - rect.left) / rect.width;
        const value = Math.round(MIN_PRICE + percent * (MAX_PRICE - MIN_PRICE));

        const minVal = Number(minSlider.value);
        const maxVal = Number(maxSlider.value);

        if (Math.abs(value - minVal) <= Math.abs(value - maxVal)) {
            minSlider.value = Math.min(value, maxVal - MIN_GAP);
        } else {
            maxSlider.value = Math.max(value, minVal + MIN_GAP);
        }
        updatePriceUI();
    });

    updatePriceUI(); // draw initial fill on modal open
}

function getFilterModalHTML() {
    const range = window.priceRange || { min: 0, max: 50000 };
    const countryCode = window.currencyCode || "US";
    const symbol = getCurrencySymbol(countryCode);

    return `
<div id="filter-modal" class="filter-modal">
<div class="filter-dialog">
<div class="filter-header">
<h2>Filters</h2>
<button id="filter-close">✕</button>
</div>
<div class="filter-body">

<div class="filter-top">
<label class="icon-label">
    <span class="filter-icon paw-icon">🐾</span> Pet-friendly only
    <input type="checkbox"><span class="checkmark"></span>
</label>
<label class="icon-label">
    <span class="filter-icon leaf-icon">🍃</span> Eco-friendly only
    <input type="checkbox"><span class="checkmark"></span>
</label>
</div>

<div class="filter-row-2col">
<div class="filter-section" id="date-section">
<h3>Select a date</h3>
<button id="modal-date-btn">
    <span>Select Date</span>
    <span class="date-icon">📅</span>
</button>
</div>

<div class="filter-section" id="guest-section">
<h3>Guests</h3>
<div class="guest-box">
<button id="guest-minus">−</button>
<span id="guest-count">0</span>
<button id="guest-plus">+</button>
</div>
</div>
</div>

<div id="price-section" class="filter-section">
<h3>Price range</h3>
<div class="price-slider">
    <div class="price-slider-track"></div>
    <div class="price-slider-fill" id="price-slider-fill"></div>
        <input id="min-price" type="range" min="${range.min}" max="${range.max}" value="${range.min}">
        <input id="max-price" type="range" min="${range.min}" max="${range.max}" value="${range.max}">
    </div>
    <div class="price-inputs">
        <span class="price-label side-label">Min price</span>
        <div class="price-input-group">
        <span class="price-currency">${countryCode} ${symbol}</span>
        <input id="min-price-value" type="number" value="${range.min}">
    </div>
    <span>—</span>
    <div class="price-input-group">
    <span class="price-currency">${countryCode} ${symbol}</span>
    <input id="max-price-value" type="number" value="${range.max}">
    </div>
    <span class="price-label side-label">Max price</span>
</div>
</div>

<div class="filter-section">
<h3>Amenities</h3>
<div class="amenities-grid">
<label>Air Conditioner<input type="checkbox"><span class="checkmark"></span></label>
<label>Balcony/terrace<input type="checkbox"><span class="checkmark"></span></label>
<label>Bedding/linens<input type="checkbox"><span class="checkmark"></span></label>
<label>Breakfast<input type="checkbox"><span class="checkmark"></span></label>
<label>Child Friendly<input type="checkbox"><span class="checkmark"></span></label>
<label>Hot Tub<input type="checkbox"><span class="checkmark"></span></label>
<label>Internet/Wifi<input type="checkbox"><span class="checkmark"></span></label>
<label>Kitchen<input type="checkbox"><span class="checkmark"></span></label>
<label>Laundry<input type="checkbox"><span class="checkmark"></span></label>
</div>
</div>

</div>
<div class="filter-footer">
<button id="filter-clear">Clear</button>
<button id="filter-search">Search</button>
</div>
</div>
</div>
`;
}