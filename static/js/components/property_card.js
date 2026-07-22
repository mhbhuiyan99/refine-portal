function renderPropertyCard(item, countryCode) {
    const p = item.Property;

    const image = p.FeatureImage || "/static/images/placeholder.jpg";

    // console.log("Image_src: ", image);
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

    const amenities =
        (p.TopListedAmenities || [])
            .slice(0, 2)
            .map(item => `
                <span class="amenity-tag">
                    <img
                        src="${getAmenityIcon(item.Name)}"
                        class="amenity-icon"
                        alt="${item.Name}"
                    >
                    ${item.Name}
                </span>
            `)
            .join("");

    const partnerLogo = getPartnerLogo(item.Feed);
    const partnerUrl = item.Partner?.URL || "#";
    //console.log("Partner URL:", partnerUrl);

    const state = item.GeoInfo?.State || "";
    const city = item.GeoInfo?.City || "";
    const distance = item.GeoInfo?.DistanceFromCenter || "";

    return `
        <div
            class="property-card"
            data-property-type="${p.PropertyType}">

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

                <div class="property-location">

                    <img
                        src="/static/images/amenities/map_pin.svg"
                        class="location-icon"
                        alt="Location">

                    ${state}

                    ${state && city ? " • " : ""}

                    ${city}

                    ${
                        distance
                            ? `<span>${distance} to center</span>`
                            : ""
                    }

                </div>

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

                <div class="property-amenities">

                    ${amenities}

                </div>

                <div class="property-footer">

                    <div class="price">

                        ${formatCurrency(p.Price, countryCode)}

                        <span>/night</span>

                    </div>

                    <div class="deal-section">

                       <a
                            class="partner-logo"
                            href="${partnerUrl}"
                            target="_blank">

                            ${
                                partnerLogo
                                    ? `<img src="${partnerLogo}" alt="Partner">`
                                    : ""
                            }

                        </a>

                        <a
                            class="deal-btn"
                            href="${partnerUrl}"
                            target="_blank">

                            View Deal

                        </a>

                    </div>

                </div>

            </div>

        </div>
    `;
}