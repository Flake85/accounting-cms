POSTGRES_HOST=localhost
POSTGRES_PORT=54321
POSTGRES_USER=cms
POSTGRES_PASSWORD=cms
POSTGRES_DB=cms

dockers:
	cd docker; docker-compose up

golang:
	cd server; go run main.go -POSTGRES_HOST=${POSTGRES_HOST} -POSTGRES_PORT=${POSTGRES_PORT} -POSTGRES_USER=${POSTGRES_USER} -POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -POSTGRES_DB=${POSTGRES_DB}

react:
	cd client; yarn && yarn run dev

run:
	make dockers && make golang && make react
