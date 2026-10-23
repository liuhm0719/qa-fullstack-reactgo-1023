FROM node:20-alpine AS build
# Astra bindings (acc5cb9): ASTRA_PATH_API addresses the backend component;
# NEXT_PUBLIC_BASE_PATH (fold) is this component's own prefix for assets.
ARG ASTRA_PATH_API=""
ARG NEXT_PUBLIC_BASE_PATH=""
ENV VITE_API_BASE=$ASTRA_PATH_API
ENV NEXT_PUBLIC_BASE_PATH=$NEXT_PUBLIC_BASE_PATH
WORKDIR /app
COPY package.json ./
RUN npm install
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
