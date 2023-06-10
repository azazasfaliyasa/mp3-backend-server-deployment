package user

type UseCase struct {
	R RepositoryInterface
}
type UseCaseInterface interface {
	ShowAll() ([]User, error)
	Create(user *User) error
	Show(ID string) (*User, error)
	Update(ID string, user User) (*User, error)
	Destroy(ID string) (*User, error)
}

func NewUseCase(r RepositoryInterface) UseCaseInterface {
	return &UseCase{
		R: r,
	}
}

func (u *UseCase) ShowAll() ([]User, error) {
	return u.R.ShowAll()
}

func (u *UseCase) Create(user *User) error {
	return u.R.Create(user)
}

func (u *UseCase) Show(ID string) (*User, error) {
	return u.R.Show(ID)
}

func (u *UseCase) Update(ID string, user User) (*User, error) {
	return u.R.Update(ID, user)
}

func (u *UseCase) Destroy(ID string) (*User, error) {
	return u.R.Destroy(ID)
}
