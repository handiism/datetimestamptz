package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err.Error())
	}

	dbUrl := viper.GetString("DATABASE_URL")

	conn, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	now := time.Now()
	interval := pgtype.Interval{Months: 10, Microseconds: 10, Days: 10, Valid: true}

	var id int
	var mytime pgtype.Time
	var mytimestamp pgtype.Timestamp
	var mytimestamptz pgtype.Timestamptz
	var mydate pgtype.Date
	var myinterval pgtype.Interval

	err = conn.QueryRow(context.Background(),
		`INSERT INTO datetimestamptz(
			time,
			timestamp,
			timestamptz,
			date,
			interval
		 ) VALUES($1,$2,$3,$4,$5)
		RETURNING
			id,
			time,
			timestamp,
			timestamptz,
			date,
			interval`,
		now, now, now, now, interval).Scan(&id, &mytime, &mytimestamp, &mytimestamptz, &mydate, &myinterval)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(id)
	fmt.Println(mytime)
	fmt.Println(mytimestamp)
	fmt.Println(mytimestamptz)
	fmt.Println(mydate)
	fmt.Println(myinterval)
}
