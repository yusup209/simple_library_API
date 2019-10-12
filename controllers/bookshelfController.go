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

func GetBookshelves(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	bookshelf, err := repositories.FindAllBookshelf(configs.DB)
	if err != nil {
		fmt.Println("error while get bookshelves :v\n %v", err)
	}

	data["bookshelf"] = bookshelf
	helpers.JSONify(ctx, data)
}

func GetBookshelfByID(ctx *fasthttp.RequestCtx) {
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	bookshelfid, err := strconv.Atoi(idvalue)
	var data = make(map[string]interface{})

	bookshelf, err := repositories.FindBookshelfByID(configs.DB, uint(bookshelfid))
	if err != nil {
		fmt.Println("error while get bookshelves :v\n %v", err)
	}

	data["bookshelf"] = bookshelf
	helpers.JSONify(ctx, data)
}

func CreateBookshelf(ctx *fasthttp.RequestCtx) {
	//mendapatkan value json post
	postValues := ctx.PostBody()

	//membuat variabel map
	var data = make(map[string]interface{})

	//convert json => struct
	bookshelf := models.Bookshelf{}
	if err := json.Unmarshal(postValues, &bookshelf); err != nil {
		fmt.Println("ada error di controller->bookshelfcontroller bagian unMarshal\n", err)
		data["message"] = "invalid json field"
		data["error"] = err.Error()
		helpers.JSONify(ctx, data)
	}

	//panggil Fungsi save di repo
	_, err := repositories.InsertBookshelf(configs.DB, bookshelf)

	if err != nil {
		data["message"] = "gagal insert data bookshelf"
		data["error"] = err
	} else {
		data["message"] = "sukses menyimpan data bookshelf"
	}

	helpers.JSONify(ctx, data)
}

func UpdateBookshelf(ctx *fasthttp.RequestCtx) {
	//get data data user (json)
	postValues := ctx.PostBody()

	//baut variabel map
	var data = make(map[string]interface{})

	//convert json to struct
	bookshelf := models.Bookshelf{}
	if err := json.Unmarshal(postValues, &bookshelf); err != nil {
		data["message"] = "invalid json"
		data["error"] = err
		helpers.JSONify(ctx, data)
	}
	//panggil fungsi update  bookshelf dari repo Bookshelf
	_, err := repositories.UpdateBookshelf(configs.DB, &bookshelf)
	if err != nil {
		data["message"] = "gagal update data Bookshelf"
		data["error"] = err
	} else {
		data["message"] = "sukses update data Bookshelf"
	}

	helpers.JSONify(ctx, data)

}

func DeleteBookshelf(ctx *fasthttp.RequestCtx) {
	//get id
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	bookshelfid, err := strconv.Atoi(idvalue)

	if err != nil {
		fmt.Println("Error convert id int->string")
	}

	//memanggil fungsi delete yang ada di repo
	_, err = repositories.DeleteBookshelf(configs.DB, bookshelfid)

	//buat map
	var data = make(map[string]interface{})

	//cek err
	if err != nil {
		data["message"] = "gagal menghapus bookshelf"
		data["error"] = err
	} else {
		data["message"] = "sukses menghapus bookshelf dengan id = "
	}

	//output json
	helpers.JSONify(ctx, data)
}
