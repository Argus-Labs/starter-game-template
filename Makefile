

clear:
	docker compose down --volumes

start:
	cd cardinal; go mod vendor
	docker compose up nakama cardinal --build cardinal

test:
	cd cardinal; go mod vendor
	make clear
	docker compose up --abort-on-container-exit --exit-code-from testsuite --attach testsuite
