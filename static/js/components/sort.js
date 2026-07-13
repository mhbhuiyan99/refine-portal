const params = new URLSearchParams(window.location.search);

params.set("order", selectedOrder);

window.location.search = params.toString();