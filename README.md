# topdownloads

Algorithm :
===========
1. Every 5 secs, POST request is sent to Artifactory to get a list of all artifacts in a repository.
2. Next, a GET request is sent to get all file's stats matching the criteria of repo type (eg: jcentre-cache) & file type (eg: .jar) and top K downloads.
3. The GET response is compared with cached data to see if any new data is available.
4. If new data is available,a Priority-queue based frequency check algorithm is run to find top K downloads.

Notes :
=======

Project :
=========
The project consists of Go-based repository and has unit testing included.

CI/CD Pipeline :
================
1. Go-based topdownloads binary
2. Travis CI : Used for Build, Test and Deploy and Continuous Integration
3. Github : Source code repository connected to Travis CI
3. Docker Hub : Docker image pushed to Hub
4. Heroku : Build is deployed on Heroku app for Continous Deployment

Github Repo:
============
https://github.com/drishtu11/topdownloads

Binary Usage :
==============
1. Shell: ./topdownloads 104.154.94.138 jcenter-cache .jar 3
2. Docker : docker run topdownloads ./topdownloads 104.154.94.138 jcenter-cache .jar 3
3. Heroku: heroku run /app/bin/topdownloads 104.154.94.138 jcenter-cache .jar 3 -a topdownloads

