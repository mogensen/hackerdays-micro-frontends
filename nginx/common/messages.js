
function rerender(sku) {
    // removeListeners();
    document.getElementsByTagName("red-info")[0].setAttribute("sku", sku);
    document.getElementsByTagName("green-recos")[0].setAttribute("sku", sku);
    document.getElementsByTagName("orange-recos")[0].setAttribute("sku", sku);
    // addListeners();
}

window.skuchanged = function(event, element) {
    
    event.preventDefault();

    var sku = element.dataset.sku
    window.history.pushState(null, null, sku);
    rerender(element.dataset.sku)
}