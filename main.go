package main

import (
	"context"
	"fmt"
	"hw/feature_sql/connection"
)

func main() {
	ctx := context.Background()
	conn, err := connection.ConnectionDB(ctx)
	if err != nil {
		panic(err)
	}
	err = conn.Ping(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Успешнок подключение!")

}
