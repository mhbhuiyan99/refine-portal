function updateCategoryPrices() {

    const currency =
        localStorage.getItem("currency") ||
        window.locationData.GeoInfo.CountryCode;

    document.querySelectorAll(".price").forEach(priceEl => {

        const usdPrice =
            Number(priceEl.dataset.price);

        priceEl.innerHTML = `
            ${formatCurrency(usdPrice, currency)}
            <span>/night</span>
        `;

    });

}


