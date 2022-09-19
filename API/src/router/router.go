package router

import "github.com/gorilla/mux"

// Gerar vai retornar um router com as rotdas configuradas
func Gerar() *mux.Router {
	return mux.NewRouter()
}
