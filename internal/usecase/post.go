package usecase

type PostUseCase struct {
	p Post
}

func NewPostUseCase(p Post) *PostUseCase {
	return &PostUseCase{
		p,
	}
}

func (pu *PostUseCase) Create() {}
func (pu *PostUseCase) Update() {}
func (pu *PostUseCase) Delete() {}
