# Dockerfile-nextjs-checkout
FROM node:16.15-alpine

CMD ["/bin/sh"]

# set working directory
WORKDIR /usr/src/app

# install app dependencies
COPY package.json ./
COPY package-lock.json ./

RUN npm install -g npm@8.8.0
RUN npm install --silent \
    npm install react-scripts@5.0.1 -g --silent
