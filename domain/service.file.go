package domain

type FileServiceInterface interface {
	GetPath(path string) string
}

type FileService struct {
	FileRepository FileRepositoryInterface
}

func NewFileService() *FileService {
	return &FileService{
		FileRepository: NewFileRepository(),
	}
}

func (s *FileService) GetPath(path string) string {
	return s.FileRepository.GetPath(path)
}
