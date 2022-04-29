# Dockerfile-react-ambassador
FROM node:16.15-alpine

CMD ["/bin/sh"]

# set working directory
WORKDIR /usr/src/app

RUN apk update && \
    npm install -g create-react-app
    
COPY package.json ./

RUN npm install