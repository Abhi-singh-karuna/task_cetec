package main

import (
	"log"

	"github.com/Abhi-singh-karuna/config"
	"github.com/Abhi-singh-karuna/router"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// initilize the gin
	app := gin.Default()

	// stablish connection with database
	config.ConnectDB()

	Create_table()
	
	// load the environment file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router.Routes(app)

	// Run attaches the router to a http.Server and starts listening and serving HTTP requests.
	log.Fatal(app.Run(":8000"))
}

func Create_table() {
	db := config.ConnectDB()
	db.Exec("create table person(id int not null key auto_increment,  name varchar(255), age int)")
	db.Exec("insert into person(id, name, age) values (1, 'mike', 31), (2, 'John', 20), (3, 'Joseph', 20)")

	db.Exec("create table phone(id int not null key auto_increment,  number varchar(255), person_id int)")
	db.Exec("insert into phone(id, person_id, number) values (1,1, '444-444-4444'), (8,2, '123-444-7777'), (3,3, '445-222-1234')")

	db.Exec("create table address(id int not null key auto_increment,  city varchar(255), state varchar(255), street1 varchar(255), street2 varchar(255), zip_code varchar(255))")
	db.Exec("insert into address(id ,  city , state , street1 , street2 , zip_code ) values (1,'Eugene', 'OR', '111 Main St', '', '98765'), (2, 'Sacramento', 'CA', '432 First St', 'Apt 1', '22221'), (3, 'Austin', 'TX', '213 South 1st St', '', '78704')")

	db.Exec("create table address_join(id int not null key auto_increment,  person_id int, address_id int)")
	db.Exec("insert into address_join(id, person_id, address_id) values (1,1,3),(2,2,1),(3,3,2)")

}
