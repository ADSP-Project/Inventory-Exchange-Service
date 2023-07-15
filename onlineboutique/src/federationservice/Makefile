ui:
	docker build . -f Dockerfile-UI -t ui
	docker run --rm -it --net=host ui
db:
	docker build . -f Dockerfile-DB -t db
	docker run --rm -it --net=host db
be:
	docker build . -f Dockerfile-BE -t be
	docker run --rm -it --net=host be
