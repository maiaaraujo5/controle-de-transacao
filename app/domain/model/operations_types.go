package model

type OperationType int64

const (
	CashPayment OperationType = iota + 1
	InstallmentPayment
	Withdraw
	Payment
)

var operationTypeNames = map[OperationType]string{
	CashPayment:        "Cash Payment",
	InstallmentPayment: "Installment Payment",
	Withdraw:           "Withdraw",
	Payment:            "Payment",
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

func (ot OperationType) IsPayment() bool {
	return ot == Payment
}
