const data = {
  temperature: ['1'],
  pressure: ['3', '6', '4'],
  precipitation: ['1', '8', '7'],
};

export default function recoApi(req, res) {
  setTimeout(() => {
    res.send(data[req.query.sku]);
  }, 1000);
}
