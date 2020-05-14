package wallet

type Tx interface {
}

type Address interface {
}

type Wallet interface {
	SendAmtTo(amt float64, address Address) Tx
	Balance() float64
}
