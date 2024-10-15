package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"
	"strconv"
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
	//newUserDto.Name = request.Form.Get("name")
	//newUserDto.Email = request.Form.Get("email")
	//newUserDto.Password = request.Form.Get("password")
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
	fmt.Println("will generate")
	//findId = c.UserService.FindById()
	tokenString, err := utils.GenerateToken(loginUser.Email)
	fmt.Println("generated")
	if err != nil{
		utils.ErrorResponse(response,http.StatusUnauthorized,"cannot generate")
		return
	}

	//utils.JSONResponse(response,http.StatusCreated,usr)
	token := &dto.TokenResponse{
		AccessToken: tokenString,
	}
	utils.JSONResponse(response,http.StatusAccepted,token)
	//hashedpassword := bcrypt.CompareHashAndPassword()
	//bcrypt.CompareHashAndPassword()
}
func (c *UserController) UserMe(response http.ResponseWriter, request *http.Request){
	//var udto dto.UserLoginDto
	//response.Header().Set("Content-Type","application/json")

	//request.
	//user, ok := request.Context().Value(&loginUser).(dto.UserLoginDto)
	//fmt.Println("gagal disini")
	email := response.Header().Get("email")
	// if !ok {
	// 	utils.ErrorResponse(response,http.StatusBadRequest,"error dashboard")
	// 	return
	// }
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

// func googleAuth(response http.ResponseWriter, request http.Request){
// 	session, _ = store
// }