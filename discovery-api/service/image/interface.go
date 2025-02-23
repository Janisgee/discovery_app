package image

type ImageService interface {
	GetImageURL(search string) (*ImageURl, error)
}
