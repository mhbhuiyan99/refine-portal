<div class="property-card">

    <div class="image-wrapper">

        <img
            class="property-image"
            src="/static/images/placeholder.jpg"
            alt="{{.Property.PropertyName}}">

    </div>

    <div class="property-body">

        <div class="property-type">

            {{.Property.PropertyType}}

        </div>

        <h3 class="property-title">

            {{.Property.PropertyName}}

        </h3>

        <div class="property-rating">

            ⭐ {{.Property.ReviewScore}}

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

            <a
                class="deal-btn"
                href="{{.Partner.URL}}"
                target="_blank">

                View Deal

            </a>

        </div>

    </div>

</div>