all: kill clean build run 
		
build: build_base build_jupyter

build_base:
	docker build --network=host \
	--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
	--build-arg USER=$(shell id -u --name) --build-arg GROUP=$(shell id -g --name) \
	-f Dockerfile -t jcabrero/he_mia_base ..

build_jupyter:
	docker build --network=host \
	--build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) \
	--build-arg USER=$(shell id -u --name) --build-arg GROUP=$(shell id -g --name) \
	-f Dockerfile_Jupyter -t jcabrero/he_mia_jupyterlab ..

he_mia:
	docker run  \
		--net host \
		--hostname=$@ --name=$@ \
		-it --rm -v $(shell pwd)/../:/home/jocabrer/mnt \
		jcabrero/he_mia_base

run: he_mia 

run_jupyterlab:
	cd .. && jupyter lab --port=10093 --no-browser --ip=0.0.0.0 --allow-root

jupyterlab:
	docker run  \
		--net host \
		--hostname=$@ --name=$@ \
		-it --rm -v $(shell pwd)/../:/home/jocabrer/mnt \
		jcabrero/he_mia_jupyterlab 

clean:
	docker rmi $(shell docker images --filter dangling=true -q --no-trunc) || true

kill:
	docker kill he_mia jupyterlab || true

