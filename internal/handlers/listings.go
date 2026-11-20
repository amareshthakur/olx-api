package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type listing struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	City        string    `json:"city"`
	CreatedAt   time.Time `json:"cretaion time"`
}

func List(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("all ok"))
		rows, err := db.Query(
			`SELECT * FROM listings
			ORDER BY created_at DESC
			LIMIT 100`)
		if err != nil {
			log.Printf("list all <-> List Handler Query : %v", err)
			http.Error(w, "internal error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		listings := []listing{}
		for rows.Next() {
			var l listing // 1->one ,l->L, i,I->eye
			if err := rows.Scan(&l.ID, &l.Title, &l.Description, &l.Price, &l.City, &l.CreatedAt); err != nil {
				log.Printf("rows.Scan : -> %v ", err)
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
			listings = append(listings, l)
		}
		if err := rows.Err(); err != nil {
			log.Printf("rows.Err: %v", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		// json encoder
		_ = json.NewEncoder(w).Encode(listings)
	}
}
