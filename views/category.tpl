{{template "layouts/header.tpl" .}}

<div class="container">

    <section id="hero">

        <div class="hero-overlay">

            <h1>
                Owner Direct Vacation Rentals in
                {{.Category.GeoInfo.Name}}
            </h1>

            <p>
                The Best Vacation Rentals in {{.Category.GeoInfo.Name}} 
                - Plan Your Next Vacation to St. Augustine Today!
            </p>

            <div class="hero-search">

                <div class="search-field" id="destination-field">

                    <img src="/static/images/location.png" alt="">

                    <input
                        id="destination-input"
                        type="text"
                        value=" {{.Category.GeoInfo.Name}}"
                        autocomplete="off"
                        placeholder="Destination">

                </div>

                <div class="search-field" id="date-field">

                    <img src="/static/images/calendar.png" alt="">
                    <input
                        id="category-date"
                        readonly
                        placeholder="Dates"
                    >

                </div>

                <div class="search-field" id="guest-field">

                    <img src="/static/images/user.png" alt="">

                    <span id="guest-text">
                        Guests
                    </span>

                </div>

                <button id="browse-rentals-btn">
                    Browse Rentals
                </button>
            </div>

        </div>

    </section>

    <section class="breadcrumb">

        <a href="/">Home</a>

        {{range .Category.GeoInfo.Breadcrumbs}}

            <span class="separator">›</span>

            <a href="/all/{{.Slug}}">
                {{.Name}}
            </a>

        {{end}}

    </section>

    <section id="category-intro">

        <h2>
            Explore the Best
            {{.Category.GeoInfo.Name}}
            Vacation Homes
        </h2>

        <p>
            Looking for the perfect place to stay in
            {{.Category.GeoInfo.ShortName}}?
            Discover a curated collection of vacation rentals offering comfort,
            convenience, and authentic local experiences.
            Whether you want a beachfront condo, a cozy family home, or a pet-friendly stay close to downtown, 
            Owner Direct makes it easy to compare verified options for your next trip.
        </p>

    </section>

    <hr>
    <div id="property-type-tabs" class="property-type-tabs"></div>
    <section id="category-sections">

    {{range .Category.Result.Sections}}

    <section class="category-section">

        <div class="section-header">

            <h2>
                {{.Title}}
            </h2>

            <p>
                {{.SubTitle}}
            </p>

        </div>

        <div class="property-grid">

            {{range .Items}}

                {{template "components/property_card.tpl" .}}

            {{end}}

        </div>

    </section>

    {{end}}

    </section>
</div>



{{template "layouts/footer.tpl" .}}