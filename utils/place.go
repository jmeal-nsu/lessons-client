package utils

import "fmt"

type Place struct {
	Cabinet  string `json:"cabinet"`
	Pavilion string `json:"pavilion"`
}

type Places []Place

func (p Places) PrintPlaces() {
	for id, place := range p {
		place.PrintPlace(id)
	}
}

func (p Place) PrintPlace(id int) {
	if p.Pavilion == "" {
		fmt.Println(id, p.Cabinet)
		return
	}
	fmt.Println(id, p.Pavilion, p.Cabinet)
}
