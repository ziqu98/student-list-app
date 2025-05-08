package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// Student struct to hold student data
type student struct {
	NRP   string  `json:"nrp"`
	Name  string  `json:"name"`
	Major string  `json:"major"`
	GPA   float64 `json:"gpa"`
}

// In-memory database
var students = []student{
	{NRP: "152021333", Name: "Cloud", Major: "Informatics", GPA: 3.77},
	{NRP: "112022111", Name: "Computing", Major: "Informatics", GPA: 3.56},
}

func main() {
	router := gin.Default()
	router.GET("/", getHome)
	router.GET("/students", getStudents)
	router.GET("/students/:nrp", getStudentByID)
	router.POST("/students", postStudents)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run("0.0.0.0:" + port)
}

// Handlers
func getHome(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"message": "welcome to student lists"})
}

func getStudents(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, students)
}

func postStudents(c *gin.Context) {
	var newStudent student

	if err := c.BindJSON(&newStudent); err != nil {
		return
	}

	students = append(students, newStudent)
	c.IndentedJSON(http.StatusCreated, newStudent)
}

func getStudentByID(c *gin.Context) {
	nrp := c.Param("nrp")

	for _, a := range students {
		if a.NRP == nrp {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "student not found"})
}
