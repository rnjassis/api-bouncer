package sqllite

func getRequestsSql() SQL {
	sql := SQL{sql: `SELECT id, verb, url, return FROM request WHERE projectId = ?`}
	return sql
}
