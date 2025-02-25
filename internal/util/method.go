package util

import "net/http"

func Post(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
		next.ServeHTTP(res, req)
	})
}

func Get(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
		next.ServeHTTP(res, req)
	})
}

func Put(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
		next.ServeHTTP(res, req)
	})
}

func Delete(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
		next.ServeHTTP(res, req)
	})
}

func Patch(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPatch {
			http.Error(res, "Method not allowed", http.StatusMethodNotAllowed)
		}
		next.ServeHTTP(res, req)
	})
}
