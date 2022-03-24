package parser

import (
	"airplane-seating/internal/seat/seatmap"
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type InputParser struct {
	reader *bufio.Reader
}

func NewInputParser(r *bufio.Reader) *InputParser {
	return &InputParser{
		reader: r,
	}
}

func (p *InputParser) ParseInput() (seatmap.Layout, int, error) {
	fmt.Printf("Enter the layout: ")
	inputLayout, err := p.reader.ReadString('\n')
	inputLayout = strings.TrimSpace(inputLayout)
	if err != nil {
		return nil, 0, err
	}

	fmt.Printf("Enter the number of passengers: ")
	inputNumberOfPassengers, err := p.reader.ReadString('\n')
	inputNumberOfPassengers = strings.TrimSpace(inputNumberOfPassengers)
	if err != nil {
		return nil, 0, err
	}

	layout, err := convertStringToLayout(inputLayout)
	if err != nil {
		return nil, 0, err
	}
	numberOfPassengers, err := convertStringToNumberOfPassengers(inputNumberOfPassengers)
	if err != nil {
		return nil, 0, err
	}

	return layout, numberOfPassengers, nil
}

func convertStringToNumberOfPassengers(passengers string) (int, error) {
	n, err := strconv.ParseInt(passengers, 10, 8)
	if n < 0 {
		return 0, errors.New("invalid number of passengers, please provide a positive number")
	}
	return int(n), err
}

func convertStringToLayout(inputString string) (seatmap.Layout, error) {
	layoutGroups := make(seatmap.Layout, 0)
	sanitised := separateIndividual1DArrayElements(inputString)
	for _, s := range sanitised {
		split := strings.Split(s, ",")
		r := strings.TrimSpace(split[1])
		c := strings.TrimSpace(split[0])

		row, err := strconv.ParseInt(r, 10, 8)
		if err != nil {
			return nil, err
		}
		column, err := strconv.ParseInt(c, 10, 8)
		if err != nil {
			return nil, err
		}
		layoutGroups = append(layoutGroups, []int{int(row), int(column)})
	}

	return layoutGroups, nil

}

func separateIndividual1DArrayElements(input string) []string {
	input = strings.Replace(input, "[[", "", 1)
	input = strings.Replace(input, "]]", "", 1)

	split := strings.Split(input, "],")
	var sanitised []string
	for _, s := range split {
		s = strings.TrimSpace(s)
		s = strings.Replace(s, "[", "", -1)
		s = strings.Replace(s, "]", "", -1)
		sanitised = append(sanitised, s)
	}

	return sanitised
}
