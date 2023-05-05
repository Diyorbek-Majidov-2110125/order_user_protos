package models

type Promocode struct {
	Id              int     `json:"promocode_id"`
	Name            string  `json:"name"`
	Discount        int     `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type CreatePromoCode struct {
	Name            string  `json:"name"`
	Discount        int     `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type PromocodePrimaryKey struct {
	PromocodeId int `json:"promocode_id"`
}

type GetListPromocodeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPromocodeResponse struct {
	Count      int          `json:"count"`
	Promocodes []*Promocode `json:"promocode"`
}

type Date struct {
	Day string `json:"day"`
}

type Staffs struct {
	StaffName string  `json:"employe"`
	Category  string  `json:"category"`
	Product   string  `json:"product"`
	Quantity  int     `json:"quantity"`
	Summ      float32 `json:"summ"`
}

type Id struct {
	Order_id   string `json:"order_id"`
	Promo_Code string `json:"promo_code"`
}

type Disc struct {
	Order_id   int     `json:"order_id"`
	List_price float64 `json:"list_price"`
	Discount   float64 `json:"discount"`
}