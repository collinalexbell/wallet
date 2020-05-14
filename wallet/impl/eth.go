package eth

import "github.com/ethereum/go-ethereum/accounts"

var DefaultRootDerivationPath = accounts.DefaultRootDerivationPath

// see: https://medium.com/myetherwallet/hd-wallets-and-derivation-paths-explained-865a643c7bf2
var path string = "m/44'/60'/0'/0"

type Wallet struct {
}
