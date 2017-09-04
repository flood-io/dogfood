.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

_ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd dockerfiles/$(image) && docker build -t floodio/$(image) .

push:
	docker push floodio/$(image)

install: ## Install target dependencies
	@echo "Installing target dependencies for $(image)"
	@cd dockerfiles/$(image) && docker build -t floodio/$(image) .
	@docker push floodio/$(image)
	@cd terraform/$(image)/elb && terraform apply -var-file=$(_ROOT_DIR)/terraform.tfvars
	@cd terraform/$(image)/asg && TF_VAR_dd_api_key=$(DD_API_KEY) terraform apply -var-file=$(_ROOT_DIR)/terraform.tfvars

health: ## Check target health
	@aws --profile=flooded --region us-west-2 elb describe-instance-health --load-balancer-name $(image)-elb | jq -r '.InstanceStates[] | .State' | sort | uniq -c | sort
	$(eval ELB_DNS_NAME := $(shell terraform output -state=terraform/$(image)/elb/terraform.tfstate dns_name))
	@echo "Checking ELB at $(ELB_DNS_NAME) for $(image)"
	@curl -I --silent --connect-timeout 3 http://$(ELB_DNS_NAME)/

destroy:
	@cd terraform/$(image)/asg && TF_VAR_dd_api_key=$(DD_API_KEY) terraform destroy -var-file=$(_ROOT_DIR)/terraform.tfvars
	@cd terraform/$(image)/elb && terraform destroy -var-file=$(_ROOT_DIR)/terraform.tfvars

