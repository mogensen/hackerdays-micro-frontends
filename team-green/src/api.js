const data = {
  temperature: ['1'],
  pressure: ['2'],
  precipitation: ['7'],
};

export default function recoApi(req, res) {
  setTimeout(() => {
    res.send(data[req.query.sku]);
  }, 1000);
}
