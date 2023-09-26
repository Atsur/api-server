package controllers

import (
	// "errors"
	// "github.com/atsur/api-server/pkg/http_err"

	"log"
	"net/http"
	"os"

	"github.com/skip2/go-qrcode"
	//models "github.com/atsur/api-server/internal/pkg/models"
	//"github.com/atsur/api-server/internal/pkg/persistence"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	//"github.com/atsur/api-server/pkg/crypto"
	awsPkg "github.com/atsur/api-server/pkg/aws"

	"github.com/gin-gonic/gin"
)

type QRInput struct {
	UUID     string `json:"uuid" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

var (
	AccessKeyID     = os.Getenv("AWS_ACCESS_KEY_ID")
	SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	MyRegion        = os.Getenv("AWS_REGION")
	MyBucket        = os.Getenv("BUCKET_NAME")
)

// GetUserById godoc
// @Summary Retrieves user based on given ID
// @Description get User by ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} users.User
// @Router /api/users/{id} [get]
// @Security Authorization Token
func CreateQR(c *gin.Context) {

	// fetch entry
	// s := persistence.GetUserRepository()
	// id := c.Param("entry-id")

	// if user, err := s.Get(entry-id); err != nil {
	// 	http_err.NewError(c, http.StatusNotFound, errors.New("user not found"))
	// 	log.Println(err)
	// } else {
	// 	c.JSON(http.StatusOK, user)
	// }


	// //generate QR code
	// qr, err := qrsvg.New("https://registry.atsur.art/entry/00012")
	// if err != nil {
	//    log.Fatalln("failed to generate qr ", err)
	// }
	// buff := new(bytes.Buffer)
	// // encode image to buffer
	// err = svg.Encode(buff, qr)
	// if err != nil {
	// 	log.Fatalln("failed to create buffer ", err)
	// }
	// // convert buffer to reader
	// reader := bytes.NewReader(buff.Bytes())


	err := qrcode.WriteFile("https://registry.atsur.art/entry/123456", qrcode.Medium, 256, "qr.png")
	if err != nil {
		log.Fatalln("failed to create file ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to create file",
		})
		return
	}
	//file, _, err := c.Request.FormFile("qr.png")
	//filename := header.Filename
	file, err := os.Open("qr.png")
    if err != nil {
		log.Fatalln("failed to read file ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to read file",
		})
		return
	}
	defer file.Close()

	filename := "new QR code"
    r := file

	// start aws session
	sess, err := awsPkg.InitAws()
	if err != nil {
		log.Fatalln("failed to init aws session", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to start upload session",
		})
		return
	}
	uploader := s3manager.NewUploader(sess)
	contentType := "image/png"

	//upload to the s3 bucket
	up, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(MyBucket),
		// ACL:    aws.String("public-read"),
		Key:    aws.String(filename),
		ContentType: &contentType,
		Body: r,
	})

	if err != nil {
		log.Fatalln("failed to upload file ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":    "Failed to upload file",
			"uploader": up,
		})
		return
	}
	filepath := "https://" + MyBucket + "." + "s3-" + MyRegion + ".amazonaws.com/" + filename
	c.JSON(http.StatusOK, gin.H{
		"filepath": filepath,
	})

}
