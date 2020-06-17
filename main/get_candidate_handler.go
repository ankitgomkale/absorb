package main

import (
	"github.com/ankitgomkale/absorb/models"
	"encoding/json"
	"fmt"
	"github.com/go-pg/pg"
	"net/http"
)

func GetCandidateHandler(w http.ResponseWriter, r *http.Request)  {
	//db, err := gorm.Open("postgres", "host=35.154.51.56 port=5432 user=absorb_go dbname=absorb password=AqYHWZr2Q7aDAMku")
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "absorb_go",
		Password: "AqYHWZr2Q7aDAMku",
		Database: "absorb",
	})

	if err := db.Ping(nil); err != nil {
	defer db.Close()
		panic(err)
	}

	var candidates []models.Candidate
	err := db.Model(&candidates).Select()
	if err != nil {
		panic(err)
	}

	candidateListBytes, err := json.Marshal(candidates)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(candidateListBytes)
	return
}
