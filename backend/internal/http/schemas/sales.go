package schemas

// Upload contract — ONLY headers required and allowed fields
type SalesImportSchema struct {
	OrderID        string `csv:"order_id" validate:"required"`
	OrderDate      string `csv:"order_date" validate:"required"`
	ProductID      string `csv:"product_id" validate:"required"`
	Quantity       string `csv:"quantity" validate:"required"`
	Channel        string `csv:"channel,omitempty"`
	Marketplace    string `csv:"marketplace,omitempty"`
	OrderStatus    string `csv:"order_status,omitempty"`
	PaymentMethod  string `csv:"payment_method,omitempty"`
	ProductName    string `csv:"product_name,omitempty"`
	Category       string `csv:"category,omitempty"`
	Brand          string `csv:"brand,omitempty"`
	Variant        string `csv:"variant,omitempty"`
	UnitPrice      string `csv:"unit_price,omitempty"`
	Revenue        string `csv:"revenue,omitempty"`
	Discount       string `csv:"discount,omitempty"`
	Promotion      string `csv:"promotion,omitempty"`
	CustomerRegion string `csv:"customer_region,omitempty"`
	CustomerState  string `csv:"customer_state,omitempty"`
	IsNewCustomer  string `csv:"is_new_customer,omitempty"`
	ShippingMethod string `csv:"shipping_method,omitempty"`
	ShippingDays   string `csv:"shipping_days,omitempty"`
	ShippingCost   string `csv:"shipping_cost,omitempty"`
	DeliveryDate   string `csv:"delivery_date,omitempty"`
	Campaign       string `csv:"campaign,omitempty"`
	TrafficSource  string `csv:"traffic_source,omitempty"`
	Device         string `csv:"device,omitempty"`
}
