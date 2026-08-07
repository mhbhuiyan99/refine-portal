const params = new URLSearchParams(window.location.search);

const search = window.refineConfig.search || "";
const order = window.refineConfig.order || 1;

//console.log("Search: ", search);
//console.log("Order: ", order);

async function init() {
  renderSkeletonCards(32);
  renderFilters();
  renderSort();

  try {
    const location = await getLocation(search);
    //console.log("Location:", location);
    window.locationData = location;

    renderHeader(location);
    renderBreadcrumb(location);

    const properties = await getProperties(
      location.GeoInfo.LocationSlug,
      location.GeoInfo.CountryCode,
      order,
    );
    //console.log("Properties:", properties);

    const propertyIDs = properties.Result.ItemIDs;
    //console.log("IDs:", propertyIDs);
    //console.log("Count:", propertyIDs.length);

    const propertyDetails = await getPropertyDetails(propertyIDs);
    //.log("Details:", propertyDetails);

    const countryCode = location.GeoInfo.CountryCode;

    const params = new URLSearchParams(window.location.search);

    const startDate = params.get("dateStart");
    const endDate   = params.get("dateEnd");
    const pax       = params.get("pax");

    window.currencyCode =
      localStorage.getItem("currency") ||
      countryCode;
    
    window.priceRange =
      computePriceRange(
          propertyDetails.Items,
          window.currencyCode
      );
    window.allProperties = propertyDetails.Items;

    window.filterState = {
        startDate: startDate ? new Date(startDate) : null,
        endDate: endDate ? new Date(endDate) : null,

        guests: pax ? Number(pax) : 0,

        minPrice: window.priceRange.min,
        maxPrice: window.priceRange.max,

        amenities: [],

        petFriendly: false,
        ecoFriendly: false
    };

    renderTiles(window.allProperties, window.currencyCode);
    //console.log(window.filterState);
    updateFilterButtons();
    initializeCurrencyDropdown();
  } catch (error) {
    console.log(error);
  }
}

init();
