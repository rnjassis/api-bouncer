package sqllite

func getResponseSQL() SQL {
    sql := SQL{sql: `SELECT id, status_code, active, body, mime FROM response WHERE id = ?`}
    return sql
}

