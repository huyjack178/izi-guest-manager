package mysqldb


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type ConnectOpt struct {
	UserName string
	Database string
	Password string
	Host     string
}

type Instance  struct {
	opts ConnectOpt
	conn *sql.DB
}

func NewInstance(opts ConnectOpt) (ins *Instance, err error) {
	ins = &Instance{
		opts: opts,
	}

	connectionStr := opts.UserName + ":" + opts.Password + "@tcp(" + opts.Host + ")/" + opts.Database
	log.Println(connectionStr)
	ins.conn, err = sql.Open("mysql", connectionStr)

	return ins, err
}

func(this*Instance) Test(){
	rows, err := this.conn.Query("SELECT * FROM guest", 1)

	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	log.Println(rows)

}





