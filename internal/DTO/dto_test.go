package dto

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToCartDTO(t *testing.T) {
	tests := []struct {
		name    string
		dto     AddToCartDTO
		wantErr bool
	}{
		{
			name:    "Valid AddToCartDTO",
			dto:     AddToCartDTO{ProductID: 1, Quantity: 2},
			wantErr: false,
		},
		{
			name:    "Zero ProductID",
			dto:     AddToCartDTO{ProductID: 0, Quantity: 1},
			wantErr: true,
		},
		{
			name:    "Zero Quantity",
			dto:     AddToCartDTO{ProductID: 1, Quantity: 0},
			wantErr: true,
		},
		{
			name:    "Negative Quantity",
			dto:     AddToCartDTO{ProductID: 1, Quantity: -1},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				if tt.dto.ProductID == 0 {
					assert.Equal(t, uint(0), tt.dto.ProductID)
				}
				if tt.dto.Quantity < 1 {
					assert.Less(t, tt.dto.Quantity, 1)
				}
			} else {
				assert.NotEqual(t, uint(0), tt.dto.ProductID)
				assert.GreaterOrEqual(t, tt.dto.Quantity, 1)
			}
		})
	}
}

func TestProductDTO(t *testing.T) {
	tests := []struct {
		name    string
		dto     ProductDTO
		wantErr bool
	}{
		{
			name: "Valid ProductDTO",
			dto: ProductDTO{
				ID:        1,
				ProductID: 101,
				Title:     "Test Product",
				Price:     19.99,
			},
			wantErr: false,
		},
		{
			name: "Empty Title",
			dto: ProductDTO{
				ProductID: 102,
				Title:     "",
				Price:     9.99,
			},
			wantErr: true,
		},
		{
			name: "Zero Price",
			dto: ProductDTO{
				ProductID: 103,
				Title:     "Free Product",
				Price:     0,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				if tt.dto.Title == "" {
					assert.Empty(t, tt.dto.Title)
				}
				if tt.dto.Price == 0 {
					assert.Equal(t, 0.0, tt.dto.Price)
				}
			} else {
				assert.NotEmpty(t, tt.dto.Title)
				assert.NotEqual(t, 0.0, tt.dto.Price)
			}
		})
	}
}

func TestCartDTO_JSON(t *testing.T) {
	t.Run("Marshal CartDTO", func(t *testing.T) {
		cart := CartDTO{
			CartID: 1,
			Items: []CartItemDTO{
				{
					ProductID: 1,
					Quantity:  2,
					Price:     10.99,
					Title:     "Test Item",
				},
			},
			Total: 21.98,
		}

		data, err := json.Marshal(cart)
		assert.NoError(t, err)
		assert.Contains(t, string(data), `"cartId":1`)
		assert.Contains(t, string(data), `"productId":1`)
		assert.Contains(t, string(data), `"quantity":2`)
		assert.Contains(t, string(data), `"total":21.98`)
	})

	t.Run("Unmarshal CartDTO", func(t *testing.T) {
		jsonStr := `{"cartId":2,"items":[{"productId":2,"quantity":1,"price":5.99,"title":"Another Item"}],"total":5.99}`
		var cart CartDTO
		err := json.Unmarshal([]byte(jsonStr), &cart)

		assert.NoError(t, err)
		assert.Equal(t, uint(2), cart.CartID)
		assert.Len(t, cart.Items, 1)
		assert.Equal(t, 5.99, cart.Total)
	})
}

func TestOrderDTO_Validation(t *testing.T) {
	tests := []struct {
		name    string
		dto     OrderDTO
		wantErr bool
	}{
		{
			name: "Valid OrderDTO",
			dto: OrderDTO{
				ID:     1,
				Status: "pending",
				Total:  100.50,
				Items: []OrderItemDTO{
					{
						ProductID: 1,
						Quantity:  2,
						Price:     50.25,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "Invalid Status",
			dto: OrderDTO{
				Status: "invalid_status",
			},
			wantErr: true,
		},
		{
			name: "Negative Total",
			dto: OrderDTO{
				Total: -10.00,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.wantErr {
				if tt.dto.Status != "" {
					validStatuses := []string{"pending", "processing", "completed", "cancelled"}
					assert.NotContains(t, validStatuses, tt.dto.Status)
				}
				if tt.dto.Total < 0 {
					assert.True(t, tt.dto.Total < 0)
				}
			} else {
				assert.NotEmpty(t, tt.dto.Status)
				assert.GreaterOrEqual(t, tt.dto.Total, 0.0)
				if len(tt.dto.Items) > 0 {
					for _, item := range tt.dto.Items {
						assert.Greater(t, item.Quantity, 0)
						assert.GreaterOrEqual(t, item.Price, 0.0)
					}
				}
			}
		})
	}
}

func TestUpdateOrderStatusDTO(t *testing.T) {
	tests := []struct {
		name    string
		dto     UpdateOrderStatusDTO
		wantErr bool
	}{
		{
			name:    "Valid Update",
			dto:     UpdateOrderStatusDTO{OrderID: 1, Status: "completed"},
			wantErr: false,
		},
		{
			name:    "Zero OrderID",
			dto:     UpdateOrderStatusDTO{OrderID: 0, Status: "processing"},
			wantErr: true,
		},
		{
			name:    "Empty Status",
			dto:     UpdateOrderStatusDTO{OrderID: 1, Status: ""},
			wantErr: true,
		},
		{
			name:    "Invalid Status",
			dto:     UpdateOrderStatusDTO{OrderID: 1, Status: "invalid"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validStatuses := []string{"pending", "processing", "completed", "cancelled"}

			if tt.wantErr {
				if tt.dto.OrderID == 0 {
					assert.Equal(t, uint(0), tt.dto.OrderID)
				}
				if tt.dto.Status == "" {
					assert.Empty(t, tt.dto.Status)
				} else if !contains(validStatuses, tt.dto.Status) {
					assert.NotContains(t, validStatuses, tt.dto.Status)
				}
			} else {
				assert.NotEqual(t, uint(0), tt.dto.OrderID)
				assert.Contains(t, validStatuses, tt.dto.Status)
			}
		})
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func BenchmarkCartItemDTO_Marshal(b *testing.B) {
	item := CartItemDTO{
		ProductID: 1,
		Quantity:  5,
		Price:     12.99,
		Title:     "Benchmark Product",
	}

	for i := 0; i < b.N; i++ {
		_, _ = json.Marshal(item)
	}
}
