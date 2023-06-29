#!/bin/sh

echo "pull project yunyuansheng"
git clone git@github.com:shootingclub/yunyuansheng.git
cd module3
echo "build images"
docker build -t httpserver:last .
echo "image list"
docker images
echo "run container"
docker run -p 8999:8999 --name httpserver -d httpserver:last
