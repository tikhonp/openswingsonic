.PHONY: run dev

run: dev

dev:
	docker build -t oswingsonic --build-arg DEBUG=true --build-arg SWINGSONIC_BASE_URL=${SWINGSONIC_BASE_URL} . && docker run -p 1991:1991 -v ./:/app/ -v ./users:/app/users oswingsonic
