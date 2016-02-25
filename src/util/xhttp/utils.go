package xhttp

import (

	"net/http"
	gorillaContext "github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type Composer  struct {
	middleWares []func(http.Handler) http.Handler
}

func Adapt(handler http.Handler) httprouter.Handle  {
	return httprouter.Handle(
		func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
			ctx := &Context{
				Params: params,
			}

			gorillaContext.Set(r, keyContext, ctx)

			handler.ServeHTTP(w, r)
		})
}

func BindJSON(req *http.Request, obj interface{}) error {
	decoder := json.NewDecoder(req.Body)
	return decoder.Decode(obj)
}