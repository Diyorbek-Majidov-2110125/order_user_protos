
package postgresql

import (
	"app/api/models"
	"fmt"
	"context"
	"time"
	"github.com/jackc/pgx/v4/pgxpool"
)

type promoRepo struct {
	db *pgxpool.Pool
}

func NewPromoRepo(db *pgxpool.Pool) *promoRepo {
	return &promoRepo{
		db: db,
	}
}

func (r promoRepo) Create(ctx context.Context, req *models.CreatePromoCode) (int, error) {

	query := `INSERT INTO promocode("id","name",
		"discount",
		"discount_type",
		"order_limit_price"
		) 
		VALUES((SELECT MAX(id) + 1 FROM promocode),$1,$2,$3,$4) RETURNING id`
	id := 0
	err := r.db.QueryRow(ctx, query, req.Name, req.Discount, req.DiscountType, req.OrderLimitPrice).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r promoRepo) GetByID(ctx context.Context, req *models.PromocodePrimaryKey) (*models.Promocode, error) {
	var (
		query     string
		promocode models.Promocode
	)

	query = `SELECT * FROM promocode WHERE id =$1`

	err := r.db.QueryRow(ctx, query, req.PromocodeId).Scan(
		&promocode.Id, &promocode.Name, &promocode.Discount, &promocode.DiscountType, &promocode.OrderLimitPrice)
	if err != nil {
		return nil, err
	}

	return &promocode, nil
}

func (r promoRepo) GetList(ctx context.Context, req *models.GetListBrandRequest) (resp *models.GetListPromocodeResponse, err error) {

	resp = &models.GetListPromocodeResponse{}

	var (
		query  string
		filter = " WHERE TRUE "
		offset = " OFFSET 0"
		limit  = " LIMIT 10"
	)

	query = `
		SELECT
			COUNT(*) OVER(),
			id, 
			name, 
			discount,
			discount_type,
			order_limit_price
		FROM promocode
	`

	if len(req.Search) > 0 {
		filter += " AND name ILIKE '%' || '" + req.Search + "' || '%' "
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	query += filter + offset + limit

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var promocode models.Promocode
		err = rows.Scan(
			&resp.Count,
			&promocode.Id,
			&promocode.Name,
			&promocode.Discount,
			&promocode.DiscountType,
			&promocode.OrderLimitPrice,
		)
		if err != nil {
			return nil, err
		}

		resp.Promocodes = append(resp.Promocodes, &promocode)
	}

	return resp, nil
}

func (r promoRepo) Delete(ctx context.Context, req *models.PromocodePrimaryKey) (int64, error) {
	query := `
		DELETE 
		FROM promocode
		WHERE id = $1
	`

	result, err := r.db.Exec(ctx, query, req.PromocodeId)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}


func (r promoRepo) EveryStaff(ctx context.Context, req *models.Date) (res []models.Staffs, err error) {
	query := `SELECT
    staffs.first_name || ' ' || staffs.last_name AS "employe",  categories.category_name AS "category",
       products.product_name AS "product",   order_items.quantity AS "quantity",   order_items.list_price * order_items.quantity AS "summ"
FROM orders
         JOIN order_items ON orders.order_id = order_items.order_id
         JOIN products ON order_items.product_id = products.product_id
         JOIN categories ON products.category_id = categories.category_id
         JOIN staffs ON orders.staff_id = staffs.staff_id
WHERE orders.order_date = $1`

	var year string

	if req.Day == "" {
		dt := time.Now()
		year = dt.Format("2023-05-05")
	} else {
		year = req.Day
	}

	date, error := time.Parse("2023-05-05", year)
	if error != nil {
		fmt.Println(error)
		return
	}

	rows, err := r.db.Query(ctx, query, date)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var s models.Staffs
		err = rows.Scan(
			&s.StaffName,
			&s.Category,
			&s.Product,
			&s.Quantity,
			&s.Summ,
		)
		res = append(res, s)
		if err != nil {
			return res, err
		}
	}
	return res, nil

}


func (r promoRepo) Summ(ctx context.Context, req *models.Id) (res models.Disc, err error) {

	query := `select order_id, sum(list_price) AS "list_price" , sum(discount) AS "discount"
from order_items
WHERE order_id = $1 GROUP BY  order_id`

	err = r.db.QueryRow(ctx, query, req.Order_id).Scan(
		&res.Order_id,
		&res.List_price,
		&res.Discount,
	)

	if err != nil {
		return res, err
	}

	if req.Promo_Code == "" {
		return res, nil
	}

	res.List_price -= res.Discount

	return res, nil
}
