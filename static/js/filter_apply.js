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

    const min =
        Number(document.getElementById("min-price-value").value);

    const max =
        Number(document.getElementById("max-price-value").value);

    return properties.filter(item => {

        const price =
            convertPrice(
                item.Property.Price,
                window.currencyCode
            );

        return price >= min && price <= max;
    });
}

function filterByGuests(properties) {

    const guests =
        Number(
            document.getElementById("guest-count").textContent
        );

    if (guests === 0) {
        return properties;
    }

    return properties.filter(item => {

        return (
            item.Property.Counts &&
            item.Property.Counts.Occupancy >= guests
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