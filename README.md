## Load Balancing helloworld with Golang / Minikube

### Dependencies
- go
- minikube
- kubectl
- make

### Description of steps performed
1. Build a static binary of the hello world application
2. Build a docker container of the application within the minikube VM
3. Deploy a k8s service on minikube which will load balance across 3 replicas of the application
4. `curl` the service endpoint every 1 second to demonstrate load balanced requests


### Usage
```bash
make up
```

### Example output
```bash
Hello World from Pod hello-world-3200471853-wxpsm - 2017-10-11 12:31:24

Hello World from Pod hello-world-3200471853-t848r - 2017-10-11 12:31:25

Hello World from Pod hello-world-3200471853-t848r - 2017-10-11 12:31:27

Hello World from Pod hello-world-3200471853-1djfh - 2017-10-11 12:31:28

Hello World from Pod hello-world-3200471853-wxpsm - 2017-10-11 12:31:30

Hello World from Pod hello-world-3200471853-1djfh - 2017-10-11 12:31:31

Hello World from Pod hello-world-3200471853-1djfh - 2017-10-11 12:31:33
```