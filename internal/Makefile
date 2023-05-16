image:
	cd .. && docker build --platform linux/amd64 -t $(REPO)/go-workshop:latest -f internal/Dockerfile .

push:
	docker push $(REPO)/go-workshop:latest
