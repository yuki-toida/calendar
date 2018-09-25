#!/bin/sh

# node process kill
killall node

cd assets
rm -rf dist
yarn deploy
cd ../

ENV=$1
REGISTRY=asia.gcr.io/planet-pluto-dev
IMAGE=knowme

# build and push latest
gcloud builds submit --config=_cloudbuild.yaml .

# delete untag
digest=`gcloud container images list-tags $REGISTRY/$IMAGE --filter='-tags:*' --format='get(digest)'`
if [ "$digest" != "" ]; then
  echo digest: $digest
  gcloud container images delete --quiet $REGISTRY/$IMAGE@$digest
fi
