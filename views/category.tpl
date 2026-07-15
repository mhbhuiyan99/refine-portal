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

    <p>Total Sections:
        {{len .Category.Result.Sections}}
    </p>

</div>

{{template "layouts/footer.tpl" .}}