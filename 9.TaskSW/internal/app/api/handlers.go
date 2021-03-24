package api

import (
	"encoding/json"
	"github.com/Lexa-san/spc-go2/9.TaskSW/internal/app/models"
	"math"
	"net/http"
)

type Message struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	IsError    bool   `json:"is_error"`
}

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

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

	DB["A"] = q.A
	DB["B"] = q.B
	DB["C"] = q.C
	api.logger.Info("Store Quadratic", DB)
	writer.WriteHeader(201)
	//json.NewEncoder(writer).Encode()

}

func (api *API) GetQuadratic(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Quadratic")

	writer.WriteHeader(200)
	//json.NewEncoder(writer).Encode(DB)

}

func solve(a float64, b float64, c float64) (int, error) {
	var (
		D  float64
		x1 float64
		x2 float64
	)

	if a == 0 && b == 0 {
		//fmt.Println("корней нет")
		return 0, nil
	}

	if a == 0 {
		//fmt.Println("один корень")
		return 1, nil
	}

	D = math.Pow(b, 2) - 4*a*c

	if D < 0 {
		//fmt.Println("корней нет")
		return 0, nil
	}

	if D == 0 {
		//fmt.Println("один корень")
		return 1, nil
	}

	if D > 0 {
		x1 = (-b + math.Sqrt(D)) / 2 / a
		x2 = (-b - math.Sqrt(D)) / 2 / a
		if x1 == x2 {
			//fmt.Println("один корень")
			return 1, nil
		}

		//fmt.Println("два корня")
		return 2, nil
	}

	return 0, nil
}

func (api *API) SolveQuadratic(writer http.ResponseWriter, req *http.Request) {
	initHeaders(writer)
	api.logger.Info("Get Quadratic")

	res, _ := solve(float64(DB["A"]), float64(DB["B"]), float64(DB["C"]))
	q := models.Quadratic{
		A:      DB["A"],
		B:      DB["B"],
		C:      DB["C"],
		Nroots: res,
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(q)
}
