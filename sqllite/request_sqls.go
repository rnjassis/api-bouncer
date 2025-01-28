package sqllite

func getRequestsSql(isActive bool) SQL {
	query := `SELECT id, verb, url, active FROM request WHERE project_id = ?`
	if isActive {
		query += ` and active = true`
	}
	sql := SQL{sql: query}
	return sql
}

func getRequestByProjectUrlSql() SQL {
	sql := SQL{sql: `SELECT req.id, req.verb, req.url, req.active
				FROM request req
				INNER JOIN project proj on proj.id = req.project_id
				WHERE proj.name = ? and req.url = ?`}
	return sql
}

func createRequestSql() SQL {
	sql := SQL{`INSERT INTO request (project_id, verb, url, active)
					SELECT pr.id, ?, ?, ?
					FROM project pr
					WHERE pr.name = ?
				`}
	return sql
}
