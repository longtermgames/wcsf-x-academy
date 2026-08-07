package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type registration struct {
	ID         int64     `json:"id"`
	FullName   string    `json:"full_name"`
	Phone      string    `json:"phone"`
	Discipline string    `json:"discipline"`
	CreatedAt  time.Time `json:"created_at"`
}

var allowedDisciplines = map[string]string{
	"bmx":        "BMX",
	"workout":    "Воркаут-баттл",
	"trampoline": "Батут",
}

type registerRequest struct {
	FullName   string `json:"full_name"`
	Phone      string `json:"phone"`
	Discipline string `json:"discipline"`
	// honeypot field: real users never fill it, bots that autofill every input do
	Website string `json:"website"`
}

func (s *server) withCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", s.allowedOrigin)
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func (s *server) handleOptions(w http.ResponseWriter, r *http.Request) {
	s.withCORS(w)
	w.WriteHeader(http.StatusNoContent)
}

func (s *server) handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func (s *server) handleRegister(w http.ResponseWriter, r *http.Request) {
	s.withCORS(w)

	var req registerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	if req.Website != "" {
		w.WriteHeader(http.StatusCreated)
		return
	}

	req.FullName = strings.TrimSpace(req.FullName)
	req.Phone = strings.TrimSpace(req.Phone)

	if req.FullName == "" || req.Phone == "" {
		http.Error(w, "full_name and phone are required", http.StatusBadRequest)
		return
	}
	if _, ok := allowedDisciplines[req.Discipline]; !ok {
		http.Error(w, "invalid discipline", http.StatusBadRequest)
		return
	}

	var reg registration
	err := s.db.QueryRowContext(r.Context(),
		`INSERT INTO registrations (full_name, phone, discipline) VALUES ($1, $2, $3)
		 RETURNING id, full_name, phone, discipline, created_at`,
		req.FullName, req.Phone, req.Discipline,
	).Scan(&reg.ID, &reg.FullName, &reg.Phone, &reg.Discipline, &reg.CreatedAt)
	if err != nil {
		log.Printf("insert registration: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(reg)
}

func (s *server) requireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if s.adminUser == "" || !ok || user != s.adminUser || pass != s.adminPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="admin"`)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func (s *server) fetchRegistrations(r *http.Request) ([]registration, error) {
	rows, err := s.db.QueryContext(r.Context(),
		`SELECT id, full_name, phone, discipline, created_at FROM registrations ORDER BY created_at DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []registration
	for rows.Next() {
		var reg registration
		if err := rows.Scan(&reg.ID, &reg.FullName, &reg.Phone, &reg.Discipline, &reg.CreatedAt); err != nil {
			return nil, err
		}
		out = append(out, reg)
	}
	return out, rows.Err()
}

func (s *server) handleAdmin(w http.ResponseWriter, r *http.Request) {
	regs, err := s.fetchRegistrations(r)
	if err != nil {
		log.Printf("fetch registrations: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	counts := map[string]int{}
	for _, reg := range regs {
		counts[reg.Discipline]++
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := adminTmpl.Execute(w, adminPageData{
		Registrations: regs,
		Counts:        counts,
		Names:         allowedDisciplines,
	}); err != nil {
		log.Printf("render admin: %v", err)
	}
}

func (s *server) handleExportCSV(w http.ResponseWriter, r *http.Request) {
	regs, err := s.fetchRegistrations(r)
	if err != nil {
		log.Printf("fetch registrations: %v", err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/csv; charset=utf-8")
	w.Header().Set("Content-Disposition", `attachment; filename="registrations.csv"`)

	cw := csv.NewWriter(w)
	cw.Write([]string{"ID", "Имя", "Телефон", "Дисциплина", "Дата"})
	for _, reg := range regs {
		cw.Write([]string{
			strconv.FormatInt(reg.ID, 10),
			reg.FullName,
			reg.Phone,
			allowedDisciplines[reg.Discipline],
			reg.CreatedAt.Format("2006-01-02 15:04"),
		})
	}
	cw.Flush()
}
