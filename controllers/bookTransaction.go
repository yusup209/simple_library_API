package controllers

import (
	"fmt"
	"libraryapi/configs"
	helpers "libraryapi/helpers"
	models "libraryapi/models"
	repositories "libraryapi/repositories"
	"strconv"

	"encoding/json"

	"github.com/valyala/fasthttp"
)

func GetBookTransaction(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	book_transaction, err := repositories.FindAllBookTransaction(configs.DB)
	if err != nil {
		fmt.Println("error while get book transaction :v\n %v", err)
	}

	data["book_transaction"] = book_transaction
	helpers.JSONify(ctx, data)
}

func CreateBookTransaction(ctx *fasthttp.RequestCtx) {
	//mendapatkan value json post
	postValues := ctx.PostBody()

	//membuat variabel map
	var data = make(map[string]interface{})

	//convert json => struct
	book_transaction := models.BookTransaction{}
	if err := json.Unmarshal(postValues, &book_transaction); err != nil {
		fmt.Println("ada error di controller->booktransaction bagian unMarshal\n", err)
		data["message"] = "invalid json field"
		data["error"] = err.Error()
		helpers.JSONify(ctx, data)
	}

	//panggil Fungsi save di repo
	// _, err := repositories.Inser.BookTransaction(configs.DB, &book_transaction)
	_, err := repositories.InsertBookTransaction(configs.DB, &book_transaction)

	if err != nil {
		data["message"] = "gagal insert data book_transaction"
		data["error"] = err
	} else {
		data["message"] = "sukses menyimpan data book_transaction"
	}

	helpers.JSONify(ctx, data)
}

func ReturnBook(ctx *fasthttp.RequestCtx) {
	//get id
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	bookid, err := strconv.Atoi(idvalue)

	if err != nil {
		fmt.Println("Error convert id int->string")
	}

	//memanggil fungsi delete yang ada di repo
	_, err = repositories.ReturnBook(configs.DB, uint(bookid))

	//buat map
	var data = make(map[string]interface{})

	//cek err
	if err != nil {
		data["message"] = "gagal mengembalikan buku"
		data["error"] = err
	} else {
		data["message"] = "sukses mengembalikan buku dengan id = " + idvalue
	}

	//output json
	helpers.JSONify(ctx, data)
}
