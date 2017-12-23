FROM node:slim

WORKDIR /fx
COPY . /fx
RUN npm i --only=prod

CMD [ "node", "/fx/index.js" ]
