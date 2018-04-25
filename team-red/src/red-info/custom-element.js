
import renderHtml from './render';

class RedInfo extends HTMLElement {

  static get observedAttributes() {
    return ['sku'];
  }

  sendUpdateEvent(sku) {
    var event = new CustomEvent('skuUpdated', { sku: sku });
    this.dispatchEvent(event);
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
    this.innerHTML = renderHtml(sku)
  }

  disconnectedCallback() {
    const sku = this.getAttribute('sku');
    this.log('disconnected', sku);
  }

  log(...args) {
    console.log('🖼️ orange-recos', ...args);
  }
}

export default RedInfo;
