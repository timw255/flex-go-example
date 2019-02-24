var path = require('path');
var express = require('express');
var app = express();
 
express.static.mime.types['wasm'] = 'application/wasm';

app.use('/wasm', express.static(path.join(__dirname + '/../')));
app.use(express.static(path.join(__dirname + '/')));

console.log('listening on :8080');
app.listen(8080);