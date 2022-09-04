package exception

import "fmt"

type DataNotFoundException struct {
	Name string
}

func NewDataNotFoundException(name string) *DataNotFoundException {
	return &DataNotFoundException{Name: name}
}

func (e *DataNotFoundException) Error() string {
	return fmt.Sprintf("%s not found", e.Name)
}
