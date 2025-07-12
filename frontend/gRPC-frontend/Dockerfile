# Stage 1: Build the Vite app
FROM node:20 AS build

WORKDIR /app
COPY . .

RUN npm install
RUN npm run build

# Stage 2: Serve using Nginx
FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
