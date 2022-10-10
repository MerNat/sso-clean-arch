package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// httpClient := &http.Client{Timeout: time.Duration(5) * time.Second}
	ctx := context.Background()
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic("Failed to connect to RabbitMQ")
	}
	defer conn.Close()

	//Opening the channel to r.mq
	ch, err := conn.Channel()
	if err != nil {
		panic("Failed to open a channel")
	}

	defer ch.Close()
	//We need to make sure we declare our custom exchange in rabbitMq
	err = ch.ExchangeDeclare(
		"crypto-info", // name
		"fanout",      // type of fanout exchange
		true,          // durable - we need to reserve it's state after a restart of rabbitMQ and or for reserving after unbind
		false,         // auto-deleted - we need to auto-delete it if rabbitMQ restarts or
		false,         // internal - not internal exchange topology
		false,         // no-wait - no need to wait for a confirmation from server
		nil,           // arguments
	)

	if err != nil {
		panic("Failed to declare an exchange")
	}

	COIN_API_KEY := os.Getenv("COIN_API_KEY")

	fmt.Println(COIN_API_KEY)

	if COIN_API_KEY == "" {
		panic("Token Not provided")
	}
	quitProcess := make(chan bool)
	go requestAndSend(ch, ctx, quitProcess, COIN_API_KEY)

	//wait for 5 minutes and quit
	time.Sleep(1 * time.Minute)
	quitProcess <- true

	log.Println("\nDone requesting info and sending to r.mq")

}

func requestAndSend(ch *amqp.Channel, ctx context.Context, quitProcess chan bool, token string) {
	var result json.RawMessage
	counter := 0
	for {
		select {
		case <-quitProcess:
			return
		default:
			req, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil)
			if err != nil {
				panic(err)
			}
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("X-CoinAPI-Key", token)
			resp, err := http.DefaultClient.Do(req)

			if err != nil {
				panic(err)
			}

			if resp.StatusCode != http.StatusOK {
				panic("err")
			}

			err = json.NewDecoder(resp.Body).Decode(&result)

			if err != nil {
				panic(err)
			}

			byteData, err := result.MarshalJSON()

			if err != nil {
				panic(err)
			}

			//Send to r.mq
			//It's okay to drop messages if no consumer is connected i.e. (empty routing key)
			err = ch.PublishWithContext(
				ctx,
				"crypto-info", // name of exchange
				"",            // routing key
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        byteData,
				},
			)

			if err != nil {
				panic(err)
			}

			counter++

			log.Printf("Data sent -> %d times\n", counter)
			resp.Body.Close()

			time.Sleep(1 * time.Second)
		}
	}
}
