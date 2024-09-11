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
	"return" TEXT,
	FOREIGN KEY("projectId") REFERENCES project("id")
	)`}
	return sql
}
