package cloud

import (
	"context"
	"errors"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// CloudinaryService defines the contract for file operations.
// This interface can be moved to the service layer if you want strict strict hexagonal architecture.
type CloudinaryService interface {
	UploadFile(ctx context.Context, file interface{}, filename string) (string, string, error)
	DeleteFile(ctx context.Context, publicID string) error
}

type cloudinaryService struct {
	cld          *cloudinary.Cloudinary
	uploadFolder string
}

func NewCloudinaryService(cloudName, apiKey, apiSecret, uploadFolder string) (CloudinaryService, error) {
	cld, err := cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		return nil, err
	}

	return &cloudinaryService{
		cld:          cld,
		uploadFolder: uploadFolder,
	}, nil
}

// UploadFile handles uploading a file (io.Reader or path) to Cloudinary.
// It returns the Secure URL of the uploaded asset.
func (s *cloudinaryService) UploadFile(ctx context.Context, file interface{}, filename string) (string, string, error) {
	// 1. Set upload parameters
	// public_id: Custom name for the file (optional, but recommended for tracking)
	// folder: Organizing assets into specific directories
	overwrite := true
	uploadParams := uploader.UploadParams{
		PublicID:     filename,
		Folder:       s.uploadFolder,
		ResourceType: "auto",     // Automatically detect image/video/raw
		Overwrite:    &overwrite, // Overwrite if same public_id exists
	}

	// 2. Execute upload
	// 'file' can be a path (string) or an io.Reader (multipart.File)
	resp, err := s.cld.Upload.Upload(ctx, file, uploadParams)
	if err != nil {
		return "", "", err
	}

	// 3. Return the secure URL (HTTPS)
	return resp.SecureURL, resp.PublicID, nil
}

// DeleteFile removes an asset from Cloudinary using its Public ID.
func (s *cloudinaryService) DeleteFile(ctx context.Context, publicID string) error {
	// 1. Set destroy parameters
	Invalidate := true
	destroyParams := uploader.DestroyParams{
		PublicID:     publicID,
		ResourceType: "image",     // Default to image, change if you handle videos
		Invalidate:   &Invalidate, // Invalidate CDN cache
	}

	// 2. Execute destroy
	resp, err := s.cld.Upload.Destroy(ctx, destroyParams)
	if err != nil {
		return err
	}

	// 3. Check result status
	if resp.Result != "ok" {
		return errors.New("failed to delete file from cloud storage")
	}

	return nil
}
