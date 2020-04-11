#!/bin/sh
set -e # Stop script from running if there are any errors

docker rm -f $(docker ps -a -q)
docker rmi -f $(docker images -q)
docker run drishtu11/topdownloads ./topdownloads 104.154.94.138 jcenter-cache .jar 2