package service

// CmdServiceImpl is cmd service impl
type CmdServiceImpl interface {
	SetParam(args []string) error
	Validate() error
	Execute() error
}
