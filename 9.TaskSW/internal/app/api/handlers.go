package api

import (
	"encoding/json"
	"github.com/Lexa-san/spc-go2/9.TaskSW/internal/app/models"
	"net/http"
)

//Common Message struct to API answers.
type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

//Inin common API headers.
func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

//Save parameters of uadratic.
func (api *API) GrabQuadratic(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Post Quadratic POST /grab")
	var q models.Quadratic

	err := json.NewDecoder(req.Body).Decode(&q)
	if err != nil {
		api.logger.Info("Invalid json recieved from client")
		msg := Message{
			StatusCode: 400,
			Message:    "Provided json is invalid",
			IsError:    true,
		}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	DB = &q
	DB.Nroots = 0
	api.logger.Info("Store Quadratic", *DB)
	writer.WriteHeader(201)
	//json.NewEncoder(writer).Encode(DB)

}

//// Return Quadratic model
//func (api *API) GetQuadratic(writer http.ResponseWriter, req *http.Request) {
//	initHeaders(writer)
//	api.logger.Info("Get Quadratic")
//
//	writer.WriteHeader(200)
//	json.NewEncoder(writer).Encode(DB)
//}

//Return Quadratic model with quantity of roots.
func (api *API) SolveQuadratic(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Quadratic")

	res, _ := DB.Solve()
	DB.Nroots = res

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(DB)
}
