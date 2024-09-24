package sqllite

func getResponseSQL() SQL {
    sql := SQL{sql: `SELECT id, statusCode, body FROM response WHERE id = ?`}
    return sql
}

