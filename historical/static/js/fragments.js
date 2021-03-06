/* eslint-disable no-use-before-define, no-console, class-methods-use-this */
/* globals HTMLElement, window, CustomEvent 123*/

class HistoricalGraph extends HTMLElement {
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
        div.setAttribute("id", "historical-chart");

        this.innerHTML = div.outerHTML;
        var dataline = 'Value';
        var ylabel = "";
        var useSku = sku;
        if (sku == "temperature") {
            ylabel = 'Temp (C)';
        } else if (sku == "precipitation") {
            ylabel = 'Luftfugtughed i %'
            useSku = "humidity";
        } else if (sku == "pressure") {
            ylabel = 'Mm kviksølv'
        }

        var xhr = new XMLHttpRequest();
        xhr.open('GET', "/historical/iot/" + useSku);
        xhr.onload = function () {
            if (xhr.status === 200) {
                var chart = c3.generate({
                    bindto: "#historical-chart",
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
        console.log('🖼️ historical-graph', ...args);
    }
}

window.customElements.define('historical-graph', HistoricalGraph);