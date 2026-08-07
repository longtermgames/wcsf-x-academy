package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("open db: %v", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("ping db: %v", err)
	}
	if err := migrate(ctx, db); err != nil {
		log.Fatalf("migrate: %v", err)
	}

	srv := &server{
		db:            db,
		adminUser:     os.Getenv("ADMIN_USER"),
		adminPass:     os.Getenv("ADMIN_PASS"),
		allowedOrigin: envOr("ALLOWED_ORIGIN", "*"),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", srv.handleHealth)
	mux.HandleFunc("GET /api/health", srv.handleHealth)
	mux.HandleFunc("POST /api/register", srv.handleRegister)
	mux.HandleFunc("OPTIONS /api/register", srv.handleOptions)
	mux.HandleFunc("GET /admin", srv.requireAdmin(srv.handleAdmin))
	mux.HandleFunc("GET /admin/export.csv", srv.requireAdmin(srv.handleExportCSV))

	port := envOr("PORT", "8080")
	log.Printf("listening on :%s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
