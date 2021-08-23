PROJECT=demo
DOCKER_USER=watermelon
build:
	go build -o ${PROJECT}

# image:
# 	docker build --build-arg TOKEN=${TOKEN} --build-arg VERSION=${GITTAG} --build-arg COMMIT=${COMMIT} --build-arg BUILDTIME=${BUILD_TIME} -t ${DOCKER_USER}/${PROJECT}:latest .