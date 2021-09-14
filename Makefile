# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

tidy: ## Clean up inputs
	go mod tidy

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

BIN_DIR=$(shell pwd)/bin
test: tidy ## Run tests against a local `kube-apiserver` and `etcd`
	mkdir -p ${BIN_DIR}
	test -f ${BIN_DIR}/setup-envtest.sh || curl -sSLo ${BIN_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/v0.8.3/hack/setup-envtest.sh
	source ${BIN_DIR}/setup-envtest.sh; fetch_envtest_tools "$(BIN_DIR)/.."; setup_envtest_env $(BIN_DIR);
	export PATH="${BIN_DIR}:${PATH}"
	KUBEBUILDER_ASSETS=$(BIN_DIR) kubectl kuttl test --config test/kuttl-test.yaml

build: tidy vet
	mkdir -p ${BIN_DIR}
	go build -a -o ${BIN_DIR}/main main.go