package auth

import (
	t "github.com/kmallari/cashr-Server/.gen/cashr/public/table"
	"github.com/kmallari/cashr-Server/.gen/cashr/public/model"
	"github.com/kmallari/cashr-Server/db"
	"log"
)

func createDefaultCategoriesForUser(userID string) {
	stmt := t.Category.INSERT(t.Category.UserID, t.Category.Icon, t.Category.Label, t.Category.IsExpense).
		VALUES(userID, "🍴", "Dining Out", true).
		VALUES(userID, "🍔", "Food Delivery", true).
		VALUES(userID, "🏡", "Rent", true).
		VALUES(userID, "🔌", "Electricity", true).
		VALUES(userID, "🌊", "Water", true).
		VALUES(userID, "📱", "Phone", true).
		VALUES(userID, "🌐", "Internet", true).
		VALUES(userID, "🚗", "Transportation", true).
		VALUES(userID, "🥓", "Groceries", true).
		VALUES(userID, "🍸", "Drinks", true).
		VALUES(userID, "👕", "Clothing", true).
		VALUES(userID, "🧺", "Laundry", true).
		VALUES(userID, "📺", "Subscriptions", true).
		VALUES(userID, "💊", "Medicine", true).
		VALUES(userID, "💸", "Salary", false).
		VALUES(userID, "💼", "Business", false).
		RETURNING(t.Category.AllColumns)

	var dest []struct {
		model.Category
	}

	err := stmt.Query(db.Client, &dest)
	if err != nil {
		log.Fatal(err.Error())
	}
}
