#!/usr/bin/env bash

IMG=charliekenney23/dockerhub-update

echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin

if [ -n "$TRAVIS_TAG" ]; then
  docker tag $IMG $IMG:$TRAVIS_TAG
fi

docker push $IMG

docker run -v $TRAVIS_BUILD_DIR/README.md:/data/README.md:ro -t \
  -e DOCKERHUB_USERNAME="$DOCKER_USERNAME" \
  -e DOCKERHUB_PASSWORD="$DOCKER_PASSWORD" \
  "$IMG" -r /data/README.md $IMG
