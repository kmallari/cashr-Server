package transactions

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kmallari/cashr-Server/api/utils"
	"github.com/kmallari/cashr-Server/db"
	sqlc "github.com/kmallari/cashr-Server/db/postgres"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"log"
	"net/http"
)

func Handler(r *mux.Router) {
	r.HandleFunc("", session.VerifySession(nil, createTransaction)).Methods(http.MethodPost)
	r.HandleFunc("", session.VerifySession(nil, getUserTransactions)).Methods(http.MethodGet)
	r.HandleFunc("/{id}", updateTransaction).Methods(http.MethodPut)
	r.HandleFunc("/{id}", session.VerifySession(nil, deleteTransaction)).Methods(http.MethodDelete)
}

func createTransaction(w http.ResponseWriter, r *http.Request) {
	userId := session.GetSessionFromRequestContext(r.Context()).GetUserID()

	var createTransactionReq CreateTransactionReq
	if err := utils.ParseAndValidateRequest(r.Body, &createTransactionReq); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	transaction, err := db.Queries.CreateTransaction(r.Context(), sqlc.CreateTransactionParams{
		UserID:      userId,
		Amount:      fmt.Sprintf("%v", createTransactionReq.Amount),
		Date:        createTransactionReq.Date,
		IsExpense:   *createTransactionReq.IsExpense,
		CategoryID:  createTransactionReq.CategoryID,
		Description: &createTransactionReq.Description,
	})

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, transaction)
}

func getUserTransactions(w http.ResponseWriter, r *http.Request) {
	userId := session.GetSessionFromRequestContext(r.Context()).GetUserID()

	transactions, err := db.Queries.GetUserTransactions(r.Context(), userId)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := []sqlc.GetUserTransactionsRow{}

	if len(transactions) > 0 {
		res = transactions
	}

	utils.SendJSONResponse(w, http.StatusOK, res)
}

func updateTransaction(w http.ResponseWriter, r *http.Request) {
	//postgres := r.Context().Value("postgres").(*sql.DB)
	//transactionId := mux.Vars(r)["id"]
	if r.Body != nil {
		log.Printf("%+v", r.Body)
	}

	var updateTransactionReq UpdateTransactionReq
	if err := utils.ParseAndValidateRequest(r.Body, &updateTransactionReq); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	log.Println("SHEESH")
	log.Printf("%+v", updateTransactionReq)

	//updateStatement = Transaction.UPDATE(
	//	Transaction.Date,
	//	Transaction.CategoryID,
	//	Transaction.IsExpense,
	//	Transaction.Description,
	//	Transaction.Amount,
	//	Transaction.Amount,
	//).
	//	SET(
	//		)

	//row, _ := postgres.Query("SELECT 1;")
	//log.Printf("%+v", row)
	return
}

func deleteTransaction(w http.ResponseWriter, r *http.Request) {
	userId := session.GetSessionFromRequestContext(r.Context()).GetUserID()
	transactionId := mux.Vars(r)["id"]

	if len(transactionId) != 36 {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.InvalidQueryParamError("invalid id"))
		return
	}

	if err := db.Queries.DeleteTransaction(r.Context(), sqlc.DeleteTransactionParams{
		ID:     transactionId,
		UserID: userId,
	}); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusAccepted, utils.SuccessfulDeleteMessage("transaction", transactionId))
}
