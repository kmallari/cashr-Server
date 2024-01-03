package categories

import (
	"fmt"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/gorilla/mux"
	"github.com/kmallari/cashr-Server/.gen/cashr/public/model"
	. "github.com/kmallari/cashr-Server/.gen/cashr/public/table"
	"github.com/kmallari/cashr-Server/api/utils"
	"github.com/kmallari/cashr-Server/db"
	sqlc "github.com/kmallari/cashr-Server/db/postgres"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"log"
	"net/http"
	"strconv"
)

func Handler(r *mux.Router) {
	r.HandleFunc("", session.VerifySession(nil, createCategory)).Methods(http.MethodPost)
	//r.HandleFunc("", session.VerifySession(nil, updateCategories)).Methods(http.MethodPut)
	r.HandleFunc("", session.VerifySession(nil, getUserCategories)).Methods(http.MethodGet)
	r.HandleFunc("", session.VerifySession(nil, updateCategories)).Methods(http.MethodPut)
	r.HandleFunc("/{id}", session.VerifySession(nil, deleteCategory)).Methods(http.MethodDelete)
}

func createCategory(w http.ResponseWriter, r *http.Request) {
	userID := session.GetSessionFromRequestContext(r.Context()).GetUserID()
	var createCategoryReq CreateCategoryReq
	if err := utils.ParseAndValidateRequest(r.Body, &createCategoryReq); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	category, err := db.Queries.CreateCategory(r.Context(), sqlc.CreateCategoryParams{
		UserID:    userID,
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
	userID := session.GetSessionFromRequestContext(r.Context()).GetUserID()

	q := r.URL.Query()
	isExpense := q.Get("is_expense")

	var stmt SelectStatement

	if isExpense == "" {
		stmt =
			SELECT(Category.AllColumns).
				FROM(Category).
				WHERE(
					Category.UserID.EQ(String(userID)),
				).ORDER_BY(Category.IsExpense.DESC(), Category.Label)
	} else {
		isExpenseBool, err := strconv.ParseBool(isExpense)
		if err != nil {
			utils.SendJSONResponse(w, http.StatusBadRequest, utils.InvalidQueryParamError("is_expense"))
			return
		}
		stmt =
			SELECT(Category.AllColumns).
				FROM(Category).
				WHERE(
					Category.UserID.EQ(String(userID)).
						AND(Category.IsExpense.EQ(Bool(isExpenseBool))),
				).ORDER_BY(Category.IsExpense.DESC(), Category.Label)
	}

	query, _ := stmt.Sql()
	log.Println("QUERY:", query)

	var dest []struct {
		model.Category
	}

	err := stmt.Query(db.Client, &dest)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, dest)
	return
}

func updateCategories(w http.ResponseWriter, r *http.Request) {
	userID := session.GetSessionFromRequestContext(r.Context()).GetUserID()

	var updateCategoryReq []UpdateCategoryReq
	if err := utils.ParseAndValidateRequest(r.Body, &updateCategoryReq); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	rawQuery := `update category as c1 set icon = coalesce(c2.icon, c1.icon), label = coalesce(c2.label, c1.label), is_expense = coalesce(c2.is_expense::boolean, c1.is_expense) from (values `

	queryArgs := make(map[string]interface{})

	for i := 0; i < len(updateCategoryReq); i++ {
		var newIcon interface{}
		if updateCategoryReq[i].Icon != nil {
			newIcon = *updateCategoryReq[i].Icon
		} else {
			newIcon = nil
		}

		var newLabel interface{}
		if updateCategoryReq[i].Label != nil {
			newLabel = *updateCategoryReq[i].Label
		} else {
			newLabel = nil
		}

		var newIsExpense interface{}
		if updateCategoryReq[i].IsExpense != nil {
			newIsExpense = *updateCategoryReq[i].IsExpense
		} else {
			newIsExpense = nil
		}

		idKey := utils.ConvertToKey(i, "id")
		iconKey := utils.ConvertToKey(i, "icon")
		labelKey := utils.ConvertToKey(i, "label")
		isExpenseKey := utils.ConvertToKey(i, "isExpense")

		rawQuery += fmt.Sprintf(
			`(%v, %v, %v, %v)`,
			idKey,
			iconKey,
			labelKey,
			isExpenseKey,
		)
		if i < len(updateCategoryReq)-1 {
			rawQuery += ", "
		}

		queryArgs[fmt.Sprintf("%v", idKey)] = updateCategoryReq[i].ID
		queryArgs[fmt.Sprintf("%v", iconKey)] = newIcon
		queryArgs[fmt.Sprintf("%v", labelKey)] = newLabel
		queryArgs[fmt.Sprintf("%v", isExpenseKey)] = newIsExpense
	}

	rawQuery += `) as c2(id, icon, label, is_expense) where c2.id = c1.id and c1.user_id = $userID`
	queryArgs["$userID"] = userID

	_, err := RawStatement(rawQuery, queryArgs).Exec(db.Client)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusNoContent, "Successfully updated categories.")
}

func deleteCategory(w http.ResponseWriter, r *http.Request) {
	userID := session.GetSessionFromRequestContext(r.Context()).GetUserID()
	categoryId := mux.Vars(r)["id"]

	if len(categoryId) != 36 {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.InvalidQueryParamError("invalid id"))
		return
	}

	if err := db.Queries.DeleteCategory(r.Context(), sqlc.DeleteCategoryParams{
		ID:     categoryId,
		UserID: userID,
	}); err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJSONResponse(w, http.StatusAccepted, utils.SuccessfulDeleteMessage("transaction", categoryId))
	return
}
