package server

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/xamust/qtimTestQuiz/internal/app/counter"
	"github.com/xamust/qtimTestQuiz/internal/app/model"
	"io/ioutil"
	"net/http"
)

type Handlers struct {
	logger  *logrus.Logger
	counter *counter.Counter
}

//handle...
func (h *Handlers) Detect(w http.ResponseWriter, r *http.Request) {
	//check correct http method
	if r.Method == http.MethodPost {
		counter := h.counter
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			h.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		var strRequest model.Request
		var strResponse model.Response

		//unmarshaling...
		if err := json.Unmarshal(content, &strRequest); err != nil {
			h.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//check data and fill map...
		if err := counter.CheckRaw(&strRequest); err != nil {
			h.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//get count...
		resultInt, err := counter.GetCount(strRequest.Char)
		if err != nil {
			h.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		strResponse.Count = resultInt

		result, err := json.Marshal(strResponse)
		if err != nil {
			h.logger.Error(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		//return page with count...
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(result)
		h.logger.Infof("Request: %+v, response %+v\n", strRequest, strResponse)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}
