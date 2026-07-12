function renderTiles(propertyDetails) {

    console.log(propertyDetails);

    if (!propertyDetails.Success) {
        console.error("Property Details API failed.");
        return;
    }

    if (!propertyDetails.Items || propertyDetails.Items.length === 0) {
        console.warn("No properties found.");
        return;
    }

    console.log(propertyDetails.Items[0]);

}