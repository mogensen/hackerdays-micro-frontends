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
        this.innerHTML = '<h2>TEST</h1>';
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