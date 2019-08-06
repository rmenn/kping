### kping

Simple webapp to demonstrate kubernetes health check failures & rolling deployments

#### Start the image on the cluster

`kubectl run test image=rmenn/kping:0.0.1`

#### Expose the deployment as a service

on minikube

`kubectl expose deployment test --port=80 --type=LoadBalancer`

`minikube service test`

on docker-for-mac

`kubectl expose deployment test --port=80 --type=NodePort`

#### Update the image version to 0.0.2

Running this in a split terminal will show you the real time updates

`while true; do curl -m 0.5 -w "\t %{http_code} \n" <your_endpoint>/ping; done`

Edit the manifest and update the image in the container spec

`kubectl edit deployment/test`

You should see a few timeouts occuring


#### Exercise - Add HealthChecks to the application. Endpoint is `/healthz`
health checks should remove the container from the service when it returns a non 2XX-3XX responce

to test this

`curl -X POST <endpoint>/healthz`

this will set the healthz endpoint to return an internal server error
the pod will be restarted

#### Run curl and update the image 0.0.3

