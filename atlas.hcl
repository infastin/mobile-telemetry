env "local" {
  src = "file://server/service/repo/db/sqlcimpl/sqlc/schema"
  url = "postgres://root:root@localhost:5432/telemetry?search_path=public&sslmode=disable"
  dev = "docker://postgres/16/dev?search_path=public"

  migration {
    dir = "file://server/service/repo/db/sqlcimpl/sqlc/migrations"
  }
}
