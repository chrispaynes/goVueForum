FROM node:8-alpine as build-stage
WORKDIR /app
COPY ./ui/package.json .
RUN npm config set depth 0 \
	&& npm cache clean --force \
	&& npm i
COPY ./ui .
RUN npm run build

FROM nginx:1.13.3-alpine
COPY nginx/default.conf /etc/nginx/conf.d/
RUN rm -rf /var/www/html
COPY --from=build-stage /app/dist /var/www/html
CMD ["nginx", "-g", "daemon off;"]