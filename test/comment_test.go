package test

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_createComment(t *testing.T) {

	app := SetUpApp()

	type createCommentResp struct {
		ID       uint   `json:"ID"`
		BlogId   uint   `json:"BlogId"`
		ParentId uint   `json:"ParentId"`
		NickName string `json:"NickName"`
		Content  string `json:"Content"`
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"/api/comment",
		getCreateCommentJson(),
	)

	assert.Equalf(t, true, err == nil, "Create comment request genration failed")

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	res, err := app.Test(req, -1)

	assert.Equalf(t, true, err == nil, "Create comment failed")

	assert.Equalf(t, 200, res.StatusCode, "Create comment doesn't return 200 status code.")

	body, err := ioutil.ReadAll(res.Body)

	assert.Nilf(t, err, "Create comment body parse failed")

	respComment := createCommentResp{}

	err = json.Unmarshal(body, &respComment)
	assert.Nilf(t, err, "Create comment json parse failed")

	assert.Greaterf(t, respComment.ID, uint(0), "Create comment id invaild")
	assert.Equalf(t, "Very first comment", respComment.Content, "Created comment doesn't match with posted comment.")
}

func getCreateCommentJson() io.Reader {
	comment := map[string]interface{}{
		"name":      "Jenish Ja",
		"articleId": 1,
		"commentId": 0,
		"content":   "Very first comment",
	}

	jsonStr, _ := json.Marshal(comment)

	reader := strings.NewReader(string(jsonStr))

	return reader
}
