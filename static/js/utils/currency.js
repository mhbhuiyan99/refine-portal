const EXCHANGE_RATE = {
    BD: 123,
    US: 1,
    CA: 1.37,
    AU: 1.53,
    GB: 0.74,
    EU: 0.86,
};

const CURRENCY_SYMBOL = {
    BD: "৳",
    US: "$",
    CA: "C$",
    AU: "A$",
    GB: "£",
    EU: "€",
};

function getExchangeRate(countryCode) {
    return EXCHANGE_RATE[countryCode] ?? 1;
}

function getCurrencySymbol(countryCode) {
    return CURRENCY_SYMBOL[countryCode] ?? "$";
}

function convertPrice(priceUSD, countryCode) {
    return Math.round(priceUSD * getExchangeRate(countryCode));
}

function formatCurrency(price, countryCode) {
    if (!price) return "";

    const symbol = getCurrencySymbol(countryCode);
    const converted = convertPrice(price, countryCode);

    return `${symbol}${converted.toLocaleString()}`;
}