package app

type ServiceApp struct{}

func NewServiceApp() ServiceApp {
	return ServiceApp{}
}

func (ServiceApp) Run(opts Options) error {

}
