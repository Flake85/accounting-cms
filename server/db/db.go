package db

func BuildDbConnectionStr(host string, port string, user string, db string, password string) string {
	return "host="+host+" port="+port+" user="+user+" dbname="+db+" password="+password+" sslmode=disable"
}