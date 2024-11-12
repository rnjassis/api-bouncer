package sqllite

func getRequestsSql() SQL {
	sql := SQL{sql: `SELECT id, verb, url FROM request WHERE project_id = ?`}
	return sql
}

func getRequestByProjectUrlSql() SQL {
	sql := SQL{sql: `SELECT req.id, req.verb, req.url 
				FROM request req
				INNER JOIN project proj on proj.id = req.project_id
				WHERE proj.name = ? and req.url = ?`}
	return sql
}

func createRequestSql() SQL {
	sql := SQL{`INSERT INTO request (project_id, verb, url)
					SELECT pr.id, ?, ?
					FROM project pr
					WHERE pr.name = ?
				`}
	return sql
}
