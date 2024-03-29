package model

type OperationType int64

const (
	CompraAVista OperationType = iota + 1
	CompraParcelada
	Saque
	Pagamento
)

var operationTypeNames = map[OperationType]string{
	CompraAVista:    "Compra a Vista",
	CompraParcelada: "Compra Parcelada",
	Saque:           "Saque",
	Pagamento:       "Pagamento",
}

func (ot OperationType) String() string {
	name, ok := operationTypeNames[ot]
	if !ok {
		return "Unknown"
	}
	return name
}

func (ot OperationType) IsValid() bool {
	_, ok := operationTypeNames[ot]
	return ok
}

func (ot OperationType) Index() int64 {
	return int64(ot)
}
