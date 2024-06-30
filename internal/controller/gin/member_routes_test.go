package controller

import (
	"chat-app/domain"
	"chat-app/dto"
	"chat-app/internal/configuration"
	"chat-app/internal/logging"
	repository "chat-app/internal/repository/postgres"
	"chat-app/internal/service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/smartystreets/goconvey/convey"
	"gorm.io/gorm"
)

var testServer *gin.Engine
var testRepository *gorm.DB

func TestMain(m *testing.M) {
	testSetup()
	code := m.Run()
	testTeardown()
	os.Exit(code)
}

func testSetup() {
	baseDir, err := findBaseDir()
	if err != nil {
		panic(err)
	}

	if err := configuration.InitializeWithSpecifiedWorkingDirectory(baseDir, "test"); err != nil {
		logging.Log.Error(err)
		panic(err)
	}

	if err := repository.InitializeConnection(); err != nil {
		logging.Log.Error(err)
		panic(err)
	}
	repository.InitializeRepositories(repository.DB)

	service.InitializeService(repository.DB, repository.MemberRepository)

	InitializeGinServer("")
	InitializeRoutes(Server, service.MemberService)
	testServer = Server
	testRepository = repository.DB
}

// findBaseDir searches for the base directory by looking for a known marker file
func findBaseDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Look for a known marker file in parent directories
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); !os.IsNotExist(err) {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	return "", os.ErrNotExist
}

func testTeardown() {
	testRepository.Exec("DELETE FROM ca_member")
	testRepository.Exec("DELETE FROM ca_transaction")
}

func TestCreateMember(t *testing.T) {
	convey.Convey("given a request body", t, func() {
		requestBody := dto.CreateMemberRequest{
			Name:     "test-username",
			Email:    "test@gmail.com",
			Gender:   "MALE",
			Password: "test123",
		}

		jsonReqBody, _ := json.Marshal(requestBody)
		request, _ := http.NewRequest(http.MethodPost, "/v1/member", strings.NewReader(string(jsonReqBody)))

		httpRec := httptest.NewRecorder()
		Server.ServeHTTP(httpRec, request)

		convey.Convey("Make sure there is a data inserted into the database", func() {
			var expectedMember domain.Member
			testRepository.Raw("SELECT * FROM ca_member WHERE email = 'test@gmail.com'").Find(&expectedMember)
			assert.Equal(t, http.StatusOK, httpRec.Code)
			assert.Equal(t, domain.Member{}, expectedMember)
		})
	})
}

func TestCreateMemberWithExistingMemberEmail(t *testing.T) {
	convey.Convey("given an already created member and a request body", t, func() {
		testRepository.Exec("INSERT INTO ca_member(Name, Email, Gender, Password) VALUES ('test-username', 'test@gmail.com', 'MALE', '12345')")

		convey.Convey("when the API is hit", func() {
			requestBody := dto.CreateMemberRequest{
				Name:     "test-username",
				Email:    "test@gmail.com",
				Gender:   "MALE",
				Password: "test123",
			}
			jsonReqBody, _ := json.Marshal(requestBody)
			request, _ := http.NewRequest(http.MethodPost, "/v1/member", strings.NewReader(string(jsonReqBody)))

			httpRec := httptest.NewRecorder()
			Server.ServeHTTP(httpRec, request)

			convey.Convey("Make sure there is a data inserted into the database", func() {
				var expectedMembers []domain.Member
				testRepository.Raw("SELECT * FROM ca_member WHERE email = 'test@gmail.com'").Find(&expectedMembers)
				assert.Equal(t, http.StatusInternalServerError, httpRec.Code)
				assert.Equal(t, 1, len(expectedMembers))
			})
		})
	})
}
