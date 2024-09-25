package sqllite

func getProjectsSql() SQL {
	sql := SQL{sql: `SELECT id, port, name, description FROM project`}
	return sql
}

func getProjectByNameSql() SQL {
	sql := SQL{sql: `SELECT id, port, name, description FROM project WHERE name = ?`}
	return sql
}
