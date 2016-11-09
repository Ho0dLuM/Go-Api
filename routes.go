package main

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

type User struct {
	Id        int64  `db:"id" json:"id"`
	Firstname string `db:"first_name" json:"first_name"`
	Lastname  string `db:"last_name" json:"last_name"`
}

func main() {
	app := gin.Default()
	router := app.Group("api/v1")
	{
		router.GET("/users", GetUsers)
		router.GET("/users/:id", GetUser)
		router.POST("/users", PostUser)
		router.PUT("/users/:id", UpdateUser)
		router.DELETE("/users/:id", DeleteUser)
	}
	app.Run(":8080")
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("postgres", "postgres://Ho0dLuM:password@localhost/league_api?sslmode=disable")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetUsers(c *gin.Context) {
	var users []User
	_, err := dbmap.Select(&users, "SELECT * FROM users")
	if err == nil {
		c.JSON(200, users)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
}

func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=$1", id)
	if err == nil {
		user_id, _ := strconv.ParseInt(id, 0, 64)
		content := &User{
			Id:        user_id,
			Firstname: user.Firstname,
			Lastname:  user.Lastname,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func PostUser(c *gin.Context) {
	var user User
	c.Bind(&user)
	if user.Firstname != "" && user.Lastname != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO users (first_name, last_name) VALUES ($1, $2)`, user.Firstname, user.Lastname); insert != nil {
			c.JSON(201, gin.H{"success": "Added User"})
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=$1", id)
	if err == nil {
		var changed User
		c.Bind(&changed)

		if changed.Firstname == "" && changed.Lastname == "" {
			c.JSON(422, gin.H{"error": "fields are empty"})
		}

		if changed.Firstname != "" {
			user.Firstname = changed.Firstname
		}

		if changed.Lastname != "" {
			user.Lastname = changed.Lastname
		}

		if update, err2 := dbmap.Exec(`UPDATE users SET first_name=$1, last_name=$2 WHERE id=$3`, user.Firstname, user.Lastname, id); update != nil {
			c.JSON(200, user)
		} else {
			checkErr(err2, "Updated failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}

func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	err := dbmap.SelectOne(&user, "SELECT * FROM users WHERE id=$1", id)
	if err == nil {
		if delete, err2 := dbmap.Exec(`DELETE FROM users WHERE id=$1`, id); delete != nil {
			c.JSON(200, gin.H{"id #" + id: " deleted"})
		} else {
			checkErr(err2, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}
}
