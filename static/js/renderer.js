
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

    const sliders =
        container.querySelectorAll(
            ".property-slider"
        );

    sliders.forEach(slider => {

        if (slider.dataset.loaded === "true") {
            return;
        }

        slider.dataset.loaded = "true";

        initializePropertySlider(slider);

    });
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

const trigger =
    document.getElementById("load-more-trigger");

if (trigger) {
    observer.observe(trigger);
}

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

async function fetchPropertyImages(propertyId) {
    try {
        const response = await fetch(
            `/api/property/images/v1?propertyId=${propertyId}`
        );

        if (!response.ok) {
            throw new Error(`HTTP ${response.status}`);
        }

        return await response.json();

    } catch (error) {

        console.warn(
            `[PropertySlider] Failed to load images | propertyId=${propertyId}`,
            error
        );

        return null;
    }
}

function showSlide(
    slider,
    index
) {

    const images =
        slider.querySelectorAll(".slider-image");

    const dots =
        slider.querySelectorAll(".slider-dot");

    if (images.length === 0) {
        return;
    }

    if (index < 0) {
        index = images.length - 1;
    }

    if (index >= images.length) {
        index = 0;
    }

    images.forEach(image =>
        image.classList.remove("active")
    );

    dots.forEach(dot =>
        dot.classList.remove("active")
    );

    images[index].classList.add("active");

    if (dots[index]) {
        dots[index].classList.add("active");
    }

    slider.dataset.currentIndex = index;
}

async function initializePropertySlider(slider) {

    const propertyId =
        slider.dataset.propertyId;

    const result =
        await fetchPropertyImages(propertyId);

    if (
        !result ||
        !result.Success ||
        !Array.isArray(result.Images) ||
        result.Images.length === 0
    ) {
        return;
    }

    const firstImage =
        slider.querySelector(".property-image");

    if (!firstImage) {
        return;
    }

    // Extract image base url from the existing feature image
    const imageBaseURL =
        firstImage.src.substring(
            0,
            firstImage.src.lastIndexOf("/") + 1
        );

    const previousButton =
        slider.querySelector(".prev");

    const nextButton =
        slider.querySelector(".next");

    // Hide navigation if only one image exists.
    if (result.Images.length <= 1) {

        previousButton.style.display = "none";
        nextButton.style.display = "none";

        return;
    }

    const dotsContainer =
        slider.querySelector(".slider-dots");

    if (
        !previousButton ||
        !nextButton ||
        !dotsContainer
    ) {
        return;
    }

    // Remove placeholder image
    firstImage.remove();
    
    dotsContainer.innerHTML = "";

    // Create slider images
    result.Images.forEach((imageName, index) => {

        const image =
            document.createElement("img");

        image.src = imageBaseURL + imageName;
        image.className = "property-image slider-image";

        if (index === 0) {
            image.classList.add("active");
        }

        slider.insertBefore(image, previousButton);

        // Dot
        const dot = 
            document.createElement("span");
        
        dot.className = "slider-dot";

        if(index === 0) {
            dot.classList.add("active");
        }

        dot.dataset.index = index;

        dot.addEventListener("click", () => {
            showSlide(slider, index);
        });

        dotsContainer.appendChild(dot);
    });

    slider.dataset.currentIndex = 0;


    previousButton.addEventListener("click", () => {

        const current =
            Number(slider.dataset.currentIndex);

        showSlide(
            slider,
            current - 1
        );

    });

    nextButton.addEventListener("click", () => {

        const current =
            Number(slider.dataset.currentIndex);

        showSlide(
            slider,
            current + 1
        );

    });
}