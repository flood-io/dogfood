.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

_ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

install: ## Install target dependencies
	@echo "Installing target dependencies for $(target)"
	@cd dockerfiles/$(target) && docker build -t floodio/$(target) .
	@docker push floodio/$(target)
	@cd terraform/$(target)/elb && terraform apply -var-file=$(_ROOT_DIR)/terraform.tfvars
	@cd terraform/$(target)/asg && TF_VAR_dd_api_key=$(DD_API_KEY) terraform apply -var-file=$(_ROOT_DIR)/terraform.tfvars

health: ## Check target health
	@aws --profile=flooded --region us-west-2 elb describe-instance-health --load-balancer-name $(target)-elb | jq -r '.InstanceStates[] | .State' | sort | uniq -c | sort
	$(eval ELB_DNS_NAME := $(shell terraform output -state=terraform/$(target)/elb/terraform.tfstate dns_name))
	@echo "Checking ELB at $(ELB_DNS_NAME) for $(target)"
	@curl -I --silent --connect-timeout 3 http://$(ELB_DNS_NAME)/

destroy: ## Check target dependencies
	@cd terraform/$(target)/asg && TF_VAR_dd_api_key=$(DD_API_KEY) terraform destroy -var-file=$(_ROOT_DIR)/terraform.tfvars
	@cd terraform/$(target)/elb && terraform destroy -var-file=$(_ROOT_DIR)/terraform.tfvars

test: ## Test flood-chrome locally
	@docker run -v ${_ROOT_DIR}/spec/flood-chrome:/test -it --rm floodio/chrome bash -c "npm run build /test/$(target).json"

src:
	@rm -rf dockerfiles/dogfood/src
	@swagger-codegen generate \
		-i dockerfiles/dogfood/html/swagger.json \
		-l go-server \
		-o dockerfiles/dogfood/src
	@cd dockerfiles/dogfood/src/go && sed -i '' -e 's/package/package dogfood/g' *.go
