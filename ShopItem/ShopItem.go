package shopitem

import "fmt"

type shopItemHelper struct {
	Category int     `json:"category"`
	Quantity int     `json:"quantity"`
	Cost     float64 `json:"cost"`
	Name     string  `json:"name"`
}

func (s shopItemHelper) String() string {

	return fmt.Sprintf("%s - %v\n", s.Name, s.Cost)

}
