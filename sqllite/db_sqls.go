package sqllite

func createProjectTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS project (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"port" integer NOT NULL,
	"name" TEXT NOT NULL,
	"description" TEXT NOT NULL
	)`}
	return sql
}

func createEndpointTableSql() SQL {
	sql := SQL{sql: `CREATE TABLE IF NOT EXISTS endpoint (
	"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"projectId" integer NOT NULL,
	"verb" TEXT,
	"url" TEXT,
	"responseId" integer,
	FOREIGN KEY("projectId") REFERENCES project("id"),
    FOREIGN KEY("responseId") REFERENCES response("id)
	)`}
	return sql
}

func createResponseTableSql() SQL {
    sql := SQL{sql: `CREATE TABLE IF NOT EXISTS response (
        "id" integer NOT NULL,
        "statusCode" integer NOT NULL,
        "body" TEXT,
    )`}
    return sql
}

