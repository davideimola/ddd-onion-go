package server

import (
	"context"
	orderHTTPService "davideimola.dev/ddd-onion/pkg/order/infra/http"
	pgOrderRepository "davideimola.dev/ddd-onion/pkg/order/repos/pg"
	orderService "davideimola.dev/ddd-onion/pkg/order/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/spf13/cobra"
	"os"
)

var (
	Cmd = &cobra.Command{
		Use:   "server",
		Short: "Start the HTTP server",
		Run:   run,
	}
)

func run(cmd *cobra.Command, args []string) {
	dbUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		panic("DATABASE_URL env var is required")
	}

	db, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(err)
	}

	orderHTTP := orderHTTPService.New(orderService.New(pgOrderRepository.NewOrderRepository(db)))

	router := gin.Default()

	router.POST("/order", orderHTTP.PostOrder)

	router.Run(":8080")
}
