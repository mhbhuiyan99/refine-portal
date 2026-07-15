function renderPropertyCard(item, countryCode) {

    const p = item.Property;

    const image =
        p.FeatureImage
            ? `/static/images/placeholder.jpg`
            : "/static/images/placeholder.jpg";

    const price =
        formatCurrency(
            convertPrice(
                p.Price,
                window.locationData.GeoInfo.CountryCode
            ),
            window.locationData.GeoInfo.CountryCode
        );

    const guests =
        p.Counts.Occupancy > 0
            ? `<span>${p.Counts.Occupancy} Guests</span>`
            : "";

    return `
        <div class="property-card">

            <div class="image-wrapper">

                <img
                    class="property-image"
                    src="${image}"
                    alt="${p.PropertyName}">

            </div>

            <div class="property-body">

                <div class="property-type">
                    ${p.PropertyType}
                </div>

                <h3 class="property-title">
                    ${p.PropertyName}
                </h3>

                <div class="property-rating">

                    ⭐ ${p.ReviewScore}

                </div>

                <div class="property-attribute">

                    ${p.PropertyAttribute}

                </div>

                <div class="property-meta">

                    <span>${p.Counts.Bedroom} Bedrooms</span>

                    •

                    <span>${p.Counts.Bathroom} Bathrooms</span>

                    ${guests ? "•" : ""}

                    ${guests}

                </div>

                <div class="property-footer">

                    <div class="price">

                         ${formatCurrency(p.Price, countryCode)}

                        <span>/night</span>

                    </div>

                    <button class="deal-btn">

                        View Deal

                    </button>

                </div>

            </div>

        </div>
    `;
}