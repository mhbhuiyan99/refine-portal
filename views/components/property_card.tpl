<div
    class="property-card"
    data-property-type="{{.Property.PropertyType}}">

    <div class="image-wrapper">

        <div
            class="property-slider"
            data-property-id="{{.ID}}">

            <img
                class="property-image"
                src="{{if .Property.FeatureImage}}{{.Property.FeatureImage}}{{else}}/static/images/placeholder.jpg{{end}}"
                alt="{{.Property.PropertyName}}">
            
            <div class="slider-loading hidden">
                <span></span>
                <span></span>
                <span></span>
            </div>

            <button
                class="slider-btn prev"
                aria-label="Previous Image">
                &#10094;
            </button>

            <button
                class="slider-btn next"
                aria-label="Next Image">
                &#10095;
            </button>

            <div class="slider-dots">
            </div>

        </div>

    </div>

    <div class="property-body">

        <div class="property-type">

            {{.Property.PropertyType}}

        </div>

        <h3 class="property-title">

            {{.Property.PropertyName}}

        </h3>

        <div class="property-location">

            <img
                src="/static/images/amenities/map_pin.svg"
                class="location-icon">

            {{.GeoInfo.State}}

            {{if and .GeoInfo.State .GeoInfo.City}}
                •
            {{end}}

            {{.GeoInfo.City}}

            {{if .GeoInfo.DistanceFromCenter}}
                <span>
                    {{.GeoInfo.DistanceFromCenter}} to center
                </span>
            {{end}}

        </div>

        <div class="property-rating">

            <div class="rating-badge">

                <img
                    src="/static/images/amenities/like-v2.svg"
                    class="rating-icon"
                    alt="Rating">

                <span>
                    {{.Property.ReviewScore}}
                </span>

            </div>

        </div>
        

        {{if .Property.PropertyAttribute}}

        <div class="property-attribute">

            {{.Property.PropertyAttribute}}

        </div>

        {{end}}

        <div class="property-meta">

            {{if gt .Property.Counts.Bedroom 0}}

                <span>{{.Property.Counts.Bedroom}} Bedrooms</span>

            {{end}}

            {{if gt .Property.Counts.Bathroom 0}}

                •
                <span>{{.Property.Counts.Bathroom}} Bathrooms</span>

            {{end}}

            {{if gt .Property.Counts.Occupancy 0}}

                •
                <span>{{.Property.Counts.Occupancy}} Guests</span>

            {{end}}

        </div>

        <div class="property-footer">

            <div class="price">

                ${{printf "%.0f" .Property.Price}}
                <span>/night</span>

            </div>

            <div class="deal-section">

                <div class="partner-logo">

                    {{if eq .Feed 11}}
                        <img src="/static/images/amenities/booking.svg" alt="Booking">

                    {{else if eq .Feed 12}}
                        <img src="/static/images/amenities/vrbo.svg" alt="Vrbo">

                    {{else if eq .Feed 24}}
                        <img src="/static/images/amenities/expedia.svg" alt="Expedia">

                    {{end}}

                </div>

                <a
                    class="deal-btn"
                    href="{{.Partner.URL}}"
                    target="_blank">

                    View Deal

                </a>

            </div>

        </div>

    </div>

</div>