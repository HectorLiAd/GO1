package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm" 
)

type Person struct {
    gorm.Model
    Name string
    Age  string
}

func main() {




	dsn := "docker:docker@tcp(db:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Person{})

	//s := Person{ Name: "Sean" , Age: 50 }
	//s.Name = "Jeson";

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hola HOLA",
		})
	})

	// MOSTRAR A TODAS LAS PERSONAS
	r.GET("/persons", func(c *gin.Context) {//GET ALL
		var lis []Person
		db.Find(&lis)
		c.JSON(http.StatusOK, lis)
	})

	// MOSTRAR PERSONA EN ESPECIFICO
	r.GET("/persons/:id", func(c *gin.Context) {//GET ESPECIFICO
		id := c.Param("id") /// permiten poner el 
		var d Person
		if err := db.First(&d, id).Error; err != nil {
	    	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		      "error": err.Error(),
				    })
			return
		}

		c.JSON(http.StatusOK, &d)
	})

	// ELIMINAR PERSONA POR ID
	r.DELETE("/persons/:id", func(c *gin.Context) {
		id := c.Param("id")
		var d Person
		if err := db.Where("id = ?", id).First(&d).Error; err != nil {
			// c.AbortWithStatus(http.StatusNotFound)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	      	"error": err.Error(),
			    })
			return
		}
		db.Unscoped().Delete(&d)
	})

	// ACTUALIZAR PERSONA POR ID
	r.PUT("/persons/:id", func(c *gin.Context) {//GET ESPECIFICO
		id := c.Param("id")
		var d Person
		if err := db.First(&d, id).Error; err != nil {
	    	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	      "error": err.Error(),
			    })
			return
		}
		d.Name = c.PostForm("name") Name		avdvds
		d.Age = c.PostForm("age")   Age  		16    			guardar
		db.Save(&d)
		c.JSON(http.StatusOK, &d)
	})


	// CREAR PERSONA
	r.POST("/persons/", func(c *gin.Context) {
		
		d := Person{Name: c.PostForm("name"), Age: c.PostForm("age")}
		db.Create(&d)

		c.JSON(http.StatusOK, &d)
		/*d := Person{Name: c.PostForm("name"), Age: c.PostForm("age")}

		if err := c.BindJSON(&d); err != nil {
	    	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
	      "error": err.Error(),
			    })
			return
		}
		db.Create(&d)
		c.JSON(http.StatusOK, &d)*/

	})

	/*r.POST("/person/", func(c *gin.Context) {
		d := Person{ Name: c.PostForm("name"), Age: c.PostForm("age") }
		db.Create(&d)
		id	 := c.Query("id")//Query recibe valores
		name_ := c.PostForm("name")// ->Pedimos el name 
		lastname_ := c.PostForm("lastname")// ->Pedimos el lastname
		lastname := c.PostForm("lastname")
		c.JSON(200, gin.H{ // serializador de gin
			"name ": d.Name,
			"age": d.Age,
			"name":     name,
			"lastname": lastname,
		})
	})*/

	r.Run(":8087") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
