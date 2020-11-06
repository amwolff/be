web-dev:
	docker run -p 1234:1234 -p 1235:1235 -it --rm --name experimental-web-dev \
	-v /home/artur/devel/be/apps/experimental/web/app/:/usr/src/app \
	-w /usr/src/app \
	node:latest npm run dev

web-build:
	docker run -it --rm --name experimental-web-build \
	-v /home/artur/devel/be/apps/experimental/web/app/:/usr/src/app \
	-w /usr/src/app \
	node:latest npm run build
