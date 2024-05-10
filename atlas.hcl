env "local" {
  src = "file://server/service/repo/db/sqliteimpl/sqlc/schema"
  url = "sqlite://data/sqlite/telemetry.db?_fk=1"
  dev = "sqlite://file?mode=memory"
  migration {
    dir = "file://server/service/repo/db/sqliteimpl/sqlc/migrations"
  }
}
