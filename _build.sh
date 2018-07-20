#!/bin/sh

if [ "$1" != "dev" ]; then
  echo "error. \$1 undefined env (dev)"
  exit 1
fi

# node process kill
killall node

cd assets
rm -rf dist
yarn deploy
cd ../

ENV=$1
REGISTRY=asia.gcr.io/planet-pluto-$ENV
IMAGE=knowme-$ENV

# build and push latest
gcloud container builds submit --config=_cloudbuild-${ENV}.yaml .

# delete untag
digest=`gcloud container images list-tags $REGISTRY/$IMAGE --filter='-tags:*' --format='get(digest)'`
if [ "$digest" != "" ]; then
  echo digest: $digest
  gcloud container images delete --quiet $REGISTRY/$IMAGE@$digest
fi
