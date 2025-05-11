.PHONY: run-backend run-frontend run-all

run-backend:
	cd backend && go run cmd/server/main.go

run-frontend:
	cd frontend && npm start

run-all:
	make run-backend & make run-frontend 