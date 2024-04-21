env "local" {
	src = "ent://server/service/repo/db/entimpl/ent/schema"
	url = "postgres://root:root@localhost:5432/telemetry?search_path=public&sslmode=disable"
	dev = "docker://postgres/16/dev?search_path=public"

	migration {
		dir = "file://server/service/repo/db/entimpl/ent/migrate/migrations"
	}
}
