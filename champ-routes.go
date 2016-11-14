package main

import (
	"database/sql"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"gopkg.in/gorp.v1"
)

type Champion struct {
	Id                 int64  `db:"id" json:"id"`
	Name               string `db:"name" json:"name"`
	Image              string `db:"image" json:"image"`
	Title              string `db:"title" json:"title"`
	Enemytips          string `db:"enemytips" json:"enemytips"`
	Lore               string `db:"lore" json:"lore"`
	Passivename        string `db:"passive_name" json:"passiveName"`
	Passiveimage       string `db:"passive_image" json:"passiveImage"`
	Passivedescription string `db:"passive_description" json:"passiveDescription"`
	Spellsqname        string `db:"spells_q_name" json:"spellsQname"`
	Spellsqimage       string `db:"spells_q_image" json:"spellsQimage"`
	Spellsqdescription string `db:"spells_q_description" json:"spellsQdescription"`
	Spellswname        string `db:"spells_w_name" json:"spellsWname"`
	Spellswimage       string `db:"spells_w_image" json:"spellsWimage"`
	Spellswdescription string `db:"spells_w_description" json:"spellsWdescription"`
	Spellsename        string `db:"spells_e_name" json:"spellsEname"`
	Spellseimage       string `db:"spells_e_image" json:"spellsEimage"`
	Spellsedescription string `db:"spells_e_description" json:"spellsEdescription"`
	Spellsrname        string `db:"spells_r_name" json:"spellsRname"`
	Spellsrimage       string `db:"spells_r_image" json:"spellsRimage"`
	Spellsrdescription string `db:"spells_r_description" json:"spellsRdescription"`
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	port := os.Getenv("PORT")
	app := gin.Default()
	app.Use(CORSMiddleware())
	router := app.Group("api/v1")
	{
		router.GET("/champions", GetChampions)
		router.GET("/champions/:id", GetChampion)
		router.POST("/champions", PostChampion)
		router.DELETE("/champions/:id", DeleteChampion)
		// router.PUT("/champions/:id", UpdateChampion)
	}
	app.Run(":" + port)
	// app.Run(":8080")
}

var dbmap = initDb()

func initDb() *gorp.DbMap {
	url := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", url)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}

	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func GetChampions(c *gin.Context) {
	var champions []Champion
	_, err := dbmap.Select(&champions, "SELECT * FROM champions")
	if err == nil {
		c.JSON(200, champions)
	} else {
		c.JSON(404, gin.H{"error": "no user(s) into the table"})
	}
}

func GetChampion(c *gin.Context) {
	id := c.Params.ByName("id")
	var champion Champion
	err := dbmap.SelectOne(&champion, "SELECT * FROM champions WHERE id=$1", id)
	if err == nil {
		champion_id, _ := strconv.ParseInt(id, 0, 64)
		content := &Champion{
			Id:                 champion_id,
			Name:               champion.Name,
			Image:              champion.Image,
			Title:              champion.Title,
			Enemytips:          champion.Enemytips,
			Lore:               champion.Lore,
			Passivename:        champion.Passivename,
			Passiveimage:       champion.Passiveimage,
			Passivedescription: champion.Passivedescription,
			Spellsqname:        champion.Spellsqname,
			Spellsqimage:       champion.Spellsqimage,
			Spellsqdescription: champion.Spellsqdescription,
			Spellswname:        champion.Spellswname,
			Spellswimage:       champion.Spellswimage,
			Spellswdescription: champion.Spellswdescription,
			Spellsename:        champion.Spellsename,
			Spellseimage:       champion.Spellseimage,
			Spellsedescription: champion.Spellsedescription,
			Spellsrname:        champion.Spellsrname,
			Spellsrimage:       champion.Spellsrimage,
			Spellsrdescription: champion.Spellsrdescription,
		}
		c.JSON(200, content)
	} else {
		c.JSON(404, gin.H{"error": "champion not found"})
	}
}

func PostChampion(c *gin.Context) {
	var champion Champion
	c.Bind(&champion)
	if champion.Name != "" && champion.Image != "" && champion.Title != "" && champion.Enemytips != "" && champion.Lore != "" && champion.Passivename != "" && champion.Passiveimage != "" && champion.Passivedescription != "" && champion.Spellsqname != "" && champion.Spellsqimage != "" && champion.Spellsqdescription != "" && champion.Spellswname != "" && champion.Spellswimage != "" && champion.Spellswdescription != "" && champion.Spellsename != "" && champion.Spellseimage != "" && champion.Spellsedescription != "" && champion.Spellsrname != "" && champion.Spellsrimage != "" && champion.Spellsrdescription != "" {
		if insert, _ := dbmap.Exec(`INSERT INTO champions (name, image, title, enemytips, lore, passive_name, passive_image, passive_description, spells_q_name, spells_q_image, spells_q_description, spells_w_name, spells_w_image, spells_w_description, spells_e_name, spells_e_image, spells_e_description, spells_r_name, spells_R_image, spells_r_description) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)`, champion.Name, champion.Image, champion.Title, champion.Enemytips, champion.Lore, champion.Passivename, champion.Passiveimage, champion.Passivedescription, champion.Spellsqname, champion.Spellsqimage, champion.Spellsqdescription, champion.Spellswname, champion.Spellswimage, champion.Spellswdescription, champion.Spellsename, champion.Spellseimage, champion.Spellsedescription, champion.Spellsrname, champion.Spellsrimage, champion.Spellsrdescription); insert != nil {
			c.JSON(201, gin.H{"success": "Added Champion"})
		}
	} else {
		c.JSON(422, gin.H{"error": "fields are empty"})
	}
}

func DeleteChampion(c *gin.Context) {
	id := c.Params.ByName("id")
	var champion Champion
	err := dbmap.SelectOne(&champion, "SELECT * FROM champions WHERE id=$1", id)
	if err == nil {
		if delete, err2 := dbmap.Exec(`DELETE FROM champions WHERE id=$1`, id); delete != nil {
			c.JSON(200, gin.H{"id #" + id: " deleted"})
		} else {
			checkErr(err2, "Delete failed")
		}
	} else {
		c.JSON(404, gin.H{"error": "champion not found"})
	}
}

// func UpdateChampion(c *gin.Context) {
// 	id := c.Params.ByName("id")
// 	var champion Champion
// 	err := dbmap.SelectOne(&champion, "SELECT * FROM champions WHERE id=$1", id)
// 	if err == nil {
// 		var changed Champion
// 		c.Bind(&changed)
//
// 		if changed.Firstname == "" && changed.Lastname == "" {
// 			c.JSON(422, gin.H{"error": "fields are empty"})
// 		}
//
// 		if changed.Firstname != "" {
// 			champion.Firstname = changed.Firstname
// 		}
//
// 		if changed.Lastname != "" {
// 			champion.Lastname = changed.Lastname
// 		}
//
// 		if update, err2 := dbmap.Exec(`UPDATE champions SET first_name=$1, last_name=$2 WHERE id=$3`, champion.Firstname, champion.Lastname, id); update != nil {
// 			c.JSON(200, champion)
// 		} else {
// 			checkErr(err2, "Updated failed")
// 		}
// 	} else {
// 		c.JSON(404, gin.H{"error": "champion not found"})
// 	}
// }
