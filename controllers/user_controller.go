package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/dto"
	"server/services"
	"server/utils"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/microcosm-cc/bluemonday"
)

type UserController struct {
	UserService services.UserService
}
func NewUserController(s services.UserService) UserController{
	return UserController{
		UserService: s,
	}
}
var url = "http://localhost:7000"
func (c *UserController) Register(response http.ResponseWriter, request *http.Request){
	var newUserDto dto.UserDto
	if err := json.NewDecoder(request.Body).Decode(&newUserDto); err != nil{
		utils.ErrorResponse(response,http.StatusInternalServerError,"error")
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	newUserDto.Name = sanitizer.Sanitize(newUserDto.Name)
	newUserDto.Email = sanitizer.Sanitize(newUserDto.Email)
	newUserDto.Password = sanitizer.Sanitize(newUserDto.Password)
	res, err := c.UserService.SaveOrUpdate(newUserDto)
	
	if err != nil{
		utils.ErrorResponse(response,http.StatusInternalServerError,"error")
		return
	}
	log.Println(url+"/register")
	utils.JSONResponse(response,http.StatusCreated,res)
	
}
var loginUser dto.UserLoginDto
func (c *UserController) Login(response http.ResponseWriter, request *http.Request){
	if err := json.NewDecoder(request.Body).Decode(&loginUser);err != nil {
		utils.ErrorResponse(response,http.StatusInternalServerError,"error request")
		return
	}
	sanitizer := bluemonday.UGCPolicy()
	loginUser.Email = sanitizer.Sanitize(loginUser.Email)
	loginUser.Password = sanitizer.Sanitize(loginUser.Password)
	user,err := c.UserService.FindByEmail(loginUser.Email,loginUser.Password)
	if err != nil {
		utils.ErrorResponse(response,http.StatusInternalServerError,"error credentials")
		return
	}
	tokenString, err := utils.GenerateToken(user)
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
	log.Println(url+"/login")
	utils.JSONResponse(response,http.StatusAccepted,token)
}
func (c *UserController) VerifyUser(w http.ResponseWriter, r *http.Request){
	var tokenDto dto.TokenDto
	if err := json.NewDecoder(r.Body).Decode(&tokenDto); err != nil {
		utils.ErrorResponse(w,http.StatusInternalServerError,"error request lur")
		return
	}
	err := c.UserService.VerifyTokenS(tokenDto.Token)
	if err != nil {
		utils.ErrorResponse(w,http.StatusInternalServerError,err.Error())
		return
	}
	log.Println(url+"/verify")
	utils.JSONResponse(w,http.StatusAccepted,"success")
}
func (c *UserController) UserMe(response http.ResponseWriter, request *http.Request){
	email := response.Header().Get("email")
	_,err := c.UserService.EmailAuth(email)
	if err != nil {
		utils.ErrorResponse(response,http.StatusNotFound,"error")
	}
	fmt.Println("success dashboard")
	utils.JSONResponse(response,http.StatusFound,"success")
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

func (c *UserController) UploadProfile(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("userid")
	num, err := strconv.ParseUint(id,10,64)
	if err != nil {
		fmt.Println("error")
	}
	const maxUploadSize = 10 * 1024 * 1024
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"File too large")
		return
	}
	name := r.FormValue("name")
	email := r.FormValue("email")
	if name == "" || email == ""{
		utils.ErrorResponse(w,http.StatusBadRequest,"Name and email required")
		return
	}
	file,header,err := r.FormFile("profile")
	if err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"Profile is required")
		return
	}
	defer file.Close()
	if header.Size > maxUploadSize {
		utils.ErrorResponse(w,http.StatusBadRequest,"File too large")
		return
	}
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"Config error")
		return
	}
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(),&s3.PutObjectInput{
		Bucket: aws.String("chatapp-bucket-storage"),
		Key: aws.String(header.Filename),
		Body: file,
		ACL: types.ObjectCannedACL(*aws.String("public-read")),
	})
	if err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"s3 config error")
		return
	}
	// var profile dto.ProfileDto
	
	profile := dto.ProfileDto{
		Name: name,
		Email: email,
		Profile: header.Filename,
	}
	users,err := c.UserService.UpdateProfile(profile,uint(num))
	if err != nil {
		utils.ErrorResponse(w,http.StatusBadRequest,"update error")
		return
	}
	log.Println(url+"/profile")
	fmt.Println(result)
	utils.JSONResponse(w, http.StatusCreated,users)
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