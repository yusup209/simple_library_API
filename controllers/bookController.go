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

func GetBooks(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	book, err := repositories.FindAllBook(configs.DB)
	if err != nil {
		fmt.Println("error while get all books :v\n %v", err)
	}

	data["book"] = book
	helpers.JSONify(ctx, data)
}

func GetBookByID(ctx *fasthttp.RequestCtx) {
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	bookid, err := strconv.Atoi(idvalue)
	var data = make(map[string]interface{})

	book, err := repositories.FindBookByID(configs.DB, uint(bookid))
	if err != nil {
		fmt.Println("error while get books :v\n %v", err)
	}

	data["book"] = book
	helpers.JSONify(ctx, data)
}

func CreateBook(ctx *fasthttp.RequestCtx) {
	//mendapatkan value json post
	postValues := ctx.PostBody()

	//membuat variabel map
	var data = make(map[string]interface{})

	//convert json => struct
	book := models.Book{}
	if err := json.Unmarshal(postValues, &book); err != nil {
		fmt.Println("ada error di controller->bookcontroller bagian (create) unMarshal\n", err)
		data["message"] = "invalid json field"
		data["error"] = err.Error()
		helpers.JSONify(ctx, data)
	}

	//panggil Fungsi save di repo
	_, err := repositories.InsertBook(configs.DB, &book)

	if err != nil {
		data["message"] = "gagal insert data user"
		data["error"] = err
	} else {
		data["message"] = "sukses menyimpan data user"
	}

	helpers.JSONify(ctx, data)
}

func UpdateBook(ctx *fasthttp.RequestCtx) {
	//get data data user (json)
	postValues := ctx.PostBody()

	//baut variabel map
	var data = make(map[string]interface{})

	//convert json to struct
	book := models.Book{}
	if err := json.Unmarshal(postValues, &book); err != nil {
		data["message"] = "invalid json"
		data["error"] = err
		helpers.JSONify(ctx, data)
	}
	//panggil fungsi update  book dari repo book
	_, err := repositories.UpdateBook(configs.DB, &book)
	if err != nil {
		data["message"] = "gagal update data book"
		data["error"] = err
	} else {
		data["message"] = "sukses update data book"
	}

	helpers.JSONify(ctx, data)

}

func DeleteBook(ctx *fasthttp.RequestCtx) {
	//get id
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	bookid, err := strconv.Atoi(idvalue)

	if err != nil {
		fmt.Println("Error convert id int->string")
	}

	//memanggil fungsi delete yang ada di repo
	_, err = repositories.DeleteBook(configs.DB, bookid)

	//buat map
	var data = make(map[string]interface{})

	//cek err
	if err != nil {
		data["message"] = "gagal menghapus book"
		data["error"] = err
	} else {
		data["message"] = "sukses menghapus book dengan id = "
	}

	//output json
	helpers.JSONify(ctx, data)
}
