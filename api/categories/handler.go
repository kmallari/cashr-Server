package categories

import (
	"github.com/gorilla/mux"
	"github.com/kmallari/cashr-Server/api/utils"
	db "github.com/kmallari/cashr-Server/db/postgres"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"log"
	"net/http"
	"strconv"
)

func Handler(r *mux.Router) {
	r.HandleFunc("", session.VerifySession(nil, createCategoryForUser)).Methods(http.MethodPost)
	r.HandleFunc("/user", session.VerifySession(nil, getUserCategories)).Methods(http.MethodGet)
}

func createCategoryForUser(w http.ResponseWriter, r *http.Request) {
	userId := session.GetSessionFromRequestContext(r.Context()).GetUserID()
	queries := r.Context().Value("queries").(*db.Queries)

	var createCategoryReq CreateCategoryReq
	if err := utils.ParseAndValidateRequest(r.Body, &createCategoryReq); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	category, err := queries.CreateCategory(r.Context(), db.CreateCategoryParams{
		UserID:    userId,
		IsExpense: *createCategoryReq.IsExpense,
		Label:     createCategoryReq.Label,
		Icon:      &createCategoryReq.Icon,
	})
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, category)
}

func getUserCategories(w http.ResponseWriter, r *http.Request) {
	userId := session.GetSessionFromRequestContext(r.Context()).GetUserID()
	queries := r.Context().Value("queries").(*db.Queries)

	q := r.URL.Query()
	isExpense := q.Get("is_expense")
	isExpenseBool, err := strconv.ParseBool(isExpense)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.InvalidQueryParamError("is_expense"))
		return
	}

	log.Println(userId, isExpenseBool)

	categories, err := queries.GetUserCategories(r.Context(), db.GetUserCategoriesParams{
		UserID:    userId,
		IsExpense: isExpenseBool,
	})
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := []db.Category{}

	if len(categories) > 0 {
		res = categories
	}

	utils.SendJSONResponse(w, http.StatusOK, res)
}
