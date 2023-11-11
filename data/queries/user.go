package queries

func GetUserDataQuery() string {
	return `SELECT 
				u.id, u.login, u.hash
			from users u 
			where u.login = ?`
}
