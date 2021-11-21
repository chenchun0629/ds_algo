package hashtable

import "fmt"

type Valuable interface {
	GetKey() string
}

type IntValue int

func (i IntValue) GetKey() string {
	return fmt.Sprintf("%d", i)
}

type StringValue string

func (s StringValue) GetKey() string {
	return string(s)
}
