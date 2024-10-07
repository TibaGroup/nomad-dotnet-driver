PLUGIN_BINARY=nomad-dotnet-driver
TASK_PKG=test_nomad_task.zip
export GO111MODULE=on

default: build

.PHONY: clean build build_test pkg_test serve_test clean_test
clean: ## Remove build artifacts
	rm -rf ${PLUGIN_BINARY}

build:
	go build -o ${PLUGIN_BINARY} .

build_test:
	dotnet build ./test-resources/TestNomadTask --configuration Release

pkg_test: build_test
	zip -rjv test_nomad_task.zip ./test-resources/TestNomadTask/bin/Release/net7.0

serve_test: test_nomad_task.zip ## Serve built and zipped test .Net project
	python3 -m http.server 80

clean_test: test_nomad_task.zip ## Remove zipped test .Net project
	rm -rf ${TASK_PKG}