package view

type viewType string

const (
	Weekly viewType = "Weekly"
)

func (v viewType) String() string {
	return string(v)
}
