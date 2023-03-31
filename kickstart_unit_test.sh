#!/bin/bash

# This bash script build your custom redis docker image locally.

USERNAME=bwalia
DOCKER_BIN=docker
IMAGE_NAME=go-restapi-unit-test
IMAGE_REPO=https://hub.docker.io

${DOCKER_BIN} build -f Dockerfile -t ${IMAGE_NAME} . --no-cache
${DOCKER_BIN} tag ${IMAGE_NAME} ${USERNAME}/${IMAGE_NAME}:${IMAGE_TAG}
# # ${DOCKER_BIN} push  ${IMAGE_REPO}/${IMAGE_NAME}:${IMAGE_TAG}

# # Push to docker public registry as well
# ${DOCKER_BIN} tag ${IMAGE_NAME} ${USERNAME}/${IMAGE_NAME}:latest && ${DOCKER_BIN} push ${USERNAME}/${IMAGE_NAME}:latest   
#--name ${IMAGE_NAME} -e name=Balinder
# ${DOCKER_BIN} run -d -p 6379:6379 ${USERNAME}/${IMAGE_NAME}:latest && ${DOCKER_BIN} push ${USERNAME}/${IMAGE_NAME}:latest   
${DOCKER_BIN} run ${IMAGE_NAME}