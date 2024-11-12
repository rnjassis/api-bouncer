package sqllite

func getResponseSQL() SQL {
	sql := SQL{sql: `SELECT id, status_code, active, body, mime FROM response WHERE id = ?`}
	return sql
}

func createResponse() SQL {
	sql := SQL{sql: `INSERT INTO response (request_id, status_code, active, body, mime)
                        SELECT req.id, ?, ?, ?, ? 
                        FROM request req
                        INNER JOIN project prod on req.project_id = prod.id
                        WHERE prod.name = ? and req.url = ?`}
	return sql
}
