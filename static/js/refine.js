const params = new URLSearchParams(window.location.search);

const search = params.get("search") || "";
const order = params.get("order") || "1";

console.log("Search: ", search);
console.log("Order: ", order);

async function init() {
  try {
    const location = await getLocation(search);
    console.log("Location:", location);

    const properties = await getProperties(
      location.GeoInfo.LocationSlug,
      location.GeoInfo.CountryCode,
      order,
    );
    console.log("Properties:", properties);

    const propertyIDs = properties.Result.ItemIDs;
    console.log("IDs:", propertyIDs);
    console.log("Count:", propertyIDs.length);

    const propertyDetails = await getPropertyDetails(propertyIDs);
    console.log("Details:", propertyDetails);

    renderTiles(propertyDetails);
  } catch (error) {
    console.log(error);
  }
}

init();
