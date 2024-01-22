package interceptor

//import (
//	"encoding/json"
//	"go_backend/utils/Token"
//	"log"
//	"net/http"
//)
//
//// TokenInterceptor is a middleware function to intercept and authenticate requests with token
//func TokenInterceptor(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		if r.Method == "OPTIONS" {
//			w.WriteHeader(http.StatusOK)
//			return
//		}
//		token := r.Header.Get("Authorization")
//		log.Println("token:", token)
//		if token != "" {
//			result := Token.VerifyToken(token)
//			if result {
//				log.Println("通过拦截器")
//				next.ServeHTTP(w, r)
//				return
//			}
//		}
//
//		w.Header().Set("Content-Type", "application/json; charset=utf-8")
//		w.WriteHeader(http.StatusUnauthorized)
//		response := map[string]interface{}{
//			"message":    "认证失败，未通过拦截器",
//			"resultCode": "401",
//		}
//		json.NewEncoder(w).Encode(response)
//		log.Println("认证失败，未通过拦截器")
//	})
//}
