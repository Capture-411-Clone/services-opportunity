export ROOT_DIR=$(realpath $(dir $(firstword $(MAKEFILE_LIST))))

docker-test:
	chmod +x $(ROOT_DIR)/scripts/test.sh
	$(ROOT_DIR)/scripts/test.sh docker

docker-lint:
	chmod +x $(ROOT_DIR)/scripts/lint.sh
	$(ROOT_DIR)/scripts/lint.sh docker

docker-cli:
	$(eval ARGS := $(filter-out $@,$(MAKECMDGOALS)))
	-docker exec dev_capture411_opportunity_service go run /app/cmd/chef/main.go $(ARGS)

swagger:
	$(ROOT_DIR)/scripts/openapi.sh

# -------------- Development ----------------
docker-start:
	chmod +x $(ROOT_DIR)/scripts/start.sh
	$(ROOT_DIR)/scripts/start.sh docker

docker-stop:
	chmod +x $(ROOT_DIR)/scripts/stop.sh
	$(ROOT_DIR)/scripts/stop.sh

docker-reload:
	make docker-stop && make docker-start

# TODO: add volume clean up for database if needed after container shutdown
docker-restart:
	chmod +x $(ROOT_DIR)/scripts/restart.sh
	$(ROOT_DIR)/scripts/restart.sh
# -------------- Production ----------------
docker-serve:
	chmod +x $(ROOT_DIR)/scripts/serve.sh
	$(ROOT_DIR)/scripts/serve.sh docker

docker-drop:
	chmod +x $(ROOT_DIR)/scripts/drop.sh
	$(ROOT_DIR)/scripts/drop.sh

docker-reboot:
	chmod +x $(ROOT_DIR)/scripts/reboot.sh
	$(ROOT_DIR)/scripts/reboot.sh
# ------------------ Docker -----------------
dev-logs:
	chmod +x $(ROOT_DIR)/scripts/logs-dev.sh
	$(eval ARGS := $(filter-out $@,$(MAKECMDGOALS)))
	$(ROOT_DIR)/scripts/logs-dev.sh $(ARGS)
# ------------------- Git -------------------
prepare-hooks:
	chmod +x $(ROOT_DIR)/scripts/prepare-hooks.sh
	$(ROOT_DIR)/scripts/prepare-hooks.sh

# ------------------- Git -------------------
# this can be used for easier with token login
# stripe login --api-key sk_test_51OOsIJHiup0sF8jzXlcoCPUetHNAO04DAt
stripe-login:
	stripe login
stripe-listen:
	stripe listen --forward-to localhost:7878/api/v1/stripe/bank/webhooks