package sqllite

func getResponseByRequestIdSql(isActive bool) SQL {
	query := `SELECT id, status_code, active, body, mime, identifier, redirect FROM response WHERE request_id = ?`
	if isActive {
		query += ` and active = true`
	}

	sql := SQL{sql: query}
	return sql
}

func getResponseByProjectRequestResponseSql() SQL {
	sql := SQL{sql: `SELECT id, status_code, active, body, mime, identifier, redirect
					FROM response resp
					INNER JOIN request req on req.id = resp.request_id
					INNER JOIN project proj on proj.id = req.project_id
					WHERE project.name = ? and req.url = ? and resp.identification = ?`}
	return sql
}
func createResponseSql() SQL {
	sql := SQL{sql: `INSERT INTO response (request_id, status_code, active, body, mime, identifier, redirect)
                        SELECT req.id, ?, ?, ?, ?, ?
                        FROM request req
                        INNER JOIN project prod on req.project_id = prod.id
                        WHERE prod.name = ? and req.url = ?`}
	return sql
}
