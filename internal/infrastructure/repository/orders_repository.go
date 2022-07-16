package repository

import (
	"context"
	"l0/internal/domain/dto"
	"l0/internal/domain/entity"
	"l0/internal/infrastructure/dbmodel"

	"github.com/gocraft/dbr"
)

type OrdersRepository struct {
	db    DBConn
	cache ordersCache
}

func NewOrdersRepository(db DBConn) *OrdersRepository {
	return &OrdersRepository{
		db:    db,
		cache: newOrdersCache(),
	}
}
func (r *OrdersRepository) InsertOrder(ctx context.Context, w entity.Order) error {
	err := r.insertOrder(ctx, w)
	if err != nil {
		return err
	}

	r.cache.setOrders(w)

	return err
}

func (r *OrdersRepository) GetOrdersByID(ctx context.Context, id string) (entity.Order, error) {
	if r.cache.len() != 0 {
		return r.cache.find(id)
	}

	o, err := r.orders(ctx)
	if err != nil {
		return entity.Order{}, err
	}

	r.cache.setOrders(o...)

	return r.cache.find(id)
}
func (r *OrdersRepository) orders(ctx context.Context) ([]entity.Order, error) {
	var orders []dbmodel.Order

	err := r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.
			Select("*").
			From("orders").
			LeftJoin("delivery", "orders.order_uid = delivery.order_uid").
			LeftJoin("items", "orders.order_uid = items.order_uid").
			LeftJoin("payment", "orders.order_uid = payment.order_uid").
			LoadContext(ctx, &orders)
		if err != nil {
			return err
		}

		return err
	})
	if err != nil {
		return []entity.Order{}, err
	}

	return dto.OrdersFromDB(orders), err
}

func (r *OrdersRepository) insertOrder(ctx context.Context, w entity.Order) error {
	order := dto.OrderToDB(w)

	return r.db.BeginTx(ctx, func(tx *dbr.Tx) error {
		_, err := tx.InsertInto("orders").
			Pair("order_uid", order.OrderUID).
			Pair("track_number", order.TrackNumber).
			Pair("entry", order.Entry).
			Pair("locale", order.Locale).
			Pair("internal_signature", order.InternalSignature).
			Pair("customer_id", order.CustomerID).
			Pair("delivery_service", order.DeliveryService).
			Pair("shardkey", order.Shardkey).
			Pair("sm_id", order.SmID).
			Pair("date_created", order.DateCreated).
			Pair("oof_shard", order.OofShard).
			Exec()
		if err != nil {
			return err
		}
		_, err = tx.InsertInto("delivery").
			Pair("order_uid", order.OrderUID).
			Pair("name", order.Delivery.Name).
			Pair("phone", order.Delivery.Phone).
			Pair("zip", order.Delivery.Zip).
			Pair("city", order.Delivery.City).
			Pair("address", order.Delivery.Address).
			Pair("region", order.Delivery.Region).
			Pair("email", order.Delivery.Email).
			Exec()
		if err != nil {
			return err
		}
		_, err = tx.InsertInto("payment").
			Pair("order_uid", order.OrderUID).
			Pair("transaction", order.Payment.Transaction).
			Pair("request_id", order.Payment.RequestID).
			Pair("currency", order.Payment.Currency).
			Pair("provider", order.Payment.Provider).
			Pair("amount", order.Payment.Amount).
			Pair("payment_dt", order.Payment.PaymentDt).
			Pair("bank", order.Payment.Bank).
			Pair("delivery_cost", order.Payment.DeliveryCost).
			Pair("goods_total", order.Payment.GoodsTotal).
			Pair("custom_fee", order.Payment.CustomFee).
			Exec()
		if err != nil {
			return err
		}
		for _, v := range order.Items {
			_, err = tx.InsertInto("items").
				Pair("order_uid", v.OrderUID).
				Pair("chrt_id", v.ChrtID).
				Pair("track_number", v.TrackNumber).
				Pair("price", v.Price).
				Pair("rid", v.Rid).
				Pair("name", v.Name).
				Pair("sale", v.Sale).
				Pair("size", v.Size).
				Pair("total_price", v.TotalPrice).
				Pair("nm_id", v.NmID).
				Pair("brand", v.Brand).
				Pair("status", v.Status).
				Exec()
			if err != nil {
				return err
			}
		}

		return err
	})
}
