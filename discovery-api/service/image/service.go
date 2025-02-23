package image

import (
	"context"
	"fmt"
	"github.com/kosa3/pexels-go"
	"log/slog"
)

type pexelImageService struct {
	imageClient *pexels.Client
}

// Create Pexel client
func NewPexelsService(key string) ImageService {
	client := pexels.NewClient(key)

	return &pexelImageService{
		imageClient: client,
	}
}

func (svc *pexelImageService) GetImageURL(search string) (*ImageURl, error) {
	ctx := context.Background()
	photoResponse, err := svc.imageClient.PhotoService.Search(ctx, &pexels.PhotoParams{Query: search, Page: 1, PerPage: 1, Size: "Medium"})
	if err != nil {
		slog.Warn(`Failure on getting image from pexels`)
		return nil, fmt.Errorf("search pexels image error: error in searching %s: %v", search, err)
	}
	fmt.Printf("PlaceInfo at getImageURL:%v", photoResponse)

	photoInfo := &ImageURl{
		ImageID:  photoResponse.Photos[0].ID,
		ImageURL: photoResponse.Photos[0].Src.Large,
	}
	return photoInfo, nil
}
