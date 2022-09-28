# nodejs images
FROM node:16 as builder

WORKDIR /src

COPY . .

RUN npm install && npm run build

# nginx images
FROM nginx:alpine

COPY --from=builder /src/dist/ /usr/share/nginx/html

COPY --from=builder /src/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

ENTRYPOINT ["nginx", "-g", "daemon off;"]