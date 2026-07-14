const params = new URLSearchParams(window.location.search);

const search = window.refineConfig.search;
const order = window.refineConfig.order;

console.log("Search: ", search);
console.log("Order: ", order);

async function init() {
  renderFilters();
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

    function computePriceRange(propertyDetails, countryCode) {
      const prices = propertyDetails.Items
          .map(item => item.Property.Price)
          .filter(p => typeof p === "number" && p > 0)
          .map(p => convertPrice(p, countryCode));

      if (prices.length === 0) {
          return { min: 0, max: 50000 };
      }

      return {
          min: Math.floor(Math.min(...prices)),
          max: Math.ceil(Math.max(...prices)),
      };
  }

    const countryCode = location.GeoInfo.CountryCode;

    window.priceRange = computePriceRange(propertyDetails, countryCode);
    window.currencyCode = countryCode;
    window.allProperties = propertyDetails.Items;

    window.filterState = {
      startDate: null,
      endDate: null,

      guests: 0,

      minPrice: window.priceRange.min,
      maxPrice: window.priceRange.max,

      amenities: [],

      petFriendly: false,
      ecoFriendly: false
  };

    renderTiles(window.allProperties, countryCode);
    updateFilterButtons();
  } catch (error) {
    console.log(error);
  }
}

init();
