// communicates with Beego APIs


async function getLocation(keyword) {
    const response = await fetch(
        `api/location?keyword=${encodeURIComponent(keyword)}`
    );

    if (!response.ok) {
        throw new Error("Failed to fetch location");
    }

    return await response.json();
}

const DEFAULT_PROPERTY_OPTIONS = {
    page: 1,
    limit: 192,
    items: 1,
    device: "desktop",
};

async function getProperties(category, locations, order) {

    const query = new URLSearchParams({
        category: category,
        location: locations,
        order,
        page: DEFAULT_PROPERTY_OPTIONS.page,
        limit: DEFAULT_PROPERTY_OPTIONS.limit,
        items: DEFAULT_PROPERTY_OPTIONS.items,
        device: DEFAULT_PROPERTY_OPTIONS.device,
    });

    const response = await fetch(
        `/api/properties?${query.toString()}`
    );

    if (!response.ok) {
        throw new Error("Failed to fetch properties");
    }

    return await response.json();
}

async function getPropertyDetails() {

}