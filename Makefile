setup-services:
	docker-compose -f docker-compose.yml up -d --build
delete-services:
	docker-compose down -v
