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

func GetUsers(ctx *fasthttp.RequestCtx) {
	var data = make(map[string]interface{})

	users, err := repositories.FindAllUsers(configs.DB)
	if err != nil {
		fmt.Println("error while getusers :v\n %v", err)
	}

	data["users"] = users
	helpers.JSONify(ctx, data)
}

func CreateUsers(ctx *fasthttp.RequestCtx) {
	//mendapatkan value json post
	postValues := ctx.PostBody()

	//membuat variabel map
	var data = make(map[string]interface{})

	//convert json => struct
	user := models.Users{}
	if err := json.Unmarshal(postValues, &user); err != nil {
		fmt.Println("ada error di controller->usercontroller bagian unMarshal\n", err)
		data["message"] = "invalid json field"
		data["error"] = err.Error()
		helpers.JSONify(ctx, data)
	}

	//panggil Fungsi save di repo
	_, err := repositories.InsertUser(configs.DB, &user)

	if err != nil {
		data["message"] = "gagal insert data user"
		data["error"] = err
	} else {
		data["message"] = "sukses menyimpan data user"
	}

	helpers.JSONify(ctx, data)
}

func UpdateUser(ctx *fasthttp.RequestCtx) {
	//get data data user (json)
	postValues := ctx.PostBody()

	//baut variabel map
	var data = make(map[string]interface{})

	//convert json to struct
	user := models.Users{}
	if err := json.Unmarshal(postValues, &user); err != nil {
		data["message"] = "invalid json"
		data["error"] = err
		helpers.JSONify(ctx, data)
	}
	//panggil fungsi update  user dari repo user
	_, err := repositories.UpdateUser(configs.DB, &user)
	if err != nil {
		data["message"] = "gagal update data user"
		data["error"] = err
	} else {
		data["message"] = "sukses update data user"
	}

	helpers.JSONify(ctx, data)

}

func DeleteUser(ctx *fasthttp.RequestCtx) {
	//get id
	idvalue := fmt.Sprintf("%v", ctx.UserValue("id"))

	//convert string to int
	courseid, err := strconv.Atoi(idvalue)

	if err != nil {
		fmt.Println("Error convert id int->string")
	}

	//memanggil fungsi delete yang ada di repo
	_, err = repositories.DeleteUsers(configs.DB, courseid)

	//buat map
	var data = make(map[string]interface{})

	//cek err
	if err != nil {
		data["message"] = "gagal menghapus user"
		data["error"] = err
	} else {
		data["message"] = "sukses menghapus user dengan id = "
	}

	//output json
	helpers.JSONify(ctx, data)
}
