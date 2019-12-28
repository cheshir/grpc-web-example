install:
	docker run -it --rm --entrypoint npm -v `pwd`:/app -w /app node install
build:
	docker run -it --rm --entrypoint npx -v `pwd`:/app -w /app node webpack client.js
