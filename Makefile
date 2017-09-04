.PHONY: list

list:
	@awk -F: '/^[A-z]/ {print $$1}' Makefile | sort

_ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd dockerfiles/$(image) && docker build -t floodio/$(image) .

push:
	docker push floodio/$(image)
