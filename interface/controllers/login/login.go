package login

import (
	"tat_gogogo/domain/repository"
	"tat_gogogo/domain/service"
	"tat_gogogo/usecase"

	"log"

	"github.com/gin-gonic/gin"
)

/*
Controller is a function for gin to handle login api
*/
func Controller(c *gin.Context) {
	studentID := c.PostForm("studentID")
	password := c.PostForm("password")

	repo := repository.NewResultRepository()
	service := service.NewResultService(repo)
	resultUsecase := usecase.NewResultUseCase(repo, service)
	result, err := resultUsecase.LoginResult(studentID, password)

	if err != nil {
		log.Panicln("failed to fetch login cookie")
		c.Status(500)
		return
	}

	if result.GetStatus() != 200 {
		c.JSON(result.GetStatus(), gin.H{
			"message": result.GetData(),
		})
		return
	}

	c.Status(200)

}
