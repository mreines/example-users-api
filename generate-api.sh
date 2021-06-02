#!/bin/bash

docker pull openapitools/openapi-generator-cli

docker run --rm -v "${PWD}:/local" openapitools/openapi-generator-cli generate \
    --global-property=modelDocs=false \
    --additional-properties=packageName=usersapi \
    --additional-properties=sourceFolder=api \
    -i /local/spec/users-api.yml \
    -g go-server \
    -o /local