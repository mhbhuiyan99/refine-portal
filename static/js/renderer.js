function renderTiles(data, countryCode) {

    const container =
        document.getElementById("property-container");

    let items = [];

    if (Array.isArray(data)) {

        items = data;

    } else if (data.Success && data.Items) {

        items = data.Items;

    } else {

        container.innerHTML = "<p>No properties found.</p>";
        return;
    }

    if (items.length === 0) {

        container.innerHTML = "<p>No properties found.</p>";
        return;
    }

    container.innerHTML =
        items
            .map(item => renderPropertyCard(item, countryCode))
            .join("");
}


