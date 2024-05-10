-- name: InsertUser :execresult
INSERT INTO users (id) VALUES (?);

-- name: UpsertUser :execresult
INSERT INTO users (id) VALUES (?)
ON CONFLICT DO NOTHING;
