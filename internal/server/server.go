package server

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mattiaslndstrm/itineraryparser/internal/api"
)

// Echo specific webserver
type EchoServer struct {
	e    *echo.Echo
	port string
}

type Server interface {
	Start() error
}

func NewServer(port string) Server {
	e := echo.New()
	e.POST("/trips", tripsHandler())
	return EchoServer{e: e, port: port}
}

func (e EchoServer) Start() error {
	return e.e.Start(e.port)
}

// takes the payload, tries to convert it to [][]string, tries to convert it to Trips ([][2]string),
// and convert it to an itinerary ([]string), returns 200 and the result as json if successful,
// else 400 and the error as json
func tripsHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload [][]string
		var err error
		var iternary []string
		var trips api.Trips
		if err = c.Bind(&payload); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if trips, err = payloadToTrips(payload); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		if iternary, err = api.TripsToItinerary(trips); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, iternary)
	}
}

// Converts the payload to Trips ([][2]string), returns an error if any of slices is not of length 2
func payloadToTrips(payload [][]string) (api.Trips, error) {
	var trips api.Trips
	for _, item := range payload {
		if len(item) != 2 {
			return nil, errors.New("every trip must contain a pair of strings")
		}
		trips = append(trips, [2]string{item[0], item[1]})
	}
	return trips, nil
}
