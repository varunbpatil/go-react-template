package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/varunbpatil/go-react-template/config"
	usersService "github.com/varunbpatil/go-react-template/domains/users/service"
	"github.com/varunbpatil/go-react-template/inbound/grpc"
	usersGRPC "github.com/varunbpatil/go-react-template/inbound/grpc/users"
	"github.com/varunbpatil/go-react-template/inbound/http"
	"github.com/varunbpatil/go-react-template/inbound/mcp"
	usersMCP "github.com/varunbpatil/go-react-template/inbound/mcp/users"
	usersRepository "github.com/varunbpatil/go-react-template/outbound/postgres/users"
	"github.com/varunbpatil/go-react-template/types"
)

const shutdownTimeout = 10 * time.Second

func main() {
	// Deferred functions do not run when os.Exit() is called in the same function.
	// That is the reason we put all the logic in run() which can then use deferred functions.
	os.Exit(run())
}

func run() int {
	// Parse application configuration
	cfg, err := config.Parse()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}

	// Setup logging
	logger := slog.New(newLogHandler(cfg.Log))
	slog.SetDefault(logger)

	// Context with signal cancellation
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	// onFatal is a closure that you can pass to any service below so that the service
	// itself can initiate a clean shutdown of the entire application on some fatal error
	onFatal := func(err error) { logger.Error("fatal service error", "error", err); cancel() }

	// Lifecycle manager
	lm := types.NewManager(logger)
	defer func() { cancel(); lm.StopAll(shutdownTimeout) }()

	// Setup users repository
	usersRepo, err := usersRepository.New()
	if err != nil {
		logger.Error("create users repository", "error", err)
		return 1
	}
	lm.AddCloser("Users repository", usersRepo)

	// Setup users service
	usersSvc, err := usersService.New(usersRepo)
	if err != nil {
		logger.Error("create users service", "error", err)
		return 1
	}
	lm.Add("Users service", usersSvc)

	// Setup the gRPC adapter
	grpcSrv := grpc.New(cfg.GRPC.Address, logger, onFatal)
	usersGRPC.Register(grpcSrv, usersGRPC.New(usersSvc))
	lm.Add("gRPC", grpcSrv)

	// Setup the MCP adapter
	mcpSrv := mcp.New()
	usersMCP.Register(mcpSrv, usersSvc)

	// Setup the HTTP adapter
	httpSrv := http.New(grpcSrv.Handler(), mcp.Handler(mcpSrv), cfg.HTTP.Address, logger, onFatal)
	lm.Add("HTTP", httpSrv)

	// Start all services
	if startErr := lm.StartAll(ctx); startErr != nil {
		logger.Error("startup failed", "error", startErr)
		return 1
	}

	// Shutdown gracefully
	<-ctx.Done()
	return 0
}

func newLogHandler(cfg config.LogConfig) slog.Handler {
	opts := &slog.HandlerOptions{Level: parseLevel(cfg.Level)}

	switch strings.ToLower(cfg.Format) {
	case "json":
		return slog.NewJSONHandler(os.Stderr, opts)
	default:
		return slog.NewTextHandler(os.Stderr, opts)
	}
}

func parseLevel(s string) slog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
