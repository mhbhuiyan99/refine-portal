const EXCHANGE_RATE = {
    BD: 123,
    US: 1,
    CA: 1.37,
};

function convertPrice(priceUSD, cuntryCode) {
    const rate = EXCHANGE_RATE[cuntryCode] || 1;
    return Math.round(priceUSD * rate);
}

function formatCurrency(price, countryCode) {
    if (!price) return "";

    const exchangeRate = {
        BD: 123,
        US: 1,
        CA: 1.37,
        AU: 1.53,
        GB: 0.74,
        EU: 0.86,
    };

    const currencySymbol = {
        BD: "৳",
        US: "$",
        CA: "C$",
        AU: "A$",
        GB: "£",
        EU: "€",
    };

    const rate = exchangeRate[countryCode] ?? 1;
    const symbol = currencySymbol[countryCode] ?? "$";

    return `${symbol}${Math.round(price * rate).toLocaleString()}`;
}