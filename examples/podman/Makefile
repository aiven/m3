.PHONY: images
images:
	podman build -t m3dbnode -f m3dbnode.Dockerfile .
	podman build -t m3coordinator -f m3coordinator.Dockerfile .
	podman build -t m3aggregator -f m3aggregator.Dockerfile .

.PHONY: cluster
cluster-up: images
	podman-compose -p m3db-cluster -f podman-compose-m3db-cluster.yml up -d

.PHONY: teardown
cluster-down:
	podman-compose -p m3db-cluster -f podman-compose-m3db-cluster.yml down

.PHONY: clean
clean: cluster-down
	podman volume prune -f
