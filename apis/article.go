package api

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"

	lConf "go-blogs-api/configs"
	lModels "go-blogs-api/models"
)

func SetArticleRoutes(api fiber.Router) {
	articleApi := api.Group("/article")
	articleApi.Get("/", getArticles)
	articleApi.Get("/:id", getArticle)
	articleApi.Get("/:id/comments", getArticleComments)
	articleApi.Post("/", createArticle)
}

func getArticles(c *fiber.Ctx) error {
	skip := c.Query("skip", "0")
	order := c.Query("order_by", "title")
	orderDir := c.Query("order_dir", "asc")

	offset, er := strconv.Atoi(skip)

	if er != nil {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Invalid value in skip query parameters"))
	}

	type ArticleListItem struct {
		ID        uint      `json:"id"`
		NickName  string    `json:"name"`
		Title     string    `json:"title"`
		CreatedAt time.Time `json:"created_at"`
	}

	var articles []ArticleListItem

	orderByClaus := clause.OrderByColumn{
		Column: clause.Column{Name: order},
		Desc:   orderDir == "desc",
	}

	lConf.DBO.
		Model(&lModels.Article{}).
		Order(orderByClaus).
		Offset(offset).
		Limit(20).
		Find(&articles)

	return c.JSON(articles)
}

func getArticle(c *fiber.Ctx) error {
	id := c.Params("id", "")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Invalid route parameters"))
	}

	var article lModels.Article

	lConf.DBO.First(&article, "id = ?", id)

	if article.ID == 0 {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Article not found for this id."))
	}

	return c.Status(fiber.StatusOK).JSON(article)
}

func createArticle(c *fiber.Ctx) error {

	payload := struct {
		Name    string `json:"name"`
		Title   string `json:"title"`
		Content string `json:"content"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return err
	}

	if len(payload.Name) == 0 || len(payload.Title) == 0 || len(payload.Content) == 0 {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Invalid parameters"))
	}

	article := lModels.Article{NickName: payload.Name, Title: payload.Title, Content: payload.Content}
	r := lConf.DBO.Create(&article)

	if r.Error != nil {
		return r.Error
	}

	return c.JSON(article)

}

func getArticleComments(c *fiber.Ctx) error {
	id := c.Params("id", "")

	if id == "" {
		return c.Status(fiber.StatusBadRequest).Send([]byte("Invalid route parameters"))
	}

	var comments []lModels.Comment

	lConf.DBO.Find(&comments, "blog_id = ?", id)

	return c.JSON(comments)
}
