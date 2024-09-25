package sqllite

func createProjectTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS project (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"port" integer NOT NULL,
	"name" TEXT NOT NULL UNIQUE,
	"description" TEXT NOT NULL
	)`}
	return sql
}

func createRequestTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS request (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"project_id" integer NOT NULL,
	"verb" TEXT,
	"url" TEXT,
	FOREIGN KEY("project_id") REFERENCES project("id"),
	)`}
	return sql
}

func createResponseTableSql() SQL {
    sql := SQL{sql: `CREATE TABLE IF NOT EXISTS response (
        "id" integer NOT NULL,
        "request_id" integer NOT NULL,
        "status_code" integer NOT NULL,
        "active" integer NOT NULL,
        "body" TEXT,
        "mime" TEXT
        FOREIGN KEY("request_id") REFERENCES request("id")
    )`}
    return sql
}

