#!/bin/bash

echo "building react app..."

# build react application

cd ./web

npm install && npm run build

echo "react build complete!" && ls -al

rm -rf $HOME/global/nginx/${DOMAIN}/ci-cd

cp dist/* $HOME/global/nginx/${DOMAIN}/ci-cd/

echo "nginx deployment done!" 

# build go application

echo "building go app"

cd ..

echo "remove origin container..."

sudo docker rm ci-cd > /dev/null 2>&1

echo "build nad run new container..."

sudo docker compose up -d

echo "go application deployment done!"
