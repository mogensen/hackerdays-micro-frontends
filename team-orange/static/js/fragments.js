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
        var dataline = "";
        var ylabel = "";

        if (sku == "temperature") {
            dataline = 'Temperature.Value'
            ylabel =  'Temp (C)';
        } else if (sku == "precipitation") {
            dataline = 'Precipitation.Value'
            ylabel = 'Regn i mm'
        } else if (sku == "pressure") {
            dataline = 'Pressure.Value'
            ylabel = 'Mm kviksølv'
        }

        var xhr = new XMLHttpRequest();
        xhr.open('GET', "/orange/api/weather/");
        xhr.onload = function () {
            if (xhr.status === 200) {
                var chart = c3.generate({
                    bindto: "#chart",
                    data: {
                        xFormat: '%Y-%m-%dT%H:%M:%SZ',
                        json: JSON.parse(xhr.responseText),
                        keys: {
                            x: 'Timestamp',
                            value: [dataline],
                        },
                        
                    },
                    axis: {
                        x: {
                            type: 'timeseries',
                            tick: { format: '%Y-%m-%d %H:%M' }
                        },
                        y: {
                            show: true,
                            label: {
                                text: ylabel,
                                position: 'outer-middle'
                            }
                        }
                    },
                    tooltip: { show: true }
                });
            }
            else {
                alert('Request failed.  Returned status of ' + xhr.status);
            }
        };
        xhr.send();
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