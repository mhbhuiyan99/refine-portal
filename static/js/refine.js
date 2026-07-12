const params = new URLSearchParams(window.location.search);

const search = params.get("search") || "";
const order = params.get("order") || "1";

console.log("Search: ", search);
console.log("Order: ", order);

async function init() {
    try {
        const location = await getLocation(search)

        const category = location.GeoInfo.LocationSlug;
        const locations = location.GeoInfo.CountryCode;

        const properties = await getProperties (
            category,
            locations,
            order
        );

        console.log(properties);
    } catch (error) {
        console.log(error);
    }

    
}

init()