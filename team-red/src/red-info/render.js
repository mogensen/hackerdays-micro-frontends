const product = {
  name: 'Skanderborg, Danmark',
  variants: [
    {
      sku: 'temperature',
      color: 'red',
      name: 'Temperatur',
      image: '/red/images/thermometer.png',
      thumb: '/red/images/thermometer.png',
      price: '66,00 €',
    },
    {
      sku: 'pressure',
      color: 'green',
      name: 'Lufttruk',
      image: '/red/images/pressure.png',
      thumb: '/red/images/pressure.png',
      price: '54,00 €',
    },
    {
      sku: 'precipitation',
      color: 'blue',
      name: 'Forventet regn',
      image: '/red/images/precipitation.png',
      thumb: '/red/images/precipitation.png',
      price: '58,00 €',
    },
  ],
};

function renderOption(variant, sku) {
  const active = sku === variant.sku ? 'active' : '';
  return `
  <a href="/${variant.sku}" class="${active}" type="button" data-sku="${variant.sku}" onclick="window.skuchanged(event, this)">
  <img src="${variant.thumb}" alt="${variant.name}" />
  </a>
  `;
}

export default function renderBuy(sku = 'temperature') {
  const variant = product.variants.find(v => sku === v.sku);
  if (!variant) { return '<pre>no product not found</pre>'; }
  return `
      <h1 id="store">YR Weather service</h1>
      
      <h2 id="name">${product.name} - ${variant.name}</h2>
      <div id="options">${product.variants.map(v => renderOption(v, sku)).join('')}</div>
      `;
}
