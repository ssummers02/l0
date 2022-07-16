package dto

import (
	"l0/internal/delivery/api/restmodel"
	"l0/internal/domain/entity"
	"l0/internal/infrastructure/dbmodel"
)

func OrderFromDB(d dbmodel.Order) entity.Order {
	return entity.Order{
		OrderUID:          d.OrderUID,
		TrackNumber:       d.TrackNumber,
		Entry:             d.Entry,
		Delivery:          DeliveryFromDB(d.Delivery),
		Payment:           PaymentFromDB(d.Payment),
		Items:             ItemsFromDB(d.Items),
		Locale:            d.Locale,
		InternalSignature: d.InternalSignature,
		CustomerID:        d.CustomerID,
		DeliveryService:   d.DeliveryService,
		Shardkey:          d.Shardkey,
		SmID:              d.SmID,
		DateCreated:       d.DateCreated,
		OofShard:          d.OofShard,
	}
}
func OrdersFromDB(d []dbmodel.Order) []entity.Order {
	t := make([]entity.Order, len(d))
	for i, v := range d {
		t[i] = OrderFromDB(v)
	}

	return t
}

func ItemFromDB(d dbmodel.Item) entity.Item {
	return entity.Item{
		OrderUID:    d.OrderUID,
		ChrtID:      d.ChrtID,
		TrackNumber: d.TrackNumber,
		Price:       d.Price,
		Rid:         d.Rid,
		Name:        d.Name,
		Sale:        d.Sale,
		Size:        d.Size,
		TotalPrice:  d.TotalPrice,
		NmID:        d.NmID,
		Brand:       d.Brand,
		Status:      d.Status,
	}
}

func ItemsFromDB(d []dbmodel.Item) []entity.Item {
	t := make([]entity.Item, len(d))
	for i, v := range d {
		t[i] = ItemFromDB(v)
	}

	return t
}
func DeliveryFromDB(d dbmodel.Delivery) entity.Delivery {
	return entity.Delivery{
		OrderUID: d.OrderUID,
		Name:     d.Name,
		Phone:    d.Phone,
		Zip:      d.Zip,
		City:     d.City,
		Address:  d.Address,
		Region:   d.Region,
		Email:    d.Email,
	}
}
func PaymentFromDB(d dbmodel.Payment) entity.Payment {
	return entity.Payment{
		OrderUID:     d.OrderUID,
		Transaction:  d.Transaction,
		RequestID:    d.RequestID,
		Currency:     d.Currency,
		Provider:     d.Provider,
		Amount:       d.Amount,
		PaymentDt:    d.PaymentDt,
		Bank:         d.Bank,
		DeliveryCost: d.DeliveryCost,
		GoodsTotal:   d.GoodsTotal,
		CustomFee:    d.CustomFee,
	}
}

func OrderToRest(e entity.Order) restmodel.Order {
	return restmodel.Order{
		OrderUID:          e.OrderUID,
		TrackNumber:       e.TrackNumber,
		Entry:             e.Entry,
		Delivery:          DeliveryToRest(e.Delivery),
		Payment:           PaymentToRest(e.Payment),
		Items:             ItemsToRest(e.Items),
		Locale:            e.Locale,
		InternalSignature: e.InternalSignature,
		CustomerID:        e.CustomerID,
		DeliveryService:   e.DeliveryService,
		Shardkey:          e.Shardkey,
		SmID:              e.SmID,
		DateCreated:       e.DateCreated,
		OofShard:          e.OofShard,
	}
}
func DeliveryToRest(e entity.Delivery) restmodel.Delivery {
	return restmodel.Delivery{
		Name:    e.Name,
		Phone:   e.Phone,
		Zip:     e.Zip,
		City:    e.City,
		Address: e.Address,
		Region:  e.Region,
		Email:   e.Email,
	}
}
func PaymentToRest(e entity.Payment) restmodel.Payment {
	return restmodel.Payment{
		Transaction:  e.Transaction,
		RequestID:    e.RequestID,
		Currency:     e.Currency,
		Provider:     e.Provider,
		Amount:       e.Amount,
		PaymentDt:    e.PaymentDt,
		Bank:         e.Bank,
		DeliveryCost: e.DeliveryCost,
		GoodsTotal:   e.GoodsTotal,
		CustomFee:    e.CustomFee,
	}
}
func ItemToRest(e entity.Item) restmodel.Item {
	return restmodel.Item{
		ChrtID:      e.ChrtID,
		TrackNumber: e.TrackNumber,
		Price:       e.Price,
		Rid:         e.Rid,
		Name:        e.Name,
		Sale:        e.Sale,
		Size:        e.Size,
		TotalPrice:  e.TotalPrice,
		NmID:        e.NmID,
		Brand:       e.Brand,
		Status:      e.Status,
	}
}

func ItemsToRest(e []entity.Item) []restmodel.Item {
	t := make([]restmodel.Item, len(e))
	for i, v := range e {
		t[i] = ItemToRest(v)
	}

	return t
}

func OrderToDB(e entity.Order) dbmodel.Order {
	return dbmodel.Order{
		OrderUID:          e.OrderUID,
		TrackNumber:       e.TrackNumber,
		Entry:             e.Entry,
		Delivery:          DeliveryToDB(e.Delivery),
		Payment:           PaymentToDB(e.Payment),
		Items:             ItemsToDB(e.Items),
		Locale:            e.Locale,
		InternalSignature: e.InternalSignature,
		CustomerID:        e.CustomerID,
		DeliveryService:   e.DeliveryService,
		Shardkey:          e.Shardkey,
		SmID:              e.SmID,
		DateCreated:       e.DateCreated,
		OofShard:          e.OofShard,
	}
}

func DeliveryToDB(e entity.Delivery) dbmodel.Delivery {
	return dbmodel.Delivery{
		OrderUID: e.OrderUID,
		Name:     e.Name,
		Phone:    e.Phone,
		Zip:      e.Zip,
		City:     e.City,
		Address:  e.Address,
		Region:   e.Region,
		Email:    e.Email,
	}
}

func PaymentToDB(e entity.Payment) dbmodel.Payment {
	return dbmodel.Payment{
		OrderUID:     e.OrderUID,
		Transaction:  e.Transaction,
		RequestID:    e.RequestID,
		Currency:     e.Currency,
		Provider:     e.Provider,
		Amount:       e.Amount,
		PaymentDt:    e.PaymentDt,
		Bank:         e.Bank,
		DeliveryCost: e.DeliveryCost,
		GoodsTotal:   e.GoodsTotal,
		CustomFee:    e.CustomFee,
	}
}
func ItemToDB(e entity.Item) dbmodel.Item {
	return dbmodel.Item{
		OrderUID:    e.OrderUID,
		ChrtID:      e.ChrtID,
		TrackNumber: e.TrackNumber,
		Price:       e.Price,
		Rid:         e.Rid,
		Name:        e.Name,
		Sale:        e.Sale,
		Size:        e.Size,
		TotalPrice:  e.TotalPrice,
		NmID:        e.NmID,
		Brand:       e.Brand,
		Status:      e.Status,
	}
}

func ItemsToDB(e []entity.Item) []dbmodel.Item {
	t := make([]dbmodel.Item, len(e))
	for i, v := range e {
		t[i] = ItemToDB(v)
	}

	return t
}

func OrderFromRest(d restmodel.Order) entity.Order {
	return entity.Order{
		OrderUID:          d.OrderUID,
		TrackNumber:       d.TrackNumber,
		Entry:             d.Entry,
		Delivery:          DeliveryFromRest(d.Delivery, d.OrderUID),
		Payment:           PaymentFromRest(d.Payment, d.OrderUID),
		Items:             ItemsFromRest(d.Items, d.OrderUID),
		Locale:            d.Locale,
		InternalSignature: d.InternalSignature,
		CustomerID:        d.CustomerID,
		DeliveryService:   d.DeliveryService,
		Shardkey:          d.Shardkey,
		SmID:              d.SmID,
		DateCreated:       d.DateCreated,
		OofShard:          d.OofShard,
	}
}

func ItemFromRest(d restmodel.Item, uid string) entity.Item {
	return entity.Item{
		OrderUID:    uid,
		ChrtID:      d.ChrtID,
		TrackNumber: d.TrackNumber,
		Price:       d.Price,
		Rid:         d.Rid,
		Name:        d.Name,
		Sale:        d.Sale,
		Size:        d.Size,
		TotalPrice:  d.TotalPrice,
		NmID:        d.NmID,
		Brand:       d.Brand,
		Status:      d.Status,
	}
}

func ItemsFromRest(d []restmodel.Item, uid string) []entity.Item {
	t := make([]entity.Item, len(d))
	for i, v := range d {
		t[i] = ItemFromRest(v, uid)
	}

	return t
}
func DeliveryFromRest(d restmodel.Delivery, uid string) entity.Delivery {
	return entity.Delivery{
		OrderUID: uid,
		Name:     d.Name,
		Phone:    d.Phone,
		Zip:      d.Zip,
		City:     d.City,
		Address:  d.Address,
		Region:   d.Region,
		Email:    d.Email,
	}
}
func PaymentFromRest(d restmodel.Payment, uid string) entity.Payment {
	return entity.Payment{
		OrderUID:     uid,
		Transaction:  d.Transaction,
		RequestID:    d.RequestID,
		Currency:     d.Currency,
		Provider:     d.Provider,
		Amount:       d.Amount,
		PaymentDt:    d.PaymentDt,
		Bank:         d.Bank,
		DeliveryCost: d.DeliveryCost,
		GoodsTotal:   d.GoodsTotal,
		CustomFee:    d.CustomFee,
	}
}
