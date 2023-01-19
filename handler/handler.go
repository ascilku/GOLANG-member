package handler

import (
	"go-test-2/error"
	"go-test-2/member"
	"go-test-2/respons"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type handler struct {
	service member.Service
}

func NewHandler(service member.Service) *handler {
	return &handler{service}
}

func (h *handler) SaveHandler(g *gin.Context) {
	var key_input_member member.InputMember
	err := g.ShouldBindJSON(&key_input_member)
	if err != nil {
		errorMessage := gin.H{"errors": error.ErrorMessage(err)}
		responsApi := respons.ResponsApi("Failed Input Data", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, responsApi)
	} else {
		new_handler, err := h.service.SaveService(key_input_member)
		if err != nil {
			errorMessage := gin.H{"errors": err}
			responsApi := respons.ResponsApi("Failed Server", http.StatusBadRequest, "Failed", errorMessage)
			g.JSON(http.StatusBadRequest, responsApi)
			return
		} else {
			formatter_data := member.FormatterResponsData(new_handler, "token token")
			responsApi := respons.ResponsApi("Success Input Data", http.StatusOK, "SUCCESS", formatter_data)
			g.JSON(http.StatusOK, responsApi)
		}
	}
}

func (h *handler) LoginHandler(g *gin.Context) {
	var loginMember member.LoginMember
	err := g.ShouldBindJSON(&loginMember)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		errorMessage := gin.H{"errors": errors}
		responsApi := respons.ResponsApi("Failed Login", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, responsApi)
		return
	} else {
		loginService, err := h.service.LoginService(loginMember)
		if err != nil {
			errorMessage := gin.H{"error": err.Error()}
			responsApi := respons.ResponsApi("Failed Login", http.StatusUnprocessableEntity, "FAILED", errorMessage)
			g.JSON(http.StatusUnprocessableEntity, responsApi)
			return
		} else {
			formatter := member.FormatterResponsData(loginService, "token token")
			responsApi := respons.ResponsApi("Success Login", http.StatusOK, "SUCCESS", formatter)
			g.JSON(http.StatusOK, responsApi)
		}

	}
}

func (h *handler) CheckEmailAvailable(g *gin.Context) {
	var keyCheckEmailIsAvailable member.CheckEmailIsAvailable
	err := g.ShouldBindJSON(&keyCheckEmailIsAvailable)
	if err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}
		errorMessage := gin.H{"errors": errors}
		responsApi := respons.ResponsApi("Failed Check Email", http.StatusUnprocessableEntity, "FAILED", errorMessage)
		g.JSON(http.StatusUnprocessableEntity, responsApi)
	} else {
		checkEmailIsAvailable, err := h.service.CheckEmailIsAvailable(keyCheckEmailIsAvailable)
		if err != nil {
			errorMessage := gin.H{"error": "Error Server"}
			responsApi := respons.ResponsApi("Failed Check Email", http.StatusUnprocessableEntity, "FAILED", errorMessage)
			g.JSON(http.StatusBadRequest, responsApi)
			return
		} else {
			data := gin.H{
				"status": checkEmailIsAvailable,
			}
			dataMessage := "Email Available"
			if checkEmailIsAvailable {
				dataMessage = "Email Not Available"
			}
			responsApi := respons.ResponsApi(dataMessage, http.StatusUnprocessableEntity, "FAILED", data)
			g.JSON(http.StatusOK, responsApi)
		}
	}
}

func (h *handler) AvatarHandler() {}
