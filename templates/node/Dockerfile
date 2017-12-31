FROM node

WORKDIR /app
COPY . /app
RUN npm i --only=prod

CMD [ "npm", "start" ]
