package users_db

import (
	"database/sql"
	"fmt"
	"log"
	//"os"

	_ "github.com/go-sql-driver/mysql"
)
//para configurar variables de entorno
/*const(
	mysql_users_username	= "mysql_users_username"
	mysql_users_password	= "mysql_users_password"
	mysql_users_host		= "mysql_users_host"
	mysql_users_shema		= "mysql_users_shema"

)*/
var(
	Client *sql.DB

	/*username 	= os.Getenv(mysql_users_username)
	password 	= os.Getenv(mysql_users_password)
	host 		= os.Getenv(mysql_users_host)
	schema 		= os.Getenv(mysql_users_shema)*/
)

func init(){
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"123456",
		"127.0.0.1:3306",
		"users_db",
	)
	var err error
	Client, err = sql.Open("mysql",dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil{
		panic(err)
	}
	log.Println("Configuracion de base de datos exitoso")
}