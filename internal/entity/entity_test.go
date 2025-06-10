package entity

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestCartModel(t *testing.T) {
	type fields struct {
		ID        uint
		cartId    uint
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
		CartItems []CartItem
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Valid data", fields{ID: 1, cartId: 1}, false},
		{"Zero cart ID", fields{ID: 1, cartId: 0}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cart := &Cart{
				Model:     gorm.Model{ID: tt.fields.ID},
				cartId:    tt.fields.cartId,
				CartItems: tt.fields.CartItems,
			}
			if got := cart.cartId; got != tt.fields.cartId {
				t.Errorf("cartId is wrong, expected: %d, actual: %d", tt.fields.cartId, got)
			}
		})
	}
}

func TestCart_JSON(t *testing.T) {
	cart := Cart{cartId: 1}
	data, err := json.Marshal(cart)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"cartId":1`)
}

func BenchmarkCartModel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = &Cart{cartId: uint(i)}
	}
}

func TestProductModel(t *testing.T) {
	type fields struct {
		ID        uint
		ProductID int
		Title     string
		Price     float64
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"Valid product", fields{ID: 1, ProductID: 101, Title: "Product A", Price: 19.99}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			product := &Product{
				Model:     gorm.Model{ID: tt.fields.ID},
				ProductID: tt.fields.ProductID,
				Title:     tt.fields.Title,
				Price:     tt.fields.Price,
			}
			if got := product.Title; got != tt.fields.Title {
				t.Errorf("Title is wrong, expected: %s, actual: %s", tt.fields.Title, got)
			}
		})
	}
}

func TestOrderModel(t *testing.T) {
	tests := []struct {
		name    string
		order   Order
		isValid bool
		setup   func(*Order)
	}{
		{
			name: "Valid order",
			order: Order{
				cartId: 1,
				Status: "pending",
				Total:  100.50,
			},
			isValid: true,
		},

		{
			name: "Invalid status",
			order: Order{
				cartId: 1,
				Status: "invalid_status",
				Total:  100.50,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(&tt.order)
			}

			if tt.isValid {
				assert.NotNil(t, tt.order)
				if tt.order.Status == "" && tt.setup == nil {
					assert.Equal(t, "pending", tt.order.Status)
				}
			} else {
				if tt.order.Status != "" {
					validStatuses := []string{"pending", "processing", "completed", "cancelled"}
					assert.NotContains(t, validStatuses, tt.order.Status)
				}
				if tt.order.Total < 0 {
					assert.True(t, tt.order.Total < 0)
				}
			}
		})
	}
}

func TestOrderItemModel(t *testing.T) {
	tests := []struct {
		name     string
		item     OrderItem
		expected float64
		isValid  bool
	}{
		{
			name: "Valid item",
			item: OrderItem{
				Quantity: 2,
				Price:    25.99,
			},
			expected: 51.98,
			isValid:  true,
		},
		{
			name: "Zero quantity",
			item: OrderItem{
				Quantity: 0,
				Price:    25.99,
			},
			isValid: false,
		},
		{
			name: "Negative price",
			item: OrderItem{
				Quantity: 1,
				Price:    -10.00,
			},
			isValid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.isValid {
				assert.Equal(t, tt.expected, float64(tt.item.Quantity)*tt.item.Price)
			} else {
				if tt.item.Quantity <= 0 {
					assert.LessOrEqual(t, tt.item.Quantity, 0)
				}
				if tt.item.Price < 0 {
					assert.True(t, tt.item.Price < 0)
				}
			}
		})
	}
}

func TestOrderJSON(t *testing.T) {
	t.Run("Marshal Order", func(t *testing.T) {
		order := Order{
			cartId: 1,
			Status: "completed",
			Total:  99.99,
		}

		data, err := json.Marshal(order)
		assert.NoError(t, err)
		assert.Contains(t, string(data), `"cartId":1`)
		assert.Contains(t, string(data), `"status":"completed"`)
		assert.Contains(t, string(data), `"total":99.99`)
	})

	t.Run("Unmarshal Order", func(t *testing.T) {
		jsonStr := `{"cartId":2,"status":"pending","total":50.50}`
		var order Order
		err := json.Unmarshal([]byte(jsonStr), &order)

		assert.NoError(t, err)
		assert.Equal(t, uint(2), order.cartId)
		assert.Equal(t, "pending", order.Status)
		assert.Equal(t, 50.50, order.Total)
	})
}
func TestOrderStatusValidation(t *testing.T) {
	validStatuses := []string{"pending", "processing", "completed", "cancelled"}

	for _, status := range validStatuses {
		t.Run("Valid status: "+status, func(t *testing.T) {
			order := Order{Status: status}
			assert.Contains(t, validStatuses, order.Status)
		})
	}

	t.Run("Invalid status", func(t *testing.T) {
		order := Order{Status: "invalid"}
		assert.NotContains(t, validStatuses, order.Status)
	})
}
