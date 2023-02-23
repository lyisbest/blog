package test

import (
	"blog/apps/blog/request"
	"blog/configuration"
	"blog/router"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r *gin.Engine
var id int

func init() {
	err := configuration.InitMySQL()
	if err != nil {
		return
	}
	r = gin.Default()
	r = router.SetRouters(r)
}

var blogCreateRequest request.BlogCreateRequest = request.BlogCreateRequest{
	Title:   "333",
	Content: "44444",
}

var listBlogRequest request.ListBlogRequest = request.ListBlogRequest{
	PageNum: 1,
}

func Test_CreateBlog(t *testing.T) {
	b, _ := json.Marshal(blogCreateRequest)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/blog/create", bytes.NewReader(b))
	req.AddCookie(&http.Cookie{
		Name:  "user_cookie",
		Value: "root",
	})
	r.ServeHTTP(w, req)

	var ans map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &ans)
	code := ans["Code"].(float64)
	message := ans["Message"].(string)
	data := ans["Data"]

	assert.Equal(t, 0, int(code))
	assert.Equal(t, "success", message)
	assert.Equal(t, nil, data)
}

func Test_ListBlog(t *testing.T) {
	b, _ := json.Marshal(listBlogRequest)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blog/list", bytes.NewReader(b))
	r.ServeHTTP(w, req)

	var ans map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &ans)
	code := ans["Code"].(float64)
	message := ans["Message"].(string)
	list := ans["Data"].([]interface{})[0]
	data := list.(map[string]interface{})
	title := data["title"].(string)
	content := data["content"].(string)
	id = int(data["id"].(float64))

	assert.Equal(t, 0, int(code))
	assert.Equal(t, "success", message)
	assert.Equal(t, blogCreateRequest.Title, title)
	assert.Equal(t, blogCreateRequest.Content, content)

}

func Test_GetBlog(t *testing.T) {
	var blogGetRequest request.BlogGetRequest = request.BlogGetRequest{
		Id: id,
	}
	b, _ := json.Marshal(blogGetRequest)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/blog/get", bytes.NewReader(b))
	r.ServeHTTP(w, req)

	var ans map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &ans)
	code := ans["Code"].(float64)
	message := ans["Message"].(string)
	data := ans["Data"].(map[string]interface{})
	title := data["title"].(string)
	content := data["content"].(string)

	assert.Equal(t, 0, int(code))
	assert.Equal(t, "success", message)
	assert.Equal(t, blogCreateRequest.Title, title)
	assert.Equal(t, blogCreateRequest.Content, content)
}

func Test_UpdateBlog(t *testing.T) {
	var blogUpdateRequest request.BlogUpdateRequest = request.BlogUpdateRequest{
		Id:      id,
		Title:   "333修改",
		Content: "444修改",
	}
	b, _ := json.Marshal(blogUpdateRequest)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/blog/update", bytes.NewReader(b))
	req.AddCookie(&http.Cookie{
		Name:  "user_cookie",
		Value: "root",
	})
	r.ServeHTTP(w, req)

	var ans map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &ans)
	code := ans["Code"].(float64)
	message := ans["Message"].(string)

	assert.Equal(t, 0, int(code))
	assert.Equal(t, "success", message)

	br := request.BlogGetRequest{
		Id: id,
	}
	b, _ = json.Marshal(br)
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/blog/get", bytes.NewReader(b))
	r.ServeHTTP(w, req)

	var getAns map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &getAns)
	code = getAns["Code"].(float64)
	message = getAns["Message"].(string)
	data := getAns["Data"].(map[string]interface{})
	title := data["title"].(string)
	content := data["content"].(string)

	assert.Equal(t, 0, int(code))
	assert.Equal(t, "success", message)
	assert.Equal(t, blogUpdateRequest.Title, title)
	assert.Equal(t, blogUpdateRequest.Content, content)
}

func Test_DeleteBlog(t *testing.T) {
	var blogDeleteRequest request.BlogDeleteRequest = request.BlogDeleteRequest{
		Id: id,
	}

	b, _ := json.Marshal(blogDeleteRequest)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/blog/delete", bytes.NewReader(b))
	req.AddCookie(&http.Cookie{
		Name:  "user_cookie",
		Value: "root",
	})
	r.ServeHTTP(w, req)

	var ans map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &ans)
	code := ans["Code"].(float64)
	message := ans["Message"].(string)

	assert.Equal(t, 0, int(code))
	assert.Equal(t, "success", message)
}
