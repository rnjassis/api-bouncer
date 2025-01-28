package sqllite

func createProjectTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS project (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"port" TEXT NOT NULL,
	"name" TEXT NOT NULL UNIQUE,
	"description" TEXT NOT NULL
	)`}
	return sql
}

// TODO change verb to method
func createRequestTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS request (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"project_id" integer NOT NULL,
	"verb" TEXT,
	"url" TEXT,
	"active" integer NOT NULL,
	FOREIGN KEY("project_id") REFERENCES project("id")
	)`}
	return sql
}

func createResponseTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS response (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "request_id" integer NOT NULL,
        "status_code" integer NOT NULL,
        "active" integer NOT NULL,
        "body" TEXT,
        "mime" TEXT,
		"identifier" TEXT NOT NULL UNIQUE,
        FOREIGN KEY("request_id") REFERENCES request("id")
    )`}
	return sql
}
