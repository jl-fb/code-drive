package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

//Definir estrutura rota que o sistema utilizará
type Route struct {
	ID        string
	ClientID  string
	Positions []Position
}

type Position struct {
	Lat float64
	Lng float64
}

//definir interface para rota
type RouteInterface interface {
	LoadPositions() error
}

func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New(ErrorRouteIDNotInformed)
	}
	f, err := os.Open("destinations" + r.ID + ".txt")
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
