/* eslint-disable no-console */
import express from 'express';
import morgan from 'morgan';
import CryptoJS from 'crypto-js'
// var CryptoJS = require("crypto-js");

const app = express();

function generateAuthToken() {
  var UTCstring = new Date().toUTCString();
	var mydate = UTCstring.toLowerCase();
	var mytext = "post\ndocs\ndbs/RainbowDB/colls/rainbow\n" + mydate.toLowerCase() + "\n\n";
	var mykey = CryptoJS.enc.Base64.parse('uvDkJzLCQl8qaZDCWIN2EmzbCWSCkeqdHT2qgjrlQS117CmT2XyHoZXyA1di4zJjzm96alqdwinux9asyab4og==');
	var mysignature = CryptoJS.HmacSHA256(mytext, mykey);
	var mybase64Bits = CryptoJS.enc.Base64.stringify(mysignature);
  var myauth = encodeURIComponent("type=master&ver=1.0&sig=" + mybase64Bits);
  return {"token": myauth};
}

app.use(morgan('dev'));

app.use('/auth', (req, res) => {
	res.set({ 'content-type': 'application/json; charset=utf-8' })
  res.send(JSON.stringify(generateAuthToken()));
});

console.log("Token: " + generateAuthToken().token);
app.listen(3005);
console.log(`💚  auth running. fragments are available here:
>> http://127.0.0.1:3005/auth`);
