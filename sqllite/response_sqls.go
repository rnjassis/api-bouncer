package sqllite

func getResponseSQL() SQL {
    sql := SQL{sql: `SELECT id, status_code, body, mime FROM response WHERE id = ?`}
    return sql
}

