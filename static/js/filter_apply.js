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

    updateFilterButtons();
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

    if (!window.filterState.petFriendly) {
        return properties;
    }

    return properties.filter(item =>

        item.Property.PropertyAttribute &&
        item.Property.PropertyAttribute
            .toLowerCase()
            .includes("pet")

    );

}

function filterByEco(properties) {

    if (!window.filterState.ecoFriendly) {
        return properties;
    }

    return properties.filter(item =>

        item.Property.PropertyAttribute &&
        item.Property.PropertyAttribute
            .toLowerCase()
            .includes("eco")

    );

}

function filterByAmenities(properties) {

    if (window.filterState.amenities.length === 0) {
        return properties;
    }

    return properties.filter(item => {

        if (!item.Property.Amenities) {
            return false;
        }

        const amenities =
            Object.values(item.Property.Amenities);

        // returns only properties containing all selected(AND)
        // if want OR behavior need to replace every() with some()
        return window.filterState.amenities.every(selected =>
            amenities.includes(selected)
        );

    });

}


function clearFilters() {

    window.filterState = {

        startDate: null,
        endDate: null,

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

    updateFilterButtons();

    closeFilterModal();

}