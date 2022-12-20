package interfaces

type GoWalletManageSvrImpl struct{}

func NewWalletManageSvrImpl() *GoWalletManageSvrImpl {
	return &GoWalletManageSvrImpl{}
}

func InitializeService() *GoWalletManageSvrImpl {
	return NewWalletManageSvrImpl()
}
