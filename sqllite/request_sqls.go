package sqllite

func getRequestsSql() SQL {
	sql := SQL{sql: `SELECT id, verb, url FROM request WHERE project_id = ?`}
	return sql
}
