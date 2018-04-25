
import render from './render';

class RedInfo extends HTMLElement {

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
    this.innerHTML = render()
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
