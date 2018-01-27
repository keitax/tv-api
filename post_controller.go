package tvapi

type PostController interface {
	GetList(start int, results int) ([]*Post, error)
	GetSingle(id string) (*Post, error)
	Synchronize() error
}

type postController struct {
	postService PostService
}

func NewPostController(postService PostService) PostController {
	return &postController{
		postService: postService,
	}
}

func (pc *postController) GetList(start int, results int) ([]*Post, error) {
	return pc.postService.GetList(start, results)
}

func (pc *postController) GetSingle(id string) (*Post, error) {
	return pc.postService.GetSingle(id)
}

func (pc *postController) Synchronize() error {
	return pc.postService.Synchronize()
}
