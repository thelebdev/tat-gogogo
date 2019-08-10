package curriculum

import (
	"tat_gogogo/crawler/curriculum"

	"log"

	"github.com/gin-gonic/gin"
)

/*
CoursesController handle search courses
the default target student will be self
*/
func CoursesController(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")
	targetStudentID := c.PostForm("targetStudentID")
	year := c.PostForm("year")
	sem := c.PostForm("semester")

	result, err := curriculum.GetCourses(studentID, password, targetStudentID, year, sem)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	if result.GetStatus() != 200 {
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	c.JSON(result.GetStatus(), result.GetData())
}
