{{template "layouts/header.tpl" .}}

<div class="container">

    <h1>{{.Category.GeoInfo.Name}}</h1>

    <hr>

    <h2>Breadcrumbs</h2>

    <ul>

        {{range .Category.GeoInfo.Breadcrumbs}}

            <li>{{.Name}}</li>

        {{end}}

    </ul>

    <hr>

    <h2>Sections</h2>

    <hr>

    {{range .Category.Result.Sections}}

    <div class="category-section">
        <h3>{{.Title}}</h3>
        <p> {{.SubTitle}} </p>
        <p>
            Properties:
            {{len .Items}}
        </p>

        <ul>
            {{range .Items}}

            <li>
                <strong>{{.Property.PropertyName}}</strong>
                <br>

                Price: 
                {{.Property.Price}}
                <br>

                Type:
                {{.Property.PropertyType}}
                <br>

                URL:
                <a href="{{.Partner.URL}}" target="_blank">
                    View Partner
                </a>
                <br><br>
            </li>

            {{end}}
        </ul>

        <hr>
    </div>

    {{end}}
</div>

{{template "layouts/footer.tpl" .}}