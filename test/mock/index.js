const express = require('express');
const fs = require('fs');

const app = express();
const endpoints = [
  'create/vdisk/level/r5/disks/2.6,2.7,2.8/vd-1'
];

endpoints.forEach(endpoint => app.get(`/${endpoint}`, (_, res) => {
  let rawData;
  let path = endpoint.split('/');

  while (path.length) {
    try {
      console.log(path.join('_'));
      rawData = fs.readFileSync(`../data/${path.join('_')}.xml`);
      break;
    }
    catch (e) {
      path.splice(path.length - 1, 1);
    }
  }

  if (!rawData) {
    throw new Error('what?');
  }

  res.setHeader('content-type', 'application/xml');
  res.send(rawData);
}));

app.listen(8080);
