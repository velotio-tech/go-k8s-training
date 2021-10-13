#!/bin/sh

echo "Checking users webapp deployment status..."
echo "------------------------------------------"
kubectl get pods,deployments,svc -l app=userswebapp
echo ""
echo "Checking orders webapp deployment status..."
echo "------------------------------------------"
kubectl get pods,deployments,svc -l app=orderswebapp