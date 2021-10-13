#!/bin/sh -x

kubectl create -f ordersapp-deploy.yaml

kubectl create -f orders-svc.yaml

kubectl create -f usersapp-deploy.yaml

kubectl create -f users-svc.yaml

kubectl get pods,deployments,svc -l app=userswebapp

kubectl get pods,deployments,svc -l app=orderswebapp