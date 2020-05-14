package wallet

type Tx interface {
}

type Address interface {
}

type Wallet interface {
	SendAmtTo(amt float64, address Address) Tx
	Balance() float64
}

// see: https://medium.com/myetherwallet/hd-wallets-and-derivation-paths-explained-865a643c7bf2
var path string = "m/44'/60'/0'/0"
