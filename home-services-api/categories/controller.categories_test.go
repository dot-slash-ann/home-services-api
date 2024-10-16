package categories_test

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dot-slash-ann/home-services-api/categories"
	"github.com/dot-slash-ann/home-services-api/entities"
	"github.com/dot-slash-ann/home-services-api/lib/httpErrors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCategoriesService struct {
	mock.Mock
}

func (m *MockCategoriesService) Create(categoryDto categories.CreateCategoryDto) (entities.Category, error) {
	args := m.Called(categoryDto)

	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoriesService) FindAll() ([]entities.Category, error) {
	args := m.Called()

	return args.Get(0).([]entities.Category), args.Error(1)
}

func (m *MockCategoriesService) FindOne(id string) (entities.Category, error) {
	args := m.Called(id)

	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoriesService) FindByName(name string) (entities.Category, error) {
	args := m.Called(name)

	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoriesService) Update(id string, updateCategoryDto categories.UpdateCategoryDto) (entities.Category, error) {
	args := m.Called(id, updateCategoryDto)

	return args.Get(0).(entities.Category), args.Error(1)
}

func (m *MockCategoriesService) Delete(id string) (entities.Category, error) {
	args := m.Called(id)

	return args.Get(0).(entities.Category), args.Error(1)
}

func TestCategoriesControllerCreate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockCategoriesService)

	mockService.On("Create", mock.AnythingOfType("CreateCategoryDto")).Return(entities.Category{}, nil)

	controller := categories.NewCategoriesController(mockService)

	router := gin.Default()
	router.POST("/test_create", controller.Create)

	createRequestBody := `{"name":"mock category"}`

	req, err := http.NewRequest(http.MethodPost, "/test_create", bytes.NewBuffer([]byte(createRequestBody)))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusCreated, recorder.Code)

	mockService.AssertNumberOfCalls(t, "Create", 1)
	mockService.AssertCalled(t, "Create", categories.CreateCategoryDto{Name: "mock category"})
	mockService.AssertExpectations(t)
}

func TestCategoriesControllerFindAll(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockCategoriesService)

	mockService.On("FindAll").Return([]entities.Category{}, nil)

	controller := categories.NewCategoriesController(mockService)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)

	controller.FindAll(c)

	assert.Len(t, c.Errors, 0)
	assert.Equal(t, http.StatusOK, recorder.Code)

	mockService.AssertNumberOfCalls(t, "FindAll", 1)
	mockService.AssertCalled(t, "FindAll")
	mockService.AssertExpectations(t)
}

func TestCategoriesControllerFindOne(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockCategoriesService)

	mockService.On("FindOne", "1").Return(entities.Category{}, nil)

	controller := categories.NewCategoriesController(mockService)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	controller.FindOne(c)

	assert.Len(t, c.Errors, 0)
	assert.Equal(t, http.StatusOK, recorder.Code)

	mockService.AssertNumberOfCalls(t, "FindOne", 1)
	mockService.AssertCalled(t, "FindOne", "1")
	mockService.AssertExpectations(t)
}

func TestCategoriesControllerFindOneNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockCategoriesService)

	mockService.On("FindOne", "1").Return(entities.Category{}, errors.New("record not found"))

	controller := categories.NewCategoriesController(mockService)

	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	c.Params = append(c.Params, gin.Param{Key: "id", Value: "1"})

	controller.FindOne(c)

	errs := c.Errors

	assert.Len(t, errs, 1)
	assert.Equal(t, httpErrors.NotFoundError(errors.New("record not found"), nil).Error(), errs[0].Err.Error())

	mockService.AssertNumberOfCalls(t, "FindOne", 1)
	mockService.AssertCalled(t, "FindOne", "1")
	mockService.AssertExpectations(t)
}

func TestCategoriesControllerUpdate(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockCategoriesService)

	mockService.On("Update", "1", mock.AnythingOfType("UpdateCategoryDto")).Return(entities.Category{}, nil)

	controller := categories.NewCategoriesController(mockService)

	router := gin.Default()
	router.PATCH("/test_update/:id", controller.Update)

	updateRequestBody := `{"name":"new name"}`

	req, err := http.NewRequest(http.MethodPatch, "/test_update/1", bytes.NewBuffer([]byte(updateRequestBody)))
	req.Header.Set("Content-Type", "application/json")

	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	mockService.AssertNumberOfCalls(t, "Update", 1)
	mockService.AssertCalled(t, "Update", "1", categories.UpdateCategoryDto{Name: "new name"})
	mockService.AssertExpectations(t)
}

func TestCategoriesControllerDelete(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockCategoriesService)

	mockService.On("Delete", "1").Return(entities.Category{}, nil)

	controller := categories.NewCategoriesController(mockService)

	router := gin.Default()
	router.DELETE("/test_delete/:id", controller.Delete)

	req, err := http.NewRequest(http.MethodDelete, "/test_delete/1", nil)

	assert.NoError(t, err)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)

	mockService.AssertNumberOfCalls(t, "Delete", 1)
	mockService.AssertCalled(t, "Delete", "1")
	mockService.AssertExpectations(t)
}
