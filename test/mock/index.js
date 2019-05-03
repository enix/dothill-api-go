const express = require('express');
const fs = require('fs');
const md5 = require('md5');

const app = express();
const userpass = 'manage_!manage';
const endpoints = [
  'create/vdisk/level/r5/disks/2.6,2.7,2.8/vd-1',
  'invalid/xml',
  'status/code/1',
];

app.get('/api/login/:token', (req, res) => {
  if (md5(userpass) === req.params.token) {
    res.send(fs.readFileSync(`data/login.xml`));
    return;
  }

  res.status(401).send('nope');
});

endpoints.forEach(endpoint => app.get(`/api/${endpoint}`, (req, res) => {
  let rawData;
  let path = endpoint.split('/');

  if (req.headers['sessionkey'] !== 'HereIsTheToken') {
    res.status(401).send('nope');
    return;
  }

  while (path.length) {
    try {
      rawData = fs.readFileSync(`data/${path.join('_')}.xml`);
      break;
    }
    catch (e) {
      path.splice(path.length - 1, 1);
    }
  }

  if (!rawData) {
    throw new Error(`XML input file for ${endpoint} missing in data folder`);
  }

  res.setHeader('content-type', 'application/xml');
  res.send(rawData);
}));

app.listen(8080);
