# FROM node:8-alpine as build-stage
# WORKDIR /app
# COPY ./ui/package.json .
# RUN npm config set depth 0 \
# 	&& npm cache clean --force \
# 	&& npm i
# COPY ./ui .
# RUN npm run build

# FROM nginx:1.13.3-alpine
# COPY nginx/default.conf /etc/nginx/conf.d/
# RUN rm -rf /var/www/html
# COPY --from=build-stage /app/dist /var/www/html
# CMD ["nginx", "-g", "daemon off;"]

FROM nginx:1.14
RUN apt-get update \
    && apt-get install --no-install-recommends -y \
    build-essential \
    ca-certificates \
    curl \
    && curl -fsSLO --compressed "https://nodejs.org/dist/v10.9.0/node-v10.9.0-linux-x64.tar.xz" \
    && tar -xJf "node-v10.9.0-linux-x64.tar.xz" -C /usr/local --strip-components=1 --no-same-owner \
    && ln -s /usr/local/bin/node /usr/local/bin/nodejs \
    && apt-get purge -y curl build-essential ca-certificates \
    && rm -rf /var/lib/apt/lists/* \
    && rm node-v10.9.0-linux-x64.tar.xz
WORKDIR /app
COPY nginx/default.conf /etc/nginx/conf.d/
COPY ./ui/package.json .
RUN npm install \
    && npm cache clean --force
COPY ./ui .
RUN which npm
ENTRYPOINT npm run serve
