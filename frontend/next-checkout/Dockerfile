# Dockerfile-nextjs-checkout
FROM node:16.15-alpine

CMD ["/bin/sh"]

# set working directory
WORKDIR /usr/src/app

RUN apk update

COPY package.json ./

RUN npm install