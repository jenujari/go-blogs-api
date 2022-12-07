package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	lConf "go-blogs-api/configs"
	lUtls "go-blogs-api/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func SetUpApp() *fiber.App {
	absPath, _ := filepath.Abs("./../")
	os.Setenv("ABS_PATH", absPath)

	lUtls.SetGlobals()
	lConf.InitDB(true)

	app := lConf.InitFiber()

	lUtls.InstallRouter(app)

	return app
}

func Test_getArticleComments(t *testing.T) {
	app := SetUpApp()

	type articleCommentsResp struct {
		ID       uint   `json:"id"`
		Name     string `json:"name"`
		Content  string `json:"content"`
		BlogId   uint   `json:"blog_id"`
		ParentId uint   `json:"parent_id"`
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/comment",
		getCreateCommentJson(),
	)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	app.Test(req, -1)

	req, err := http.NewRequest(
		http.MethodGet,
		"/api/article/1/comments",
		nil,
	)

	assert.Equalf(t, true, err == nil, "Get comments request genration failed")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := app.Test(req, -1)

	assert.Equalf(t, true, err == nil, "Get comments failed")

	assert.Equalf(t, 200, res.StatusCode, "Get comments doesn't return 200 status code.")

	body, err := ioutil.ReadAll(res.Body)

	assert.Nilf(t, err, "Get comments body parse failed")

	respComments := []articleCommentsResp{}

	err = json.Unmarshal(body, &respComments)
	assert.Nilf(t, err, "Response comments json parse failed")

	assert.Greaterf(t, len(respComments), int(0), "Comments response return empty array.")

	respComment := respComments[0]

	assert.Greaterf(t, respComment.ID, uint(0), "Response comments have invalid id.")
}

func Test_createArticle(t *testing.T) {

	app := SetUpApp()

	type createArticleResp struct {
		ID       uint   `json:"ID"`
		Title    string `json:"Title"`
		NickName string `json:"NickName"`
		Content  string `json:"Content"`
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"/api/article/",
		getCreateArticleJson(),
	)

	assert.Equalf(t, true, err == nil, "Create article request genration failed")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := app.Test(req, -1)

	assert.Equalf(t, true, err == nil, "Create article failed")

	assert.Equalf(t, 200, res.StatusCode, "Create article doesn't return 200 status code.")

	body, err := ioutil.ReadAll(res.Body)

	assert.Nilf(t, err, "Create article body parse failed")

	respArticle := createArticleResp{}

	err = json.Unmarshal(body, &respArticle)
	assert.Nilf(t, err, "Create article json parse failed")

	assert.Greaterf(t, respArticle.ID, uint(0), "Create article id invaild")
}

func Test_getArticle(t *testing.T) {

	app := SetUpApp()

	type getArticleResp struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		NickName string `json:"name"`
		Content  string `json:"content"`
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/article/",
		getCreateArticleJson(),
	)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	app.Test(req, -1)

	req, _ = http.NewRequest(
		http.MethodGet,
		"/api/article/1",
		nil,
	)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := app.Test(req, -1)

	assert.Equalf(t, true, err == nil, "Get article failed")

	assert.Equalf(t, 200, res.StatusCode, "Get article doesn't return 200 status code.")

	body, err := ioutil.ReadAll(res.Body)

	assert.Nilf(t, err, "Get article body parse failed")

	respArticle := getArticleResp{}

	err = json.Unmarshal(body, &respArticle)
	assert.Nilf(t, err, "Get article json parse failed")

	assert.Greaterf(t, respArticle.ID, uint(0), "Get article id invaild")
}

func Test_getArticles(t *testing.T) {

	app := SetUpApp()

	type getArticlesResp struct {
		ID       uint   `json:"id"`
		Title    string `json:"title"`
		NickName string `json:"name"`
	}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/api/article/",
		getCreateArticleJson(),
	)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	app.Test(req, -1)

	req, _ = http.NewRequest(
		http.MethodGet,
		"/api/article?order_dir=desc&order_by=title&skip=0",
		nil,
	)

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := app.Test(req, -1)

	assert.Equalf(t, true, err == nil, "Get articles failed")

	assert.Equalf(t, 200, res.StatusCode, "Get articles doesn't return 200 status code.")

	body, err := ioutil.ReadAll(res.Body)

	assert.Nilf(t, err, "Get article body parse failed")

	respArticles := []getArticlesResp{}

	err = json.Unmarshal(body, &respArticles)
	assert.Nilf(t, err, "Get article json parse failed")

	assert.Greaterf(t, len(respArticles), int(0), "Get articles response return empty array.")

	respArticle := respArticles[0]

	assert.Greaterf(t, respArticle.ID, uint(0), "Get articles first ele id invaild")
}

func getCreateArticleJson() io.Reader {
	article := map[string]interface{}{
		"name":    "Jenish Ja",
		"title":   "My Title",
		"content": "My Content",
	}

	jsonStr, _ := json.Marshal(article)

	reader := strings.NewReader(string(jsonStr))

	return reader
}
