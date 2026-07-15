{{template "layouts/header.tpl" .}}

<div class="container">

    <section id="category-hero">

        <h1>
            {{.Category.GeoInfo.Name}} Vacation Rentals
        </h1>

        <p class="property-count">
            {{.Category.GeoInfo.PropertyCount}} vacation rentals
        </p>

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

    <hr>

        <section id="category-sections">

        {{range .Category.Result.Sections}}

        <section class="category-section">

            <div class="section-header">

                <h2>{{.Title}}</h2>

                <p>{{.SubTitle}}</p>

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

<script>
    window.imageBaseURL = "{{.ImageBaseURL}}";
</script>

{{template "layouts/footer.tpl" .}}