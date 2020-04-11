Project : 
=========
topdownloads

Problem Statement :
===================
Find the most popular and the 2nd most popular jar file (artifact) in a maven repository. The most popular artifact will be the one with highest number of downloads. In addition, to build and deploy this solution, create a robust CI/CD pipeline via Jenkins or other CI tool.

Algorithm :
===========
1. Every 5 secs, POST request is sent to Artifactory to get a list of all artifacts in a repository.
2. Next, a GET request is sent to get all file's stats matching the criteria of repo type (eg: jcentre-cache) & file type (eg: .jar) and top K downloads.
3. The GET response is compared with cached data to see if any new data is available.
4. If new data is available,a Priority-queue based frequency check algorithm is run to find top K downloads.


CI/CD Pipeline :
================
1. Go-based topdownloads binary
2. Travis CI : Used for Build, Test and Deploy and Continuous Integration
3. Github : Source code repository connected to Travis CI
3. Docker Hub : Docker image pushed to Hub
4. Heroku : Build is deployed on Heroku app for Continous Deployment

CI/CD Results :
===============
https://travis-ci.com/github/drishtu11/topdownloads/builds/159933745

Github Repo:
============
https://github.com/drishtu11/topdownloads

Binary Usage :
==============
1. Shell: ./topdownloads 104.154.94.138 jcenter-cache .jar 3
2. Docker : docker run topdownloads ./topdownloads 104.154.94.138 jcenter-cache .jar 3
3. Heroku: heroku run /app/bin/topdownloads 104.154.94.138 jcenter-cache .jar 3 -a topdownloads

Results / Output:
================
Before downloading a .jar file from Artifactory

----------------------------------------

Top Downloads

----------------------------------------
Artifact : struts2-core-2.3.14.jar
Downloads : 23

Artifact : ognl-3.0.6.jar
Downloads : 20

----------------------------------------

After downloading ognl-3.0.6.jar from Artifactory:

----------------------------------------

Top 2 Downloads

----------------------------------------
Artifact : struts2-core-2.3.14.jar
Downloads : 23

Artifact : ognl-3.0.6.jar
Downloads : 21

----------------------------------------
