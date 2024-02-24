DOCKER_IMAGE_NAME = "go-playground"
.PHONY bash:
bash:
	docker build --tag $(DOCKER_IMAGE_NAME) -f Dockerfile .
	docker run -it --rm -v .:/home/app -w /home/app $(DOCKER_IMAGE_NAME) bash
