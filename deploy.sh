#!/bin/sh
set -e # Stop script from running if there are any errors

IMAGE="drishtu11/topdownloads"                             # Docker image
GIT_VERSION=$(git describe --always --abbrev --tags --long) # Git hash and tags

# Build and tag image
docker build . -f Dockerfile.multistage -t topdownloads
docker build . -f Dockerfile.multistage -t ${IMAGE}:${GIT_VERSION}
docker tag ${IMAGE}:${GIT_VERSION} ${IMAGE}:latest

# Log in to Docker Hub and push
# travis encrypt DOCKER_PASSWORD="<password>" --pro --add
# travis encrypt DOCKER_USERNAME="<username>" --pro --add
#travis env set DOCKER_USERNAME myusername
#travis env set DOCKER_PASSWORD secretsecret

echo "${DOCKER_PASSWORD}" | docker login -u "${DOCKER_USERNAME}" --password-stdin
docker push ${IMAGE}:${GIT_VERSION}
