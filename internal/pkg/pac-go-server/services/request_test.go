package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PDeXchange/pac/internal/pkg/pac-go-server/models"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAllRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, mockDBClient, tearDown := setUp(t)
	defer tearDown()

	testcases := []struct {
		name           string
		mockFunc       func()
		requestContext testContext
		httpStatus     int
	}{
		{
			name: "get all requests successfully",
			mockFunc: func() {
				mockDBClient.EXPECT().GetRequestsByUserID(gomock.Any(), gomock.Any()).Return(getResource("get-requests-by-user-id", nil).([]models.Request), nil).Times(1)
			},
			requestContext: formContext(customValues{
				"userid": "12345",
			}),
			httpStatus: http.StatusOK,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			req, err := http.NewRequest(http.MethodGet, "/requests", nil)
			if err != nil {
				t.Fatal(err)
			}
			ctx := getContext(tc.requestContext)
			c.Request = req.WithContext(ctx)
			dbCon = mockDBClient
			GetAllRequests(c)
			assert.Equal(t, tc.httpStatus, c.Writer.Status())
		})
	}
}

func TestGetRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	_, mockDBClient, tearDown := setUp(t)
	defer tearDown()

	testcases := []struct {
		name           string
		mockFunc       func()
		requestContext testContext
		httpStatus     int
	}{
		{
			name: "get request successfully",
			mockFunc: func() {
				mockDBClient.EXPECT().GetRequestByID(gomock.Any()).Return(getResource("get-request-by-id", nil).(*models.Request), nil).Times(1)
			},
			httpStatus: http.StatusOK,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.mockFunc()
			c, _ := gin.CreateTestContext(httptest.NewRecorder())
			req, err := http.NewRequest(http.MethodGet, "/requests", nil)
			if err != nil {
				t.Fatal(err)
			}
			ctx := getContext(tc.requestContext)
			c.Request = req.WithContext(ctx)
			dbCon = mockDBClient
			GetRequest(c)
			assert.Equal(t, tc.httpStatus, c.Writer.Status())
		})
	}
}
