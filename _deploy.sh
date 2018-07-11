#!/bin/sh

if [ "$1" != "dev" ]; then
  echo "error. \$1 undefined env (dev)"
  exit 1
fi

ENV=$1
CONTAINER=knowme-$ENV
REGISTRY=asia.gcr.io/planet-pluto-$ENV

# イメージタグが指定されているか判定
if [ "$3" != "" ]; then
  IMAGE=$CONTAINER:$3
else
  # イメージに'latest'タグが設定されているか判定
  current=`kubectl get pods --selector=app=$CONTAINER -o jsonpath="{.items[*].spec.containers[0].image}"`
  if echo "$current" | grep ":latest"; then
    IMAGE=$CONTAINER
  else
    IMAGE=$CONTAINER:latest
  fi
fi

# deploy
kubectl set image deployment/knowme-deploy-$ENV $CONTAINER=$REGISTRY/$IMAGE
