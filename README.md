# README:

This setup has been created to get an understanding of kubernetes.  I started
with the "Hello Minikube" tutorial at
https://kubernetes.io/docs/tutorials/hello-minikube/

In order to be able to set up a deployment file to describe the architecture, I
followed this tutorial:
https://medium.com/geekculture/deploying-a-multi-container-pod-to-a-kubernetes-cluster-3c86d1f04af1

## Target

Provide an API which allows to trigger a calculation. Internally, the
calculation is forwarded to worker nodes, which do the actual calculation. In
this basic example, they only doing a sleep for a defined number of milliseconds
and iterations, then return.  The worker nodes are executed asynchronous to save
time.  In a kubernetes environement, the worker containers should be able to get
duplicated in situations of high load. 

## Commands

### Building 

Starting the two container without kubernetes: 

`docker build -t worker:v1 . ` similar for the `initial` image. 

`docker run -d -p 9876:9876 --name initial-container initial`

`docker run -d -p 8765:8765 --name worker-container worker`

### Running

Curl command to the initial container to start the calculation: 

`curl localhost:9876 -d '{ "calls" : 3, "iterations" : 5, "wait_time" : 123 }'`

Of course, the two container can also be started with a 
`docker-compose up -d` command, using the `docker-compose.yml` file. 

### Hints

The worker ip address is hardcoded into the initial service. In the
docker-compose setup, this could be handled better.  

## Further plans

- setup a proper Kubernetes yaml to be able to deploy the stack with 
  load balancing 
- use a Postgres to store the jobs that need to be executed by the worker
  containers
- use Postgres events to trigger the worker container, probably by a dispatcher
  containre that picks up the Postgres events and starts the worker nodes
 
