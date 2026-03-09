package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/raoptimus/archimate-mcp/client"
	"github.com/raoptimus/archimate-mcp/tools"
	"github.com/urfave/cli/v3"
)

const (
	serverName            = "archimate-mcp"
	serverVersion         = "1.0.0"
	defaultRequestTimeout = 30 * time.Second
	defaultArchiURL       = "http://localhost:9898"
)

func main() {
	var (
		archiURL string
		debug    bool
		timeout  time.Duration
	)

	cmd := &cli.Command{
		Name:    serverName,
		Usage:   "MCP server for ArchiMate models via jArchi HTTP bridge",
		Version: serverVersion,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "archi-url",
				Aliases:     []string{"u"},
				Usage:       "jArchi HTTP server URL",
				Sources:     cli.EnvVars("ARCHI_URL"),
				Value:       defaultArchiURL,
				Destination: &archiURL,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "Enable debug logging",
				Sources:     cli.EnvVars("ARCHI_DEBUG"),
				Destination: &debug,
			},
			&cli.DurationFlag{
				Name:        "timeout",
				Usage:       "HTTP request timeout",
				Sources:     cli.EnvVars("ARCHI_TIMEOUT"),
				Value:       defaultRequestTimeout,
				Destination: &timeout,
			},
		},
		Writer:    os.Stderr,
		ErrWriter: os.Stderr,
		Action: func(ctx context.Context, cmd *cli.Command) error {
			return runServer(ctx, archiURL, debug, timeout)
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatalf("Fatal error: %v", err)
	}
}

func runServer(ctx context.Context, archiURL string, debug bool, timeout time.Duration) error {
	loggerLevel := slog.LevelWarn
	if debug {
		loggerLevel = slog.LevelDebug
	}
	logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: loggerLevel,
	}))

	archiClient := client.NewClient(archiURL, timeout)

	server := mcp.NewServer(
		&mcp.Implementation{
			Name:    serverName,
			Version: serverVersion,
		},
		&mcp.ServerOptions{
			Logger: logger,
			Capabilities: &mcp.ServerCapabilities{
				Tools: &mcp.ToolCapabilities{ListChanged: true},
			},
		},
	)

	registry := tools.NewRegistry(archiClient)
	registry.RegisterAll(server)

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		cancel()
	}()

	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		return fmt.Errorf("server error: %w", err)
	}

	return nil
}
