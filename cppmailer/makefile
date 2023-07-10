BINARY=mailer
CC=g++
CFLAGS=-Wall -lcurl
SOURCE=main.cpp
IMAGE=localhost:5001/mailer

.PHONY: all run clean build_image push_image deploy god

all: ${BINARY}X

${BINARY}X: ${SOURCE}
	${CC} ${CFLAGS} -o $@ $<

run: ${BINARY}X
	./${BINARY}X

clean:
	rm -f ./${BINARY}X
	kubectl delete -f mailer.yaml
	kubectl delete -f mailhog.yaml

build_image:
	docker build -t ${IMAGE} .

push_image: build_image
	docker push ${IMAGE}

deploy: push_image
	kubectl apply -f mailer.yaml
	kubectl apply -f mailhog.yaml

god: all deploy

#prep_registry: docker run -d -p 5000:5000 --restart=always --name registry registry:2