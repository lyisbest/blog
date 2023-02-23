package constant

const (
	Create = 1
	Update = 2
	Delete = 3
)

var mapOperationTypeToString map[int]string = map[int]string{Create: "create", Update: "update", Delete: "delete"}
