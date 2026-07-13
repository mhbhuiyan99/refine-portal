{{template "layout/header.tpl" .}}
<link rel="stylesheet" href="/static/css/refine.css">

    
<div class="refine-page">

    <main class="container">

        <div id="breadcrumb"></div>

        <div class="page-header">
            <h1 id="page-title">Loading...</h1>
        </div>

        <div class="toolbar">

            <div id="filters"></div>

            <div id="sort-container"></div>

        </div>

        <section
            id="property-container"
            class="property-grid">
        </section>

    </main>

</div>

<script>
    window.refineConfig = {
        search: "{{.Search}}",
        order: "{{.Order}}"
    };
</script>

{{template "layout/footer.tpl" .}}