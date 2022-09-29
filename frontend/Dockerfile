FROM node:18-alpine3.15
WORKDIR /app
COPY app/package.json ./
RUN apk update && \
    apk upgrade && \
    npm install

CMD ["npm", "run", "dev"]