package middlewares

import (
	"fmt"
	"net/http"
	"server/dto"
	"server/utils"
)

func ProtectedHandler(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization") // Bearer eyJsahuisahihnck
		if authHeader == ""{
			utils.ErrorResponse(w,http.StatusUnauthorized,"Missing Authorization Header")
			return
		}
		_, err := r.Cookie("token")
		if err != nil{
			if err == http.ErrNoCookie{
				utils.ErrorResponse(w,http.StatusUnauthorized,"unauthorized about cookie")
				return
			}
			utils.ErrorResponse(w,http.StatusBadRequest,"bad")
			return
		}
		tokenString := authHeader[len("Bearer "):]
		//misal := authHeader[len("Bearer "):]
		fmt.Println(tokenString);
		claims, err := utils.VerifyToken(tokenString)
		if err != nil{
			utils.ErrorResponse(w,http.StatusUnauthorized,"Unauthorized protected")
			return 
		}
		// if claims.Role != roleUser{
		// 	return
		// }
		r.Header.Set("email",claims.Email)
		utils.JSONResponse(w,http.StatusContinue,dto.UserResponse{
			Name: claims.Name,
			Email: claims.Email,
		})
		next.ServeHTTP(w,r)
	})
}