FROM node:slim

WORKDIR /f
COPY . /f
RUN npm i --only=prod

CMD [ "node", "/f/index.js" ]
