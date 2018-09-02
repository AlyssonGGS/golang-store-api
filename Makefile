build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o store .
	docker build -t store .

run:
	docker run -it -d -p 80:8080 store 