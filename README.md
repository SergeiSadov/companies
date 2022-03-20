# Companies API app

This is an example of the REST API app using golang and postgres, deployable on kubernetes.

## Getting Started

### Running

To run an app using your local environment:

1. build the application using makefile commands (build/build_mac/build_windows)
2. inject env variables using example.env file as an example of required variables
3. use command ``./companies migrate up`` to run migrations

To run an app in kubernetes execute this command below:
```
kubectl apply -f k8s/postgres/postgres-configmap.yaml
kubectl apply -f k8s/postgres/postgres-storage.yaml  
kubectl apply -f k8s/postgres/postgres-deployment.yaml
kubectl apply -f k8s/postgres/postgres-service.yaml 

kubectl apply -f k8s/companies/companies-configmap.yaml
kubectl apply -f k8s/companies/companies-deployment.yaml
kubectl apply -f k8s/companies/companies-service.yaml 
```

App exposes 3000 port by default