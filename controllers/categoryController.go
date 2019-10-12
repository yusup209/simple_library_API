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

func GetCategory(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	category, err := repositories.FindAllCategory(configs.DB)
	if err != nil {
		fmt.Println("error while get category :v\n %v", err)
	}

	data["category"] = category
	helpers.JSONify(ctx, data)
}

func GetCategoryByID(ctx *fasthttp.RequestCtx) {
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	category_id, err := strconv.Atoi(idvalue)
	var data = make(map[string]interface{})

	category, err := repositories.FindCategoryByID(configs.DB, uint(category_id))
	if err != nil {
		fmt.Println("error while get category :v\n %v", err)
	}

	data["category"] = category
	helpers.JSONify(ctx, data)
}

func CreateCategory(ctx *fasthttp.RequestCtx) {
	//mendapatkan value json post
	postValues := ctx.PostBody()

	fmt.Printf("data post : %v", postValues)

	//membuat variabel map
	var data = make(map[string]interface{})

	//convert json => struct
	category := models.Category{}
	if err := json.Unmarshal(postValues, &category); err != nil {
		fmt.Println("ada error di controller->category controller (create) bagian unMarshal\n", err)
		data["message"] = "invalid json field"
		data["error"] = err
		helpers.JSONify(ctx, data)
	}

	//panggil Fungsi save di repo
	_, err := repositories.InsertCategory(configs.DB, &category)

	if err != nil {
		data["message"] = "gagal insert data category"
		data["error"] = err
		helpers.JSONify(ctx, data)
	} else {
		data["message"] = "sukses menyimpan data category"
		helpers.JSONify(ctx, data)
	}

	helpers.JSONify(ctx, data)
}

func UpdateCategory(ctx *fasthttp.RequestCtx) {
	//get data data user (json)
	postValues := ctx.PostBody()

	//baut variabel map
	var data = make(map[string]interface{})

	//convert json to struct
	category := models.Category{}
	if err := json.Unmarshal(postValues, &category); err != nil {
		data["message"] = "invalid json"
		data["error"] = err
		helpers.JSONify(ctx, data)
	}
	//panggil fungsi update  category dari repo category
	_, err := repositories.UpdateCategory(configs.DB, &category)
	if err != nil {
		data["message"] = "gagal update data category"
		data["error"] = err
	} else {
		data["message"] = "sukses update data category"
	}

	helpers.JSONify(ctx, data)

}

func DeleteCategory(ctx *fasthttp.RequestCtx) {
	//get id
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	category_id, err := strconv.Atoi(idvalue)

	if err != nil {
		fmt.Println("Error convert id int->string on category")
	}

	//memanggil fungsi delete yang ada di repo
	_, err = repositories.DeleteCategory(configs.DB, category_id)

	//buat map
	var data = make(map[string]interface{})

	//cek err
	if err != nil {
		data["message"] = "gagal menghapus kategori"
		data["error"] = err
	} else {
		data["message"] = "sukses menghapus kategori dengan id = "
	}

	//output json
	helpers.JSONify(ctx, data)
}
