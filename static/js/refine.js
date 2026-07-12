const params = new URLSearchParams(window.location.search);

const search = params.get("search") || "";
const order = params.get("order") || "1";

console.log("Search: ", search);
console.log("Order: ", order);

async function init() {
    try {
        const location = await getLocation(search)

        const properties = await getProperties (
            location.GeoInfo.LocationSlug,
            location.GeoInfo.CountryCode,
            order
        );

        const propertyIDs = properties.Result.ItemIDs;
        const propertyDetails = await getPropertyDetails(propertyIDs);

        console.log(propertyDetails);
    } catch (error) {
        console.log(error);
    }

    
}

init()