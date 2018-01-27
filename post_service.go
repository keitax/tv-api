package tvapi

type PostService interface {
	GetList(start int, results int) ([]*Post, error)
	GetSingle(id string) (*Post, error)
	Synchronize() error
}

type postService struct{}

func NewPostService() {
}

func (pc *postService) GetList(start int, results int) ([]*Post, error) {
	return nil, nil
}

func (pc *postService) GetSingle(id string) (*Post, error) {
	return nil, nil
}

func (pc *postService) Synchronize() error {
	return nil
}
