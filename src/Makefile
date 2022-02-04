all: test vet fmt lint build
.PHONY: all lib test clean

test:
	go test ./...

vet:
	go vet ./...

build:
	go build ./internal/he
	go build -buildmode=c-shared -o lib/_simulator.so ./cmd/export

package:
	python3 setup.py bdist_wheel
	#python3 setup.py sdist

clean:
	rm -r *.egg-info build dist
