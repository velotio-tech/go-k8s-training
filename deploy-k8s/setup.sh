#! /bin/bash

# setup environment variables
source config.sh

# check pre-requisites
if ! hash docker > /dev/null
then
    echo "docker could not be found."
    echo "exiting..."
    exit 1
fi

if ! hash kubectl > /dev/null
then
    echo "kubectl could not be found."
    echo "exiting..."
    exit 1
fi

if ! hash helm > /dev/null
then
    echo "helm could not be found"
    echo "exiting..."
    exit 1
fi

#build docker images for service and orders
docker build -t user:latest ../user
docker build -t order:latest ../order

kubectl create ns ecommerce
kubectl -n ecommerce create secret generic mongodb-creds --from-literal=DB_USERNAME=${DB_USERNAME} --from-literal=DB_PASSWORD="${DB_PASSWORD}" --from-literal=DB_URL=${DB_URL}

helm dependency build ./ecommerce
helm upgrade --create-namespace -n ecommerce --install --wait  ecommerce ./ecommerce
kubectl -n ecommerce delete pods --all