package helpers

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

//JSONify in ajahhh wkwkw
func JSONify(ctx *fasthttp.RequestCtx, data map[string]interface{}) {
	//koplofy
	res, err := json.Marshal(data)

	if err != nil {
		log.Println("error while converting data to json")
		data["error"] = err
		panic(err)
	}

	ctx.SetContentType("application/json")
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.Write(res)
}

func koplofy() {

}
