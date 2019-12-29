VERSION = 1.0.0
DOCKER_ORG = hexagram30
REDIX_VERSION = 1.10
DOCKER_DIR = build/docker
REDIX_DIR = $(DOCKER_DIR)/redixdb

redix-image:
	docker build -t $(DOCKER_ORG)/redixdb $(REDIX_DIR)
	docker tag $(DOCKER_ORG)/redixdb $(DOCKER_ORG)/redixdb:latest
	docker tag $(DOCKER_ORG)/redixdb $(DOCKER_ORG)/redixdb:$(REDIX_VERSION)

redix-publish:
	docker push $(DOCKER_ORG)/redixdb:latest
	docker push $(DOCKER_ORG)/redixdb:$(REDIX_VERSION)

redix-run:
	docker run -it \
		-p 7090:7090 -p 6380:6380 \
		-v `pwd`/data/redixdb:/data \
		$(DOCKER_ORG)/redixdb:latest \
		-engine boltdb \
		-storage /data \
		-verbose
