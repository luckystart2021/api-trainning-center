package database

// CreateSubscribeFriendByRequestorAndTarget executes subscribe to updates from an email address
func (db Database) CreateUserByEmail(us, pw, em, role string) error {
	query := `INSERT INTO user
	(username, "password", email, "role")
	VALUES($1, $2, $3, $4);
	;`
	_, err := db.Conn.Exec(query, us, pw, em, role)
	if err != nil {
		return err
	}
	return nil
}
