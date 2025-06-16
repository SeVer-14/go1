package v1

//
//import (
//	"net/http"
//	"net/http/httptest"
//	"reflect"
//	"testing"
//
//	"github.com/gin-gonic/gin"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"go1/internal/entity"
//	"go1/internal/service"
//)
//
//type MockServices struct {
//	Product MockProductService
//	Cart    MockCartService
//	Order   MockOrderService
//}
//
//type MockProductService struct {
//	mock.Mock
//}
//
//func (m *MockProductService) GetAllProducts() ([]entity.Product, error) {
//	args := m.Called()
//	return args.Get(0).([]entity.Product), args.Error(1)
//}
//
////func (m *MockProductService) CreateProduct(product entity.Product) (entity.Product, error) {
////	args := m.Called(product)
////	return args.Get(0).(entity.Product), args.Error(1)
////}
////
////func (m *MockProductService) UpdateProduct(id uint, product entity.Product) (entity.Product, error) {
////	args := m.Called(id, product)
////	return args.Get(0).(entity.Product), args.Error(1)
////}
////
////func (m *MockProductService) DeleteProduct(id uint) error {
////	args := m.Called(id)
////	return args.Error(0)
////}
//
//type MockCartService struct {
//	mock.Mock
//}
//
////func (m *MockCartService) GetCart(userID uint) (entity.Cart, error) {
////	args := m.Called(userID)
////	return args.Get(0).(entity.Cart), args.Error(1)
////}
////
////func (m *MockCartService) AddToCart(userID, productID uint, quantity int) (entity.CartItem, error) {
////	args := m.Called(userID, productID, quantity)
////	return args.Get(0).(entity.CartItem), args.Error(1)
////}
////
////func (m *MockCartService) RemoveFromCart(userID, productID uint) error {
////	args := m.Called(userID, productID)
////	return args.Error(0)
////}
//
//type MockOrderService struct {
//	mock.Mock
//}
//
////func (m *MockOrderService) CreateOrder(cart entity.Cart) (entity.Order, error) {
////	args := m.Called(cart)
////	return args.Get(0).(entity.Order), args.Error(1)
////}
////
////func (m *MockOrderService) GetOrders(userID uint) ([]entity.Order, error) {
////	args := m.Called(userID)
////	return args.Get(0).([]entity.Order), args.Error(1)
////}
////
////func (m *MockOrderService) UpdateOrderStatus(orderID uint, status string) error {
////	args := m.Called(orderID, status)
////	return args.Error(0)
////}
//
//func setupRouter(services *MockServices) *gin.Engine {
//	gin.SetMode(gin.TestMode)
//	r := gin.Default()
//
//	svc := &service.Services{
//		Product: &service.ProductService{},
//		Cart:    &service.CartService{},
//		Order:   &service.OrderService{},
//	}
//
//	svcValue := reflect.ValueOf(svc).Elem()
//
//	productField := svcValue.FieldByName("Product")
//	if productField.IsValid() && productField.CanSet() {
//		productField.Set(reflect.ValueOf(&services.Product))
//	}
//
//	cartField := svcValue.FieldByName("Cart")
//	if cartField.IsValid() && cartField.CanSet() {
//		cartField.Set(reflect.ValueOf(&services.Cart))
//	}
//
//	orderField := svcValue.FieldByName("Order")
//	if orderField.IsValid() && orderField.CanSet() {
//		orderField.Set(reflect.ValueOf(&services.Order))
//	}
//
//	api := r.Group("/api")
//	NewHandler(svc).Init(api)
//
//	return r
//}
//
//func TestGetAllProducts(t *testing.T) {
//	t.Run("Success", func(t *testing.T) {
//		mockServices := &MockServices{
//			Product: MockProductService{},
//			Cart:    MockCartService{},
//			Order:   MockOrderService{},
//		}
//
//		mockServices.Product.On("GetAllProducts").Return([]entity.Product{
//			{ProductID: 1, Title: "Test", Price: 10.99},
//		}, nil)
//
//		router := setupRouter(mockServices)
//		w := httptest.NewRecorder()
//		req, _ := http.NewRequest("GET", "/api/products/", nil)
//		router.ServeHTTP(w, req)
//
//		assert.Equal(t, http.StatusOK, w.Code)
//		mockServices.Product.AssertExpectations(t)
//	})
//}
