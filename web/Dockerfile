FROM node:20-alpine AS build
# Astra bindings (acc5cb9): the platform injects ASTRA_PATH_API as a build arg;
# forwarded to the Vite build-time env so the bundle bakes the backend path.
ARG ASTRA_PATH_API=""
ENV VITE_API_BASE=$ASTRA_PATH_API
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
EXPOSE 80
