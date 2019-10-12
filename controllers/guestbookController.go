package controllers

import (
	"fmt"
	"libraryapi/configs"
	helpers "libraryapi/helpers"
	models "libraryapi/models"
	repositories "libraryapi/repositories"

	"encoding/json"

	"github.com/valyala/fasthttp"
)

func GetGuestBook(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	guestbook, err := repositories.FindAllGuestbook(configs.DB)
	if err != nil {
		fmt.Println("error while get bookshelves :v\n %v", err)
	}

	data["guestbook"] = guestbook
	helpers.JSONify(ctx, data)
}

func CreateGuestBook(ctx *fasthttp.RequestCtx) {
	//mendapatkan value json post
	postValues := ctx.PostBody()

	//membuat variabel map
	var data = make(map[string]interface{})

	//convert json => struct
	guestbook := models.Guestbook{}
	if err := json.Unmarshal(postValues, &guestbook); err != nil {
		fmt.Println("ada error di controller->guestbookcontroller bagian unMarshal\n", err)
		data["message"] = "invalid json field"
		data["error"] = err.Error()
		helpers.JSONify(ctx, data)
	}

	//panggil Fungsi save di repo
	// _, err := repositories.InsertGuestbook(configs.DB, &guestbook)
	err := repositories.InsertGuestbook(configs.DB, &guestbook)

	if err != nil {
		data["message"] = "gagal insert data guestbook"
		data["error"] = err
	} else {
		data["message"] = "sukses menyimpan data guestbook"
	}

	helpers.JSONify(ctx, data)
}
