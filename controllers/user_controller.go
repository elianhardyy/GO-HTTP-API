package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"
	"strconv"
	"time"
)

type UserController struct {
	UserService services.UserService
}
func NewUserController(s services.UserService) UserController{
	return UserController{
		UserService: s,
	}
}

func (c *UserController) Register(response http.ResponseWriter, request *http.Request){
	var newUserDto dto.UserDto
	request.ParseForm()
	if err := json.NewDecoder(request.Body).Decode(&newUserDto); err != nil{
		utils.ErrorResponse(response,http.StatusInternalServerError,"error")
		return
	}
	res, err := c.UserService.SaveOrUpdate(newUserDto)
	if err != nil{
		utils.ErrorResponse(response,http.StatusInternalServerError,"error")
		return
	}
	utils.JSONResponse(response,http.StatusCreated,res)
	
}
var loginUser dto.UserLoginDto
func (c *UserController) Login(response http.ResponseWriter, request *http.Request){
	request.ParseForm()
	if err := json.NewDecoder(request.Body).Decode(&loginUser);err != nil {
		utils.ErrorResponse(response,http.StatusInternalServerError,"error request")
	}
	_,err := c.UserService.FindByEmail(loginUser.Email,loginUser.Password)
	if err != nil {
		utils.ErrorResponse(response,http.StatusInternalServerError,"error credentials")
		return
	}
	tokenString, err := utils.GenerateToken(loginUser.Email)
	fmt.Println("generated")
	if err != nil{
		utils.ErrorResponse(response,http.StatusUnauthorized,"cannot generate")
		return
	}

	token := &dto.TokenResponse{
		AccessToken: tokenString,
	}
	http.SetCookie(response, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: time.Now().Add(15 * time.Minute),
		HttpOnly: true,
	})
	utils.JSONResponse(response,http.StatusAccepted,token)
}
func (c *UserController) UserMe(response http.ResponseWriter, request *http.Request){
	email := response.Header().Get("email")
	findEmail := c.UserService.EmailAuth(email)
	fmt.Println("success dashboard")
	utils.JSONResponse(response,http.StatusFound,findEmail)
}

func (c *UserController) SingleUser(response http.ResponseWriter, request *http.Request){
	id := request.URL.Query().Get("userid")
	num, err := strconv.ParseUint(id,10,64)
	if err != nil {
		fmt.Println("error")
	}
	user := c.UserService.FindById(uint(num))
	utils.JSONResponse(response, http.StatusFound,user)
}

func (c *UserController) Logout(response http.ResponseWriter, request *http.Request){
	http.SetCookie(response, &http.Cookie{
		Name: "token",
		Value: "",
		Expires: time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
	})
	utils.JSONResponse(response,http.StatusAccepted,"logout")
}