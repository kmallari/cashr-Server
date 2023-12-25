package api

import (
	"github.com/gorilla/mux"
	"github.com/kmallari/cashr-Server/api/categories"
	"github.com/kmallari/cashr-Server/api/transactions"
)

func Handler(r *mux.Router) {
	transactions.Handler(r.PathPrefix("/transactions").Subrouter())
	categories.Handler(r.PathPrefix("/categories").Subrouter())
}
