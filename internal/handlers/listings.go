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

// constructor pattern
type ListingHandler struct {
	db *sql.DB
}

// constructor
func NewListingHandler(db *sql.DB) *ListingHandler {
	return &ListingHandler{
		db: db,
	}
}

// func GetList(db *sql.DB) http.HandlerFunc {

// (lh ListingHandler) -> Method Receiver
func (lh ListingHandler) List(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("all ok"))
	rows, err := lh.db.Query(
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

// }

// func DeleteList(db *sql.DB) http.HandlerFunc {
func (lh ListingHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	// fmt.Println(id)
	// execute query without returning any rows
	// result, err := db.Exec(
	_, err := lh.db.Exec(
		`DELETE FROM listings WHERE id = $1`, id)
	if err != nil {
		log.Printf("delete : %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	// rows, err := result.RowsAffected()
	w.WriteHeader(http.StatusNoContent)

	w.Write([]byte("ok"))
}

// }
