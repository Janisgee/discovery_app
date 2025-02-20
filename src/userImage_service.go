package main

// Import Cloudinary and other necessary libraries
//===================
import (
	"context"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryService struct {
	cloudClient  *cloudinary.Cloudinary
	cloudContext context.Context
}

// Create a google client
func NewCloudinaryService() (*CloudinaryService, error) {
	// Create a cloudinary instant
	cloudSvc, err := cloudinary.New()
	if err != nil {
		return nil, err
	}
	cloudSvc.Config.URL.Secure = true
	ctx := context.Background()
	return &CloudinaryService{
		cloudClient:  cloudSvc,
		cloudContext: ctx,
	}, nil
}

// Upload Image
func (svc *CloudinaryService) UploadImage(imageURL string) (string, error) {

	// Upload the image.
	// Set the asset's public ID and allow overwriting the asset with new versions
	resp, err := svc.cloudClient.Upload.Upload(svc.cloudContext, imageURL, uploader.UploadParams{
		PublicID:       "defaultPublicID",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		fmt.Println("error")
		return "", err
	}

	// Log the delivery URL
	fmt.Println("****2. Upload an image****\nDelivery URL:", resp.SecureURL, "\n")
	return resp.SecureURL, err
}

// Get and use image
func (svc *CloudinaryService) getAssetInfo(publicID string) ([]string, error) {
	// Get and use details of the image
	// ==============================
	resp, err := svc.cloudClient.Admin.Asset(svc.cloudContext, admin.AssetParams{PublicID: publicID})
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println("****3. Get and use details of the image****\nDetailed response:\n", resp, "\n")

	// Assign tags to the uploaded image based on its width. Save the response to the update in the variable 'update_resp'.
	if resp.Width > 900 {
		update_resp, err := svc.cloudClient.Admin.UpdateAsset(svc.cloudContext, admin.UpdateAssetParams{
			PublicID: "quickstart_butterfly",
			Tags:     []string{"large"}})
		if err != nil {
			fmt.Println("error")
			return nil, err
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", update_resp.Tags, "\n")
			return update_resp.Tags, err
		}
	} else {
		update_resp, err := svc.cloudClient.Admin.UpdateAsset(svc.cloudContext, admin.UpdateAssetParams{
			PublicID: "quickstart_butterfly",
			Tags:     []string{"small"}})
		if err != nil {
			fmt.Println("error")
			return nil, err
		} else {
			// Log the new tag to the console.
			fmt.Println("New tag: ", update_resp.Tags, "\n")
			return update_resp.Tags, err

		}

	}
}

// Transform image
func (svc *CloudinaryService) TransformImage(public_ID string) (string, error) {
	// Instantiate an object for the asset with public ID "my_image"
	qs_img, err := svc.cloudClient.Image(public_ID)
	if err != nil {
		return "", err
	}

	// Add the transformation
	qs_img.Transformation = "f_auto/q_auto/c_fill,g_face,h_250,w_250"

	// Generate and log the delivery URL
	new_url, err := qs_img.String()
	if err != nil {
		return "", err
	} else {
		print("****4. Transform the image****\nTransfrmation URL: ", new_url, "\n")
		return new_url, err
	}
}

// Run code
// func main() {
// 	cld, ctx := credentials()
// 	uploadImage(cld, ctx)
// 	getAssetInfo(cld, ctx)
// 	transformImage(cld, ctx)
// }
