package sqllite

func getEndpointsSql() SQL {
	sql := SQL{sql: `SELECT id, verb, url, return FROM endpoint WHERE projectId = ?`}
	return sql
}
