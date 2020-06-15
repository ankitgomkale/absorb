package main

import (
	"absorb/models"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetCandidateHandler(w http.ResponseWriter, r *http.Request)  {
	db, err := gorm.Open("postgres", "host=172.31.0.146 port=5432 user=absorb_go dbname=absorb password=AqYHWZr2Q7aDAMku")
	defer db.Close()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		return
	}

	candidateListBytes, err := json.Marshal(models.Candidates)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(candidateListBytes)
}
