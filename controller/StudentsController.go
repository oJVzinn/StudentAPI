package controller

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"studentapi/database"
	"studentapi/person"
)

var engine *gin.Engine

func Setup(ginEngine *gin.Engine) {
	engine = ginEngine
	listStudents()
	addStudent()
	getStudent()
	updateStudent()
	deleteStudent()
}

func listStudents() {
	engine.GET("api/students", func(c *gin.Context) {
		students, err := database.GetAll()
		stats := http.StatusOK
		if err != nil {
			stats = http.StatusInternalServerError
			students = nil
		}

		c.JSON(stats, gin.H{
			"stats":    stats,
			"response": students,
		})
	})
}

func getStudent() {
	engine.GET("api/students/:studentid", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("studentid"))
		stats := http.StatusOK
		if err != nil {
			stats = http.StatusBadRequest
		}

		person := findStudent(id)
		if person == nil {
			stats = http.StatusNotFound
		}

		c.JSON(stats, gin.H{
			"stats":    stats,
			"response": person,
		})
	})
}

func deleteStudent() {
	engine.DELETE("api/students/:studentid", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("studentid"))
		stats := http.StatusOK
		if err != nil {
			stats = http.StatusBadRequest
		}

		person := findStudent(id)
		if person == nil {
			stats = http.StatusNotFound
		}

		err = database.Remove(person)
		if err != nil {
			stats = http.StatusInternalServerError
		}

		c.JSON(stats, gin.H{
			"stats":    stats,
			"response": []string{},
		})
	})
}

func updateStudent() {
	engine.PATCH("api/students/:studentid", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		stats := http.StatusOK
		if err != nil {
			stats = http.StatusBadRequest
		}

		id, err := strconv.Atoi(c.Param("studentid"))
		if err != nil {
			stats = http.StatusBadRequest
		}

		var body map[string]interface{}
		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			stats = http.StatusBadRequest
		}

		person := findStudent(id)
		if person == nil {
			stats = http.StatusNotFound
		}

		person.Name = body["name"].(string)
		person.Email = body["email"].(string)
		person.Phone = body["phone"].(string)
		person.Age = int(body["age"].(float64))
		person.Grade = body["grade"].(string)

		err = database.Update(person)
		if err != nil {
			stats = http.StatusBadRequest
		}

		c.JSON(stats, gin.H{
			"stats":    stats,
			"response": person,
		})
	})
}

func addStudent() {
	engine.POST("api/students", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		stats := http.StatusCreated
		if err != nil {
			stats = http.StatusBadRequest
		}

		var body map[string]interface{}
		err = json.Unmarshal(bodyBytes, &body)
		if err != nil {
			stats = http.StatusBadRequest
		}

		person := person.Person{Name: body["name"].(string), Email: body["email"].(string), Phone: body["phone"].(string), Age: int(body["age"].(float64)), Grade: body["grade"].(string)}
		err = database.Create(&person)
		if err != nil {
			stats = http.StatusBadRequest
		}

		c.JSON(stats, gin.H{
			"stats":    stats,
			"response": person,
		})
	})
}

func findStudent(id int) *person.Person {
	person, err := database.Get(id)
	if err != nil {
		return nil
	}

	return person
}
