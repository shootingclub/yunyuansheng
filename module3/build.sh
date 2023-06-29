#!/bin/sh
set -e
username=mock_user
password=mock_password


echo "pull project yunyuansheng"
git clone git@github.com:shootingclub/yunyuansheng.git
cd module3
echo "build images"
docker build -t httpserver:last .
echo "image list"
docker images
echo "push images to docker repo"
docker login -u $username -p $password registry.hub.docker.com && docker push registry.hub.docker.com/httpserver:last
echo "run container"
docker run -p 8999:8999 --name httpserver -d httpserver:last
