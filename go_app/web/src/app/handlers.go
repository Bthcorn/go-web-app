package app

import (
	"net/http"
)

type User struct {
	Id       int
	Username string
	Email    string
}

// var users = []User{
// 	{Id: 1, Username: "John", Email: "John@email.com"},
// 	{Id: 2, Username: "James", Email: "James@email.com"},
// 	{Id: 3, Username: "Jane", Email: "Jane@email.com"},
// 	{Id: 4, Username: "Jessica", Email: "Jessica@email.com"},
// }

func (a *App) Index(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "index.html", map[string]any{
		"title":  "Title from data",
		"header": "Header form data",
	})
}

func (a *App) About(w http.ResponseWriter, r *http.Request) {
	a.templates.Render(w, "about.html", nil)
}

// func (a *App) Users(w http.ResponseWriter, r *http.Request) {
// 	a.templates.Render(w, "users.html", map[string]any{
// 		"user": users,
// 	})
// }

func (a *App) Users(w http.ResponseWriter, r *http.Request) {
	res, err := a.db.Query("SELECT id, username, email FROM users")

	if err != nil {
		http.Error(w, err.Error(), 500) // http.StatusInternalServerError == 500
		return
	}

	var users []User
	// Iterate over the rows of sql query result
	// and append each user to the users slice
	for res.Next() {
		var user User
		res.Scan(&user.Id, &user.Username, &user.Email)
		users = append(users, user)
	}

	a.templates.Render(w, "users.html", map[string]any{
		"users": users,
	})
}

func (a *App) User(w http.ResponseWriter, r *http.Request) {
	// username := r.PathValue("username")

	// // found := false
	// index := -1

	// for i, user := range users {
	// 	if username == user.Username {
	// 		index = i
	// 		break
	// 	}
	// }

	// if index < 0 {
	// 	// w.WriteHeader(404)
	// 	// w.Write([]byte(`404 page not found`))
	// 	// fmt.Printf("error")
	// 	http.NotFound(w, r)
	// 	return
	// }

	// a.templates.Render(w, "user.html", map[string]any{
	// 	"user": users[index],
	// })
}
