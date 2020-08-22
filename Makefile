build-mockgen:
	docker build -f mockgen.Dockerfile --tag reportingmockgen .

mockgen: build-mockgen
	docker run --rm --volume "$$(pwd):/src" reportingmockgen
