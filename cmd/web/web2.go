package main

import (
	//...
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})

	// RESTy routes for "articles" resource
	r.Route("/articles", func(r chi.Router) {
		// Subrouters:
		r.Route("/{articleID}/{sxssxs}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Post("/", getArticle) // GET /articles/123
		})
	})

	http.ListenAndServe(":3001", r)
}

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleID")
		sxssxs := chi.URLParam(r, "sxssxs")
		ctx := context.WithValue(r.Context(), "article", articleID)
		ctx2 := context.WithValue(ctx, "sxssxs", sxssxs)

		next.ServeHTTP(w, r.WithContext(ctx2))
	})
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	articleID := chi.URLParam(r, "articleID")
	sxssxs := chi.URLParam(r, "sxssxs")

	fmt.Println(articleID)
	fmt.Println(sxssxs)

	//w.Write([]byte(fmt.Sprintf("title:%s%s", articleID,sxssxs)))
}
