package api

import (
	"cf-proposal/common/logservice"
	"cf-proposal/common/types"
	"cf-proposal/domain/model"
	"cf-proposal/domain/repository"
	"cf-proposal/domain/services/urlservice"
	"cf-proposal/infrastructure/sqlite3helper"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type UrlController struct{}

const createpath types.Path = "/create"

func (uc UrlController) Routes() chi.Router {
	router := chi.NewRouter()

	router.Post(string(createpath), uc.HandleCreate)

	return router
}

func (uc UrlController) HandleCreate(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var url model.UrlDto
	err = json.Unmarshal(body, &url)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
		return
	}

	//conn, err := sql.Open("sqlite3", "./shortener.db")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer conn.Close()
	urlrepo := repository.InitUrlRepo(sqlite3helper.DbConn)
	urlservice := urlservice.Init(urlrepo)
	createdUrl, err := urlservice.Create(context.Background(), url)
	if err != nil {
		logservice.LogError("400", "GET", createpath, err)
		return
	}
	//conn.Close()

	w.Header().Set("Content-Type", "application/json")
	jsonResp, _ := json.Marshal(createdUrl)
	w.Write(jsonResp)
	logservice.LogHttpRequest("200", "POST", createpath)
}