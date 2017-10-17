package main

import (
	"./routes/callback"
	"./routes/home"
	"./routes/middlewares"
	"./routes/user"
	"./routes/login"
	"./routes/logout"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)

func StartServer() {
	r := mux.NewRouter()

	r.HandleFunc("/", home.HomeHandler)
	r.HandleFunc("/login", login.LoginHandler)
	r.HandleFunc("/logout", logout.LogoutHandler)
	r.HandleFunc("/callback", callback.CallbackHandler)
	r.Handle("/user", negroni.New(
		negroni.HandlerFunc(middlewares.IsAuthenticated),
		negroni.Wrap(http.HandlerFunc(user.UserHandler)),
	))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	http.Handle("/", r)
	http.ListenAndServe("0.0.0.0:3000", nil)
}
