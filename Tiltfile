# ------------------------------------------------------------------------------
# DOCKER COMPOSE SERVICES
# ------------------------------------------------------------------------------

docker_compose("docker-compose.yml")

dc_resource(
    "postgresql",
    labels=["dependencies"],
)

# ------------------------------------------------------------------------------
# LOCAL SERVICES
# ------------------------------------------------------------------------------

# Main service
local_resource(
    "main",
    serve_cmd="make go/run",
    resource_deps=["postgresql-ready"],
    allow_parallel=True,
    labels=["services"],
)

# React frontend
local_resource(
    "ui",
    serve_cmd="make ui/dev",
    links=["http://localhost:5173"],
    resource_deps=["main-ready"],
    allow_parallel=True,
    labels=["services"],
)

# ------------------------------------------------------------------------------
# READINESS CHECKS
# ------------------------------------------------------------------------------

local_resource(
    "postgresql-ready",
    cmd="until docker compose exec -T postgresql pg_isready -U postgres; do sleep 1; done",
    resource_deps=["postgresql"],
    allow_parallel=True,
    labels=["readiness"],
)

local_resource(
    "main-ready",
    cmd="make wait WAIT_FOR=localhost:8080",
    resource_deps=["main"],
    allow_parallel=True,
    labels=["readiness"],
)
