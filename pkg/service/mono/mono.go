package mono

type MonoService struct {
}

func NewMonoService() *MonoService {
	return &MonoService{}
}

func (ms *MonoService) Create() (int, error) {
	return 0, nil
}
