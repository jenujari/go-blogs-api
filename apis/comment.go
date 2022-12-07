package apis

import (
	"github.com/gofiber/fiber/v2"

	lConf "go-blogs-api/configs"
	lModels "go-blogs-api/models"
)

func SetCommentRoutes(api fiber.Router) {
	commentApi := api.Group("/comment")
	commentApi.Post("/", createComment)
}

func createComment(c *fiber.Ctx) error {
	payload := struct {
		Name    string `json:"name"`
		Article uint   `json:"articleId"`
		Comment uint   `json:"commentId"`
		Content string `json:"content"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if len(payload.Name) == 0 || payload.Article == 0 || len(payload.Content) == 0 {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Invalid parameters"))
	}

	comment := lModels.Comment{
		BlogId:   payload.Article,
		Content:  payload.Content,
		NickName: payload.Name,
		ParentId: 0,
	}

	if payload.Comment > 0 {
		comment.ParentId = payload.Comment
	}

	r := lConf.DBO.Create(&comment)

	if r.Error != nil {
		return r.Error
	}

	return c.JSON(comment)
}
