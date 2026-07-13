function renderTiles(response) {

    const container =
        document.getElementById("property-container");

    if (!response.Success || !response.Items) {
        container.innerHTML =
            "<p>No properties found.</p>";
        return;
    }

    const html = response.Items
        .map(renderTile)
        .join("");

    container.innerHTML = html;
}


function renderTile(item) {

    console.log(item);
    const p = item.Property;
    console.log(p);

    return `
        <div class="property-card">

            <img
                class="property-image"
                src="${p.FeatureImage}"
                alt="${p.PropertyName}">

            <div class="property-body">

                <div class="property-type">
                    ${p.PropertyType}
                </div>

                <h3>
                    ${p.PropertyName}
                </h3>

                <p>
                    ${p.City}
                </p>

                <div class="rating">
                    ⭐ ${p.ReviewScore}
                </div>

                <div class="price">
                    ৳${p.Price}/night
                </div>

            </div>

        </div>
    `;
}

