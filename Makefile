# PROJECT=demo
# DOCKER_USER=watermelon
# build:
# 	go build -o ${PROJECT}

# image:
# 	docker build --build-arg TOKEN=${TOKEN} --build-arg VERSION=${GITTAG} --build-arg COMMIT=${COMMIT} --build-arg BUILDTIME=${BUILD_TIME} -t ${DOCKER_USER}/${PROJECT}:latest .
.PHONY: all-in-one
all-in-one:
	docker build -f ./main . -t demo

# .PHONY: admin-api
# admin-api:
# 	docker build -f deployments/docker/admin-api/Dockerfile . -t admin-api
#
# .PHONY: agent-api
# agent-api:
# 	docker build -f deployments/docker/agent-api/Dockerfile . -t agent-api
#
# .PHONY: config-srv
# config-srv:
# 	docker build -f deployments/docker/config-srv/Dockerfile . -t config-srv
#
# .PHONY: micro
# micro:
# 	docker build -f deployments/docker/micro/Dockerfile . -t micro
#
# .PHONY: agent
# agent:
# 	docker build -f deployments/docker/agent/Dockerfile . -t agent
