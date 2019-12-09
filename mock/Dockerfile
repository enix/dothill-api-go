FROM node:dubnium

WORKDIR /app

COPY package.json .

RUN yarn

COPY index.js .

COPY data data

CMD yarn start
