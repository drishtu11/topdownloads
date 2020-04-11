#!/bin/sh
set -e # Stop script from running if there are any errors

# // docker not found on heroku
#docker rm -f $(docker ps -a -q)
#docker rmi -f $(docker images -q)
#docker run topdownloads ./topdownloads 104.154.94.138 jcenter-cache .jar 2

# // Heroku command
heroku ps:restart --app topdownloads && heroku run timeout 60s /app/bin/topdownloads 104.154.94.138 jcenter-cache .jar 2 -a topdownloads
#heroku ps:restart --app topdownloads && heroku run timeout 60s /app/bin/topdownloads 104.154.94.138 jcenter-cache .pom 10 -a topdownloads
