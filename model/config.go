package model

type Config struct {
	Token            string `json:"token"`
	PostgresLogin    string `json:"postgres_login"`
	PostgresPassword string `json:"postgres_password"`
	PostgresDbName   string `json:"postgres_db_name"`
	PostgresPort     string `json:"postgres_port"`
}
