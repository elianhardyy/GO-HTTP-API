package middlewares

import (
	"fmt"
	"net/http"
	"server/utils"
)

func ProtectedHandler(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization") // Bearer eyJsahuisahihnck
		if authHeader == ""{
			utils.ErrorResponse(w,http.StatusUnauthorized,"Missing Authorization Header")
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
		utils.JSONResponse(w,http.StatusContinue,claims.Email)
		next.ServeHTTP(w,r)
	})
}