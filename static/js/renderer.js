
let allItems = [];
let renderedCount = 0;
let currentCountryCode = "";
let isLoading = false;
const PAGE_SIZE = 32;

function renderTiles(data, countryCode) {

    currentCountryCode = countryCode;
    
    const container =
        document.getElementById("property-container");

    let items = [];

    if (Array.isArray(data)) {

        items = data;

    } else if (data.Success && data.Items) {

        items = data.Items;

    } else {

        container.innerHTML = "<p>No properties found.</p>";
        return;
    }

    if (items.length === 0) {

        container.innerHTML = "<p>No properties found.</p>";
        return;
    }

    // Save all properties
    allItems = items;
    renderedCount = 0;
    container.innerHTML = "";
    renderNextBatch(countryCode);
}

function renderNextBatch(countryCode) {
    const container = 
        document.getElementById("property-container");

    const nextItems = 
        allItems.slice(
            renderedCount,
            renderedCount + PAGE_SIZE
        );

    nextItems.forEach(item => {
        container.insertAdjacentHTML(
            "beforeend",
            renderPropertyCard(item, countryCode)
            );
    });

    renderedCount += nextItems.length;
}

window.addEventListener("scroll", () => {

    if (isLoading) {
        return;
    }

    if (
        window.innerHeight + window.scrollY >=
        document.body.offsetHeight - 500 &&
        renderedCount < allItems.length
    ) {

        isLoading = true;

        requestAnimationFrame(() => {

            renderNextBatch(currentCountryCode);

            isLoading = false;

        });

    }

});

const observer = new IntersectionObserver(entries => {

    if (
        entries[0].isIntersecting &&
        renderedCount < allItems.length
    ) {

        renderNextBatch(currentCountryCode);

    }

}, {
    rootMargin: "300px"
});

observer.observe(
    document.getElementById("load-more-trigger")
);

function renderSkeletonCards(count = 32) {

    const container =
        document.getElementById("property-container");

    container.innerHTML = "";

    let html = "";

    for (let i = 0; i < count; i++) {

        html += `
            <div class="skeleton-card">

                <div class="skeleton-image"></div>

                <div class="skeleton-body">

                    <div class="skeleton-line skeleton-short"></div>

                    <div class="skeleton-line skeleton-title"></div>

                    <div class="skeleton-line skeleton-medium"></div>

                    <div class="skeleton-line"></div>

                </div>

            </div>
        `;
    }

    container.innerHTML = html;
}


