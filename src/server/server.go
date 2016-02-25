package server

import (
	"net/http"
)

type Config struct {
	Server struct {
		Port string `json:"API_PORT"`
		Addr string `json:"API_ADDR"`
	} `json:"server"`

	Mongo struct{
		Port string `json:"MONGODB_PORT"`
		Addr string `json:"MONGODB_ADDR"`
		DBName string `json:"MONGODB_DBNAME"`
		Collections map[string]string `json:"MONGODB_COLLECTIONS"`
	} `json:"mongodb"`

	MySQL struct{
		UserName string `json:"MYSQL_USERNAME"`
		Database string `json:"MYSQL_DATABASE"`
		Password string `json:"MYSQL_PASSWORD"`
		Host string `json:"MYSQL_HOST"`
	} `json:"mysql"`
}

func Start(cfg Config) {
	s := setup(cfg)

	listenAddr := cfg.Server.Addr + ":" + cfg.Server.Port
	log.Println("server is listening on", listenAddr)
	http.ListenAndServe(listenAddr, s.Handler)

}
