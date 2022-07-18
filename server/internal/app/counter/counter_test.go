package counter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xamust/qtimTestQuiz/internal/app/model"
	"testing"
)

func initCounter() *Counter {
	return &Counter{
		Config: &Config{
			CaseSensitive: false,
			WithNumeric:   false,
		},
		MapChar: make(map[string]int),
	}
}

func TestCounter_CheckRaw(t *testing.T) {

	counter := initCounter()

	tests := []struct {
		name       string
		inputModel *model.Request
		err        error
	}{
		{
			name:       "Correct input model",
			inputModel: &model.Request{Str: "hello world", Char: "o"},
			err:        nil,
		}, {
			name:       "Correct input model",
			inputModel: &model.Request{Str: "Вася полетел на Луну", Char: "л"},
			err:        nil,
		},
		{
			name:       "Incorrect input model",
			inputModel: &model.Request{Str: "Вася пол3т3л на Луну", Char: "л"},
			err:        fmt.Errorf("incorrect symbol in input str: Вася пол3т3л на Луну, numeric support enable: false"),
		}, {
			name:       "Incorrect input model",
			inputModel: &model.Request{Str: "Вася полетел на Луну", Char: "ла"},
			err:        fmt.Errorf("len of input char (ла) must be equal 1"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := counter.CheckRaw(test.inputModel); err != nil {
				assert.EqualError(t, counter.CheckRaw(test.inputModel), test.err.Error())
			} else {
				assert.NoError(t, counter.CheckRaw(test.inputModel))
			}
		})
	}
}

func TestCounter_CountChar(t *testing.T) {

	counter := initCounter()

	tests := []struct {
		name     string
		inputStr string
		count    int
	}{
		{
			name:     "Test CountChar",
			inputStr: "hello world",
			count:    8,
		}, {
			name:     "Test CountChar",
			inputStr: "Вася полетел на Луну",
			count:    12,
		},
		{
			name:     "Test CountChar",
			inputStr: "Вася пол3т3л на Луну",
			count:    12,
		}, {
			name:     "Test CountChar",
			inputStr: "Вася полетел на Луну",
			count:    12,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter.CountChar(test.inputStr)
			assert.Equal(t, test.count, len(counter.MapChar))
		})
	}
}

func TestCounter_GetCount(t *testing.T) {

	counter := initCounter()

	tests := []struct {
		name      string
		inputChar string
		inputStr  string
		result    int
		err       error
	}{
		{
			name:      "Test GetCount",
			inputChar: "o",
			inputStr:  "hello world",
			result:    2,
			err:       nil,
		},
		{
			name:      "Test GetCount",
			inputChar: "л",
			inputStr:  "Вася полетел на Луну",
			result:    3,
			err:       nil,
		},
		{
			name:      "Test GetCount",
			inputChar: "3",
			inputStr:  "hello world",
			result:    0,
			err:       fmt.Errorf("char 3 not find in str, case senitive enable: false"),
		},
		{
			name:      "Test GetCount",
			inputChar: "ы",
			inputStr:  "Вася полетел на Луну",
			result:    0,
			err:       fmt.Errorf("char ы not find in str, case senitive enable: false"),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			counter.CountChar(test.inputStr)
			intChar, err := counter.GetCount(test.inputChar)
			if err != nil {
				assert.EqualError(t, err, test.err.Error())
			} else {
				assert.Equal(t, intChar, test.result)
			}
		})
	}
}
