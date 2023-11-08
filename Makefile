
start:
	cd cardinal; go mod vendor
	docker compose up nakama cardinal --build cardinal
