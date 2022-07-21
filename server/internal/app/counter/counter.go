package counter

import (
	"fmt"
	"github.com/xamust/qtimTestQuiz/internal/app/model"
	"regexp"
	"strings"
	"sync"
)

type Counter struct {
	Mu      *sync.Mutex
	Config  *Config
	MapChar map[string]int
}

func NewCounter(config *Config, mu *sync.Mutex) *Counter {
	return &Counter{
		Mu:      mu,
		Config:  config,
		MapChar: make(map[string]int),
	}
}

func (c *Counter) CheckRaw(inputModel *model.Request) error {

	//check by onlyChars...
	checkStr := regexp.MustCompile(`^[a-zA-Zа-яА-Я \\t]+$`).MatchString

	if c.Config.WithNumeric {
		//check by onlyCharsAndNumeric...
		checkStr = regexp.MustCompile(`^[a-zA-Zа-яА-Я\d \\t]+$`).MatchString
	}

	if !checkStr(inputModel.Str) {
		return fmt.Errorf("incorrect symbol in input str: %v, numeric support enable: %v", inputModel.Str, c.Config.WithNumeric)
	}

	if !checkStr(inputModel.Char) {
		return fmt.Errorf("incorrect symbol in input char: %v, numeric support enable: %v", inputModel.Char, c.Config.WithNumeric)
	}

	if len([]rune(inputModel.Char)) > 1 {
		return fmt.Errorf("len of input char (%s) must be equal 1", inputModel.Char)
	}

	//count char...
	c.CountChar(inputModel.Str)

	return nil
}

func (c *Counter) CountChar(str string) {
	//"cleanup" map...
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.MapChar = make(map[string]int)
	for _, v := range str {
		if c.Config.CaseSensitive {
			c.MapChar[string(v)] += 1
			continue
		}
		c.MapChar[strings.ToLower(string(v))] += 1
	}
}

func (c *Counter) GetCount(char string) (int, error) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	tempChar := char //for case sensitive....
	if !c.Config.CaseSensitive {
		tempChar = strings.ToLower(char)
	}
	if _, ok := c.MapChar[tempChar]; !ok {
		return 0, fmt.Errorf("char %v not find in str, case senitive enable: %v", char, c.Config.CaseSensitive)
	}
	return c.MapChar[char], nil
}
