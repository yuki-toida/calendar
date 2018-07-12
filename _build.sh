#!/bin/sh

if [ "$1" != "dev" ]; then
  echo "error. \$1 undefined env (dev)"
  exit 1
fi

# node process kill
killall node

cd assets
yarn deploy &
cd ../

ENV=$1
REGISTRY=asia.gcr.io/planet-pluto-$ENV
IMAGE=knowme-$ENV

# latest -> stable
latest=`gcloud container images list-tags $REGISTRY/$IMAGE --filter='tags:latest'`
if [ "$latest" != "" ]; then
  gcloud container images add-tag $REGISTRY/$IMAGE:latest $REGISTRY/$IMAGE:stable
fi

# build and push latest
gcloud container builds submit --config=_cloudbuild-${ENV}.yaml .

# delete untag
digest=`gcloud container images list-tags $REGISTRY/$IMAGE --filter='-tags:*' --format='get(digest)'`
if [ "$digest" != "" ]; then
  echo digest: $digest
  gcloud container images delete --quiet $REGISTRY/$IMAGE@$digest
fi
