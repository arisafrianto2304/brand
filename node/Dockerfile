FROM node:lts-alpine3.20
WORKDIR /app
RUN npm install -g npm@latest
COPY package*.json ./
RUN npm install
COPY . .
# EXPOSE 3000
CMD [ "npm", "run", "dev" ]