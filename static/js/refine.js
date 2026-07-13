const params = new URLSearchParams(window.location.search);

const search = window.refineConfig.search;
const order = window.refineConfig.order;

console.log("Search: ", search);
console.log("Order: ", order);

async function init() {
  renderSort();
  
  try {
    const location = await getLocation(search);
    console.log("Location:", location);
    window.locationData = location;

    renderHeader(location);
    renderBreadcrumb(location);

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

    const countryCode = location.GeoInfo.CountryCode;
    renderTiles(propertyDetails, countryCode);
  } catch (error) {
    console.log(error);
  }
}

init();
