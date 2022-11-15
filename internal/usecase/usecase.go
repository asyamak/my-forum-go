package usecase

type UseCase struct {
	User UserUseCase // interface that implements methods of UserREpo interface
	Post Post
}

// connect layers
// layer BL and layer below
// that works with repository
func NewUseCase(u UserRepo, p Post) *UseCase {
	return &UseCase{
		User: newUserUseCase(u),
		Post: NewPostUseCase(p),
	}
}
