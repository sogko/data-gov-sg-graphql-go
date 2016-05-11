package datagovsg_test

import (
	"github.com/kr/pretty"
	"github.com/sogko/data-gov-sg-graphql-go/lib/datagovsg"
	"os"
	"testing"
)

var API_KEY string

const TEST_API_URL = "https://api.data.gov.sg/v1/environment/24-hour-weather-forecast"

func init() {
	API_KEY = os.Getenv("DATAGOVSG_API_KEY")
	if API_KEY == "" {
		panic("Set DATAGOVSG_API_KEY environment variable before running test")
	}
}
func TestSimple(t *testing.T) {
	c := datagovsg.NewClient(API_KEY)

	ch := c.Get(TEST_API_URL, &datagovsg.TwentyFourHourWeatherForecastResponse{})

	res := <-ch
	pretty.Println(res.Body)
}

func TestCached(t *testing.T) {
	c := datagovsg.NewClient(API_KEY)

	var ch chan datagovsg.ClientResult
	var ch2 chan datagovsg.ClientResult

	ch = c.Get(TEST_API_URL, &datagovsg.TwentyFourHourWeatherForecastResponse{})

	go func() {
		ch2 = c.Get(TEST_API_URL, &datagovsg.TwentyFourHourWeatherForecastResponse{})
	}()

	res := <-ch
	res2 := <-ch2

	pretty.Println(res.Body)
	pretty.Println(res2.Body)
}
