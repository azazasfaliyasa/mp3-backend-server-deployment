package customer

type UseCase struct {
	R RepositoryInterface
}

type UseCaseInterface interface {
	ShowAll() ([]Customer, error)
	Create(customer *Customer) error
	Show(ID string) (*Customer, error)
	Update(ID string, customer Customer) (*Customer, error)
	Destroy(ID string) (*Customer, error)
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		R: r,
	}
}

func (u *UseCase) ShowAll() ([]Customer, error) {
	return u.R.ShowAll()
}

func (u *UseCase) Create(customer *Customer) error {
	return u.R.Create(customer)
}

func (u *UseCase) Show(ID string) (*Customer, error) {
	return u.R.Show(ID)
}

func (u *UseCase) Update(ID string, customer Customer) (*Customer, error) {
	return u.R.Update(ID, customer)
}

func (u *UseCase) Destroy(ID string) (*Customer, error) {
	return u.R.Destroy(ID)
}
