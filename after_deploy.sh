#!/bin/sh
set -e # Stop script from running if there are any errors

# // docker not found on heroku
#docker rm -f $(docker ps -a -q)
#docker rmi -f $(docker images -q)
#docker run topdownloads ./topdownloads 104.154.94.138 jcenter-cache .jar 2
heroku run timeout 180s /app/bin/topdownloads 104.154.94.138 jcenter-cache .jar 2 -a topdownloads
heroku run timeout 180s /app/bin/topdownloads 104.154.94.138 jcenter-cache .pom 10 -a topdownloads
