package domain

type FileServiceInterface interface {
	GetPath(path string) string
}

type FileService struct {
	FileRepository FileRepositoryInterface
}

func NewFileService(file_repository FileRepositoryInterface) *FileService {
	return &FileService{
		FileRepository: file_repository,
	}
}

func (s *FileService) GetPath(path string) string {
	return s.FileRepository.GetPath(path)
}
