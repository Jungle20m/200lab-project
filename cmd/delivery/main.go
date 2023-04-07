package main

import (
	_ "github.com/go-sql-driver/mysql"
)

//func main() {
//	db, err := sqlx.Open("mysql", "anhnv:anhnv!@#456@tcp(1.53.252.177:3307)/shopping")
//	if err != nil {
//		log.Fatalf("mysql error: %s", err)
//	}
//	mux := chi.NewMux()
//
//	appCtx := component.NewAppContext(db, mux)
//
//	restApi := rest.NewServer(appCtx)
//	restApi.ListenAndServe()
//}
