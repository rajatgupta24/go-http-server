docker.build:
	docker build -t rajatguptag/go-server:1.2.4 .

docker.push: docker.build
	docker push rajatguptag/go-server:1.2.4

k.apply: docker.push
	kubectl apply -f k8s/server-deployment.yaml