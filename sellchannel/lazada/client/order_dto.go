package client

import "encoding/json"

// envelope is the common Lazada response wrapper. code == "0" means success.
type envelope struct {
	Code      string          `json:"code"`
	Message   string          `json:"message"`
	Type      string          `json:"type"`
	RequestID string          `json:"request_id"`
	Data      json.RawMessage `json:"data"`
}

// ── Public return types ─────────────────────────────────────────────────────

type OrderDetail struct {
	OrderID         int64
	OrderNumber     int64
	CreatedAt       string
	UpdatedAt       string
	PaymentMethod   string
	Statuses        []string
	AddressShipping OrderAddress
	AddressBilling  OrderAddress
}

type OrderAddress struct {
	FirstName string
	LastName  string
	Address1  string
	Address2  string
	Address3  string
	City      string
	PostCode  string
	Country   string
	Phone     string
	Phone2    string
}

type OrderItem struct {
	OrderID            int64
	OrderItemID        int64
	Sku                string
	ShopSku            string
	Name               string
	ItemPrice          float64
	PaidPrice          float64
	VoucherAmount      float64
	VoucherPlatformLpi float64
	VoucherSellerLpi   float64
	TaxAmount          float64
	ShippingAmount     float64
	ShippingType       string
	ShippingProvType   string
	ShipmentProvider   string
	TrackingCode       string
	PackageID          string
	Status             string
	OrderType          string
	Reason             string
	ReasonDetail       string
	SlaTimeStamp       string
}

// ── /orders/get ─────────────────────────────────────────────────────────────

// orderListDTO is the `data` block of /orders/get.
type orderListDTO struct {
	CountTotal int `json:"countTotal"`
	Orders     []struct {
		OrderNumber int64 `json:"order_number"`
	} `json:"orders"`
}

// ── /order/get ──────────────────────────────────────────────────────────────

type orderDetailDTO struct {
	OrderID       int64      `json:"order_id"`
	OrderNumber   int64      `json:"order_number"`
	CreatedAt     string     `json:"created_at"`
	UpdatedAt     string     `json:"updated_at"`
	PaymentMethod string     `json:"payment_method"`
	Statuses      []string   `json:"statuses"`
	AddressShip   addressDTO `json:"address_shipping"`
	AddressBill   addressDTO `json:"address_billing"`
}

type addressDTO struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Address1  string `json:"address1"`
	Address2  string `json:"address2"`
	Address3  string `json:"address3"`
	City      string `json:"city"`
	PostCode  string `json:"post_code"`
	Country   string `json:"country"`
	Phone     string `json:"phone"`
	Phone2    string `json:"phone2"`
}

func (d orderDetailDTO) toDomain() *OrderDetail {
	return &OrderDetail{
		OrderID:         d.OrderID,
		OrderNumber:     d.OrderNumber,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
		PaymentMethod:   d.PaymentMethod,
		Statuses:        d.Statuses,
		AddressShipping: d.AddressShip.toDomain(),
		AddressBilling:  d.AddressBill.toDomain(),
	}
}

func (a addressDTO) toDomain() OrderAddress {
	return OrderAddress{
		FirstName: a.FirstName,
		LastName:  a.LastName,
		Address1:  a.Address1,
		Address2:  a.Address2,
		Address3:  a.Address3,
		City:      a.City,
		PostCode:  a.PostCode,
		Country:   a.Country,
		Phone:     a.Phone,
		Phone2:    a.Phone2,
	}
}

// ── /order/items/get ────────────────────────────────────────────────────────

type orderItemDTO struct {
	OrderID            int64   `json:"order_id"`
	OrderItemID        int64   `json:"order_item_id"`
	Sku                string  `json:"sku"`
	ShopSku            string  `json:"shop_sku"`
	Name               string  `json:"name"`
	ItemPrice          float64 `json:"item_price"`
	PaidPrice          float64 `json:"paid_price"`
	VoucherAmount      float64 `json:"voucher_amount"`
	VoucherPlatformLpi float64 `json:"voucher_platform_lpi"`
	VoucherSellerLpi   float64 `json:"voucher_seller_lpi"`
	TaxAmount          float64 `json:"tax_amount"`
	ShippingAmount     float64 `json:"shipping_amount"`
	ShippingType       string  `json:"shipping_type"`
	ShippingProvType   string  `json:"shipping_provider_type"`
	ShipmentProvider   string  `json:"shipment_provider"`
	TrackingCode       string  `json:"tracking_code"`
	PackageID          string  `json:"package_id"`
	Status             string  `json:"status"`
	OrderType          string  `json:"order_type"`
	Reason             string  `json:"reason"`
	ReasonDetail       string  `json:"reason_detail"`
	SlaTimeStamp       string  `json:"sla_time_stamp"`
}

func (i orderItemDTO) toDomain() OrderItem {
	return OrderItem{
		OrderID:            i.OrderID,
		OrderItemID:        i.OrderItemID,
		Sku:                i.Sku,
		ShopSku:            i.ShopSku,
		Name:               i.Name,
		ItemPrice:          i.ItemPrice,
		PaidPrice:          i.PaidPrice,
		VoucherAmount:      i.VoucherAmount,
		VoucherPlatformLpi: i.VoucherPlatformLpi,
		VoucherSellerLpi:   i.VoucherSellerLpi,
		TaxAmount:          i.TaxAmount,
		ShippingAmount:     i.ShippingAmount,
		ShippingType:       i.ShippingType,
		ShippingProvType:   i.ShippingProvType,
		ShipmentProvider:   i.ShipmentProvider,
		TrackingCode:       i.TrackingCode,
		PackageID:          i.PackageID,
		Status:             i.Status,
		OrderType:          i.OrderType,
		Reason:             i.Reason,
		ReasonDetail:       i.ReasonDetail,
		SlaTimeStamp:       i.SlaTimeStamp,
	}
}
