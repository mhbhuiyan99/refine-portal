function applyFilters() {

    let properties = [...window.allProperties];

    properties = filterByPrice(properties);
    properties = filterByGuests(properties);
    properties = filterByPet(properties);
    properties = filterByEco(properties);
    properties = filterByAmenities(properties);

    renderTiles(
        properties,
        window.currencyCode
    );
}

function filterByPrice(properties) {

    return properties.filter(item => {

        const price = convertPrice(
            item.Property.Price,
            window.currencyCode
        );

        return (
            price >= window.filterState.minPrice &&
            price <= window.filterState.maxPrice
        );

    });

}

function filterByGuests(properties) {

    if (window.filterState.guests === 0) {
        return properties;
    }

    return properties.filter(item => {

        return (
            item.Property.Counts &&
            item.Property.Counts.Occupancy >=
            window.filterState.guests
        );

    });

}

function filterByPet(properties) {

    return properties;
}

function filterByEco(properties) {

    return properties;
}

function filterByAmenities(properties) {

    return properties;
}

function updateFilterButtons() {

    if (window.filterState.guests > 0) {

        document.getElementById("guest-filter-btn")
            .textContent =
            `Guests (${window.filterState.guests})`;

    }

    if (
        window.filterState.minPrice > 0 ||
        window.filterState.maxPrice <
        window.priceRange.max
    ) {

        document.getElementById("price-filter-btn")
            .textContent =
            `${formatCurrency(window.filterState.minPrice, window.currencyCode)}
-
${formatCurrency(window.filterState.maxPrice, window.currencyCode)}`;

    }

}

function clearFilters() {

    window.filterState = {
        checkIn: "",
        checkOut: "",
        guests: 0,
        minPrice: window.priceRange.min,
        maxPrice: window.priceRange.max,
        amenities: [],
        petFriendly: false,
        ecoFriendly: false
    };

    renderTiles(
        window.allProperties,
        window.currencyCode
    );

    closeFilterModal();

}