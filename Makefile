define docker_compose
	@if docker compose version >/dev/null 2>&1; then \
		docker compose $(1); \
	else \
		docker-compose $(1); \
	fi
endef

DEV_SERVICES=webapp mysql traefik

dev:
	$(call docker_compose,up $(DEV_SERVICES))

start:
	$(call docker_compose,up -d --force-recreate)

stop:
	$(call docker_compose,down --remove-orphans)

restart: stop start
