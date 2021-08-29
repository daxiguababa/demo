.PHONY: all-in-one
all-in-one:
	docker build -f ./Dockerfile . -t demo