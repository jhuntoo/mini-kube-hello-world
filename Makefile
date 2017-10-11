start:
	minikube start

build-binary:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bin/main .

build: build-binary
	eval $$(minikube docker-env)
	docker build --rm=true -t mini-kube-hello-world .

up: start build
	kubectl apply -f k8s.yaml
	while true; do curl $$(minikube service hello-world --url); echo '\n'; sleep 1; done