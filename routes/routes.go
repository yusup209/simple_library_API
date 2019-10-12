package routes

import (
	"fmt"
	ctrl "libraryapi/controllers"
	"log"

	cors "github.com/AdhityaRamadhanus/fasthttpcors"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func handleIndex(ctx *fasthttp.RequestCtx) {
	fmt.Fprintln(ctx, "tes tes123")
}

func Setup() {
	fmt.Println("routes setup")

	//untuk routing (hmm.. mirip express ya yg ada di nodejs)
	r := fasthttprouter.New()
	r.GET("/", handleIndex)
	//buat users

	r.GET("/api/users", ctrl.GetUsers)
	r.POST("/api/users/create", ctrl.CreateUsers)
	r.PUT("/api/users/update", ctrl.UpdateUser)
	r.DELETE("/api/users/delete/:id", ctrl.DeleteUser)

	//category
	r.GET("/api/category", ctrl.GetCategory)
	r.GET("/api/category/:id", ctrl.GetCategoryByID)
	r.POST("/api/category/create", ctrl.CreateCategory)
	r.PUT("/api/category/update", ctrl.UpdateCategory)
	r.DELETE("/api/category/delete/:id", ctrl.DeleteCategory)

	//bookshelf
	r.GET("/api/bookshelf", ctrl.GetBookshelves)
	r.GET("/api/bookshelf/:id", ctrl.GetBookshelfByID)
	r.POST("/api/bookshelf/create", ctrl.CreateBookshelf)
	r.PUT("/api/bookshelf/update", ctrl.UpdateBookshelf)
	r.DELETE("/api/bookshelf/delete/:id", ctrl.DeleteBookshelf)

	//book
	r.GET("/api/book", ctrl.GetBooks)
	r.GET("/api/book/:id", ctrl.GetBookByID)
	r.POST("/api/book/create", ctrl.CreateBook)
	r.PUT("/api/book/update", ctrl.UpdateBook)
	r.DELETE("/api/book/delete/:id", ctrl.DeleteBook)

	//book transaction
	r.GET("/api/transaction", ctrl.GetBookTransaction)
	r.POST("/api/transaction/create", ctrl.CreateBookTransaction)
	r.PUT("/api/transaction/return/:id", ctrl.ReturnBook)

	//book guest_book
	r.GET("/api/guest_book", ctrl.GetGuestBook)
	r.POST("/api/guest_book/create", ctrl.CreateGuestBook)

	//portnya
	listenAddr := ":8000"
	fmt.Println("Run fasthttttttp di port", listenAddr)

	withCORS := cors.NewCorsHandler(cors.Options{
		// if you leave allowedOrigins empty then fasthttpcors will treat it as "*"
		AllowedOrigins: []string{"*"}, // Only allow example.com to access the resource
		// if you leave allowedHeaders empty then fasthttpcors will accept any non-simple headers
		AllowedHeaders: []string{"x-something-client", "Content-Type", "content-type"}, // only allow x-something-client and Content-Type in actual request

		// if you leave this empty, only simple method will be accepted
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"}, // only allow get or post to resource
		AllowCredentials: false,                                    // resource doesn't support credentials
		AllowMaxAge:      5600,                                     // cache the preflight result
		Debug:            true,
	})
	if err := fasthttp.ListenAndServe(listenAddr, withCORS.CorsMiddleware(r.Handler)); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
