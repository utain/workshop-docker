FROM node:10.11-alpine
ENV CI=true
WORKDIR /app
ADD ./yarn.lock /app/yarn.lock
ADD ./package.json /app/package.json
ENV API_URL=http://localhost:8080/api
RUN yarn
ADD . /app
RUN yarn test
RUN yarn build

FROM nginx:1.14.2-alpine
COPY --from=0 /app/build /usr/share/nginx/html
WORKDIR /usr/share/nginx/html
ADD ./nginx.conf /etc/nginx/
CMD ["nginx", "-g", "daemon off;"]