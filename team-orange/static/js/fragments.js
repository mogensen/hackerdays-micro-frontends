/* eslint-disable no-use-before-define, no-console, class-methods-use-this */
/* globals HTMLElement, window, CustomEvent 123*/

class OrangeRecos extends HTMLElement {
    static get observedAttributes() {
        return ['sku'];
    }
    connectedCallback() {
        const sku = this.getAttribute('sku');
        this.log('connected', sku);
        this.render();
    }
    attributeChangedCallback(attr, oldValue, newValue) {
        this.log('attributeChanged', attr, newValue);
        this.render();
    }
    render() {
        const sku = this.getAttribute('sku');
        var output = []
        var div = document.createElement("div");
        div.setAttribute("id", "chart");

        this.innerHTML = div.outerHTML;
        
        var chart = c3.generate({
            bindto: "#chart",
            data: {
              columns: [
                ['data1', 30, 200, 100, 400, 150, 250],
                ['data2', 50, 20, 10, 40, 15, 25]
              ]
            },
            tooltip: {
                show: true
            }
        });
    }
    disconnectedCallback() {
        const sku = this.getAttribute('sku');
        this.log('disconnected', sku);
    }
    log(...args) {
        console.log('🖼️ orange-recos', ...args);
    }
}

window.customElements.define('orange-recos', OrangeRecos);