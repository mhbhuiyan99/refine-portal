function renderBreadcrumb(location){

    const html = location.GeoInfo.Breadcrumbs
        .map(b => `<span>${b.Name}</span>`)
        .join(" &gt; ");

    document.getElementById("breadcrumb").innerHTML = html;
}