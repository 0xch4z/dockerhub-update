#!/usr/bin/env bash

IMG=charliekenney23/dockerhub-update

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

if [ -n "$TRAVIS_TAG" ]; then
  docker tag $IMG $IMG:$TRAVIS_TAG
fi

docker push $IMG

docker run -v $(PWD)/README.md:/data/README.md:ro -t "$IMG" \
  -e DOCKERHUB_USERNAME="$DOCKER_USERNAME" \
  -e DOCKERHUB_PASSWORD="$DOCKER_PASSWORD" \
  -r /data/README.md $IMG
