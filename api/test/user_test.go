package test

import (
	"github.com/Ghun2/fast-gorestapi/router/middleware"
	"github.com/Ghun2/fast-gorestapi/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSignUpCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"user_name":"alice","email":"alice@ghun2ee.com","auth_id":"test:0123"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, ctrl.SignUp(c))
	if assert.Equal(t, http.StatusCreated, rec.Code) {
		m := responseMap(rec.Body.Bytes(), "user")
		assert.Equal(t, "alice", m["user_name"])
		assert.Equal(t, "alice@ghun2ee.com", m["email"])
		assert.Empty(t, m["birth"])
		assert.Empty(t, m["sex"])
		assert.Empty(t, m["phone"])
		assert.NotEmpty(t, rec.Header().Get("Authorization"))
	}
}

func TestLoginCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	var (
		reqJSON = `{"auth_id": "kakao:1234"}`
	)
	req := httptest.NewRequest(echo.POST, "/api/users/login", strings.NewReader(reqJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	assert.NoError(t, ctrl.Login(c))
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes(), "user")
		assert.Equal(t, "user1", m["user_name"])
		assert.Equal(t, "user1@ghun2ee.com", m["email"])
		assert.NotEmpty(t, rec.Header().Get("Authorization"))
	}
}

func TestCurrentUserCaseSuccess(t *testing.T) {
	tearDown()
	setup()
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.GET, "/api/users/login", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(1))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return ctrl.CurrentUser(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes(), "user")
		assert.Equal(t, "user1", m["user_name"])
		assert.Equal(t, "user1@ghun2ee.com", m["email"])
	}
}

func TestUpdateUserUserName(t *testing.T) {
	tearDown()
	setup()
	user1UpdateReq := `{"user_name": "ghun"}`
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.PUT, "/api/user", strings.NewReader(user1UpdateReq))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(1))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return ctrl.UpdateUser(c)
	})(c)
	assert.NoError(t, err)
	if assert.Equal(t, http.StatusOK, rec.Code) {
		m := responseMap(rec.Body.Bytes(), "user")
		assert.Equal(t, "user1@ghun2ee.com", m["email"])
		assert.Equal(t, "ghun", m["user_name"])
	}
}

func TestDeleteUser(t *testing.T) {
	tearDown()
	setup()
	usrlst, _ := us.ListUsers()
	startCnt := len(usrlst)
	jwtMiddleware := middleware.JWT(utils.JWTSecret)
	req := httptest.NewRequest(echo.DELETE, "/api/user", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderAuthorization, utils.GenerateJWT(1))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	err := jwtMiddleware(func(context echo.Context) error {
		return ctrl.DeleteUser(c)
	})(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	usrlst, _ = us.ListUsers()
	endCnt := len(usrlst)
	assert.Equal(t, startCnt-1, endCnt)
}
