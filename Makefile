.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: test
test:
	ginkgo test -race ./...	

.PHONY: run
run:
	go run ./cmd/cnvalidator/main.go

.PHONY: run-docker
run-docker:
	docker compose down -v --remove-orphans 
	docker compose up --build

CERTS_DIR ?= certs
CA_DIR := $(CERTS_DIR)/ca
CA_UNTRUSTED_DIR := $(CERTS_DIR)/ca-untrusted
CLIENT_DIR := $(CERTS_DIR)/client
CLIENT_UNTRUSTED_DIR := $(CERTS_DIR)/client-untrusted
SERVER_DIR := $(CERTS_DIR)/server

DAYS := 365
KEY_SIZE := 2048
CN ?= authorized-client

.PHONY: generate-all-certs
generate-all-certs: clean-certs ca ca-untrusted client client-untrusted server

.PHONY: clean-certs
clean-certs:
	rm -rf $(CA_DIR) $(CA_UNTRUSTED_DIR) $(CLIENT_DIR) $(CLIENT_UNTRUSTED_DIR) $(SERVER_DIR)

.PHONY: ca
ca:
	@echo "==> Generating trusted CA"
	@rm -rf $(CA_DIR)
	@mkdir -p $(CA_DIR)
	openssl genrsa -out $(CA_DIR)/ca.key $(KEY_SIZE)
	openssl req -x509 -new -nodes -key $(CA_DIR)/ca.key -sha256 -days $(DAYS) \
		-subj "/CN=Trusted Root CA" \
		-out $(CA_DIR)/ca.crt

.PHONY: ca-untrusted
ca-untrusted:
	@echo "==> Generating untrusted CA"
	@rm -rf $(CA_UNTRUSTED_DIR)
	@mkdir -p $(CA_UNTRUSTED_DIR)
	openssl genrsa -out $(CA_UNTRUSTED_DIR)/ca.key $(KEY_SIZE)
	openssl req -x509 -new -nodes -key $(CA_UNTRUSTED_DIR)/ca.key -sha256 -days $(DAYS) \
		-subj "/CN=Untrusted Root CA" \
		-out $(CA_UNTRUSTED_DIR)/ca.crt

.PHONY: client
client:
	@echo "==> Generating client certificate (CN=$(CN))"
	@rm -rf $(CLIENT_DIR)
	@mkdir -p $(CLIENT_DIR)
	openssl genrsa -out $(CLIENT_DIR)/client.key $(KEY_SIZE)
	openssl req -new -key $(CLIENT_DIR)/client.key \
		-subj "/CN=$(CN)" \
		-out $(CLIENT_DIR)/client.csr
	openssl x509 -req -in $(CLIENT_DIR)/client.csr -CA $(CA_DIR)/ca.crt -CAkey $(CA_DIR)/ca.key \
		-CAcreateserial -out $(CLIENT_DIR)/client.crt -days $(DAYS) -sha256

.PHONY: client-untrusted
client-untrusted:
	@echo "==> Generating client cert signed by UNTRUSTED CA"
	@rm -rf $(CLIENT_UNTRUSTED_DIR)
	@mkdir -p $(CLIENT_UNTRUSTED_DIR)
	openssl genrsa -out $(CLIENT_UNTRUSTED_DIR)/client.key $(KEY_SIZE)
	openssl req -new -key $(CLIENT_UNTRUSTED_DIR)/client.key \
		-subj "/CN=authorized-client" \
		-out $(CLIENT_UNTRUSTED_DIR)/client.csr
	openssl x509 -req -in $(CLIENT_UNTRUSTED_DIR)/client.csr -CA $(CA_UNTRUSTED_DIR)/ca.crt -CAkey $(CA_UNTRUSTED_DIR)/ca.key \
		-CAcreateserial -out $(CLIENT_UNTRUSTED_DIR)/client.crt -days $(DAYS) -sha256

.PHONY: server
server:
	@echo "==> Generating server certificate with SAN"
	@rm -f $(SERVER_DIR)
	@mkdir -p $(SERVER_DIR)
	openssl genrsa -out $(SERVER_DIR)/server.key $(KEY_SIZE)
	openssl req -new -key $(SERVER_DIR)/server.key \
		-config $(CERTS_DIR)/openssl.cnf \
		-out $(SERVER_DIR)/server.csr
	openssl x509 -req -in $(SERVER_DIR)/server.csr -CA $(CA_DIR)/ca.crt -CAkey $(CA_DIR)/ca.key \
		-CAcreateserial -out $(SERVER_DIR)/server.crt -days $(DAYS) -sha256 \
		-extfile $(CERTS_DIR)/openssl.cnf -extensions v3_req