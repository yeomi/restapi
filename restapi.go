package main 

import (
	"github.com/julienschmidt/httprouter"
	"github.com/yeomi/restapi/models"
	"net/http"
	"encoding/json"
)

func main() {
	r := httprouter.New()



	r.GET("/user/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {


		user := &models.User{
			Name: "gabriel",
			Gender: "Male",
			Age: 27,
			Id: p.ByName("id"),
		}

		json, err := json.Marshal(user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(json)
	})

	http.ListenAndServe("localhost:8080", r)
}