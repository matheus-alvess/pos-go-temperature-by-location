package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"pos-go-temperature-by-location/internal/core/domain"
	"pos-go-temperature-by-location/internal/ports/mock"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	gomock "go.uber.org/mock/gomock"
)

func TestGetWeather(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name                 string
		cep                  string
		mockReturn           *domain.Temperature
		mockError            error
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Valid CEP",
			cep:  "12345678",
			mockReturn: &domain.Temperature{
				TempC: 25.5,
				TempF: 50.6,
				TempK: 77.9,
			},
			mockError:            nil,
			expectedStatusCode:   http.StatusOK,
			expectedResponseBody: `{"temp_C":25.5,"temp_F":50.6,"temp_K":77.9}`,
		},
		{
			name:                 "Invalid CEP format",
			cep:                  "123",
			expectedStatusCode:   http.StatusUnprocessableEntity,
			expectedResponseBody: `{"message":"invalid zipcode"}`,
		},
		{
			name:                 "CEP not found",
			cep:                  "87654321",
			mockError:            errors.New("CEP not found"),
			expectedStatusCode:   http.StatusNotFound,
			expectedResponseBody: `{"message":"can not find zipcode"}`,
		},
		{
			name:                 "Internal server error",
			cep:                  "12345678",
			mockError:            errors.New("some internal error"),
			expectedStatusCode:   http.StatusInternalServerError,
			expectedResponseBody: `{"message":"failed to get temperature"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockService := mocks.NewMockWeatherService(ctrl)
			handler := NewWeatherHandler(mockService)

			if tt.mockError != nil || tt.mockReturn != nil {
				mockService.EXPECT().GetWeather(tt.cep).Return(tt.mockReturn, tt.mockError)
			}

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = gin.Params{gin.Param{Key: "cep", Value: tt.cep}}

			handler.GetWeather(c)

			assert.Equal(t, tt.expectedStatusCode, w.Code)
			assert.JSONEq(t, tt.expectedResponseBody, w.Body.String())
		})
	}
}
