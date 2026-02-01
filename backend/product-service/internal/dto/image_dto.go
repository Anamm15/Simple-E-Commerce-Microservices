package dto

type ProductImageRequestDTO struct {
	Filename    string
	ContentType string
	Data        []byte
	Size        int64
}

type AddImageProductRequestDTO struct {
	ProductID string
	Images    []ProductImageRequestDTO
}
