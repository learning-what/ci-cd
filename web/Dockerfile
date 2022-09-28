# nodejs images
FROM node:16

WORKDIR /src

COPY . .

RUN npm install && npm run build

# nginx images
FROM nginx:alpine

CMD ["rm", "-r", "./*"]

COPY --from=node:16 /src/dist/ /usr/share/nginx/html

COPY --from=node:16 /src/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]