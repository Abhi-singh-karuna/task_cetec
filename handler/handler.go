package handler

import (
	"github.com/Abhi-singh-karuna/config"
	"github.com/Abhi-singh-karuna/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetData(c *gin.Context) {

	id := c.Param("id")
	db := config.ConnectDB()
	Sel_DB, err := db.Query("select person.name , phone.number ,address.city , address.state , address.street1 ,address.street2 ,address.zip_code  from address_join   join address on address_join.address_id = address.id join person on address_join.person_id = person.id join phone on person.id  = phone.person_id where person.id = " + id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed in find the data from database",
		})
		return
	}

	defer Sel_DB.Close()

	var data model.Info
	for Sel_DB.Next() {
		err = Sel_DB.Scan(&data.Name, &data.Number, &data.City, &data.State, &data.Street1, &data.Street2, &data.Zip_code)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  0,
				"message": err,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":       200,
		"responseData": data,
	})

}

func CreateData(c *gin.Context) {
	var s model.Info

	if err := c.ShouldBind(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  0,
			"error":   err,
			"message": "Error in binding data",
		})
		return
	}

	db := config.ConnectDB()

	insert_person, err := db.Query("INSERT INTO person(name) VALUES(?)", s.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  0,
			"message": "insertion in person is failed",
		})
	}
	defer insert_person.Close()

	insert_phone, err := db.Query("INSERT INTO phone(number) VALUES(?)", s.Number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  0,
			"message": "insertion in phone is failed",
		})
	}
	defer insert_phone.Close()

	insert_address, err := db.Query("INSERT INTO address(city,state,street1,street2,zip_code) VALUES(?,?,?,?,?)", s.City, s.State, s.Street1, s.Street2, s.Zip_code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  0,
			"message": "insertion in address is failed",
		})
	}
	defer insert_address.Close()

	c.JSON(http.StatusOK, gin.H{
		"status":       200,
		"responseData": s,
	})

}
