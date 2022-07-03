package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

//definir interface para rota
type RouteInterface interface {
	LoadPositions() error
	ExportJSONPositions() ([]string, error)
}

func NewRoute() RouteInterface {
	return &Route{
		ID:        "",
		ClientID:  "",
		Positions: []Position{},
	}
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New(ErrorRouteIDNotInformed)
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return err
		}
		lng, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return err
		}
		r.Positions = append(r.Positions, Position{Lat: lat, Lng: lng})
	}

	return nil
}

// ExportJSONPositions implements RouteInterface
func (r *Route) ExportJSONPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)
	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Lng}
		route.Finished = false
		if total-1 == k {
			route.Finished = true
		}
		jsonRoute, err := json.Marshal(route)
		if err != nil {
			return nil, err
		}
		result = append(result, string(jsonRoute))
	}
	return result, nil
}
