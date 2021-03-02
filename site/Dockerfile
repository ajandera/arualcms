FROM node:10

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

EXPOSE 9000

CMD ["npm", "run", "serve"]
