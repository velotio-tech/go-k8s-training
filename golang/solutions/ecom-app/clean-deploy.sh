#!/bin/sh -x

kubectl delete -f ordersapp-deploy.yaml

kubectl delete -f orders-svc.yaml

kubectl delete -f usersapp-deploy.yaml

kubectl delete -f users-svc.yaml

kubectl get pods,deployments,svc -l app=userswebapp

kubectl get pods,deployments,svc -l app=orderswebapp