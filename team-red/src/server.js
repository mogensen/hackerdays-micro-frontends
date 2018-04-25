/* eslint-disable no-console */
import express from 'express';
import morgan from 'morgan';

const app = express();
app.use(morgan('dev'));
app.use('/red/images', express.static('./images'));
app.use('/red', express.static('./build'));


app.listen(3003);
console.log(`💚  team red running. fragments are available here:
>> http://127.0.0.1:3003/red-info`);
