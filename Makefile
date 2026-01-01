.PHONY: run dev

run: dev

dev:
	docker build -t oswingsonic . && \
	docker run -p 1991:1991 -v ./:/app/ oswingsonic
