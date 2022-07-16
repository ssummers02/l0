package main

import (
	"l0/internal/delivery/api/restmodel"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

//nolint:gosec
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	return string(b)
}

//nolint
func main() {
	// Connect to a server
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		log.Println(err)
	}

	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	o := restmodel.Order{
		OrderUID:    RandStringBytes(10),
		TrackNumber: RandStringBytes(10),
		Entry:       RandStringBytes(10),
		Delivery: restmodel.Delivery{
			Name:    RandStringBytes(10),
			Phone:   RandStringBytes(10),
			Zip:     RandStringBytes(10),
			City:    RandStringBytes(10),
			Address: RandStringBytes(10),
			Region:  RandStringBytes(10),
			Email:   RandStringBytes(10),
		},
		Payment: restmodel.Payment{
			Transaction:  RandStringBytes(10),
			RequestID:    RandStringBytes(10),
			Currency:     RandStringBytes(10),
			Provider:     RandStringBytes(10),
			Amount:       rand.Intn(10),
			PaymentDt:    rand.Intn(10),
			Bank:         RandStringBytes(10),
			DeliveryCost: rand.Intn(10),
			GoodsTotal:   rand.Intn(10),
			CustomFee:    rand.Intn(10),
		},
		Items: []restmodel.Item{
			{
				ChrtID:      rand.Intn(10),
				TrackNumber: RandStringBytes(10),
				Price:       rand.Intn(10),
				Rid:         RandStringBytes(10),
				Name:        RandStringBytes(10),
				Sale:        rand.Intn(10),
				Size:        RandStringBytes(10),
				TotalPrice:  rand.Intn(10),
				NmID:        rand.Intn(10),
				Brand:       RandStringBytes(10),
				Status:      rand.Intn(10),
			},
		},
		Locale:            RandStringBytes(10),
		InternalSignature: RandStringBytes(10),
		CustomerID:        RandStringBytes(10),
		DeliveryService:   RandStringBytes(10),
		Shardkey:          RandStringBytes(10),
		SmID:              rand.Intn(10),
		DateCreated:       time.Now(),
		OofShard:          RandStringBytes(10)}
	// Publish the message
	if err := ec.Publish("orders", &o); err != nil {
		log.Fatal(err)
	}
}
