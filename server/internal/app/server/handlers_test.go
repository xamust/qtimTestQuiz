package server

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/xamust/qtimTestQuiz/internal/app/counter"
	"net/http"
	"net/http/httptest"
	"testing"
)

func initCounter() *counter.Counter {
	return &counter.Counter{
		Config: &counter.Config{
			CaseSensitive: false,
			WithNumeric:   false,
		},
		MapChar: make(map[string]int),
	}
}

func TestServer_Detect(t *testing.T) {

	//init new logrus with panic level, for ignore logger...
	lg := logrus.New()
	lg.SetLevel(logrus.PanicLevel)

	//init Handlers with test mock...
	serverTest := Handlers{lg, initCounter()}
	handlerTest := http.HandlerFunc(serverTest.Detect)

	testCases := []struct {
		name   string
		method string
		body   string
		want   []byte
	}{
		{
			name:   "Testing detect handler...",
			method: http.MethodPost,
			body:   fmt.Sprintf(`{"str" : "%s", "char" : "%s"}`, "hello world", "o"),
			want:   []byte(fmt.Sprintf(`{"count":%d}`, 2)),
		},
		{
			name:   "Testing detect handler...",
			method: http.MethodPost,
			body:   fmt.Sprintf(`{"str" : "%s", "char" : "%s"}`, "Вася полетел на Луну", "л"),
			want:   []byte(fmt.Sprintf(`{"count":%d}`, 3)),
		},
		{
			name:   "Testing detect handler...",
			method: http.MethodPost,
			body:   fmt.Sprintf(`{"str" : "%s", "char" : "%s"}`, "Вася пол3т3л на Луну", "л"),
			want:   []byte(fmt.Sprintf("incorrect symbol in input str: Вася пол3т3л на Луну, numeric support enable: false")),
		},
		{
			name:   "Testing detect handler...",
			method: http.MethodPost,
			body:   fmt.Sprintf(`{"str" : "%s", "char" : "%s"}`, "Вася полетел на Луну", "ла"),
			want:   []byte(fmt.Sprintf("len of input char (ла) must be equal 1")),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, err := http.NewRequest(tc.method, "/detect", bytes.NewBuffer([]byte(tc.body)))
			req.Header.Add("Content-Type", "application/json")
			handlerTest.ServeHTTP(rec, req)
			assert.NoError(t, err)
			assert.Equal(t, tc.want, rec.Body.Bytes())
		})
	}
}
