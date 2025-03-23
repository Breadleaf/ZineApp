#
# stage 1: build app
#

# init Node image
FROM node:23.10.0-alpine3.21 AS builder
WORKDIR /app

# copy files and install dependencies
# create a cache layer for less frequently changing code
COPY package.json ./
COPY package-lock.json* ./
RUN npm install
# create a cache layer for more frequently changing code
COPY . .

# build app for production
RUN npm run build

#
# stage 2: serve the built app
#

# init Nginx image
FROM nginx:1.27.4-alpine3.21

# copy the builder output to Nginx's public folder
COPY --from=builder /app/dist /usr/share/nginx/html
COPY certs /etc/nginx/certs
COPY nginx.conf /etc/nginx/conf.d/default.conf

# expose the port the Nginx server listens on
EXPOSE 80

# start Nginx when container launches
CMD ["nginx", "-g", "daemon off;"]