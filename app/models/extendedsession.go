package models

func UsersSessionBySessionID(db XODB, session_id string) (*UsersSession, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, user_id, modified, created ` +
		`FROM public.users_sessions ` +
		`WHERE session_id = $1`

	// run query
	XOLog(sqlstr, session_id)
	us := UsersSession{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, session_id).Scan(&us.ID, &us.UserID, &us.Modified, &us.Created)
	if err != nil {
		return nil, err
	}

	return &us, nil
}
