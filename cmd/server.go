package cmd

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"testproject/confession/controller"
	"testproject/confession/env"
	"testproject/confession/router"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start Server",
	Run: func(cmd *cobra.Command, args []string) {
		server()
	},
}

func init() {
	RootCmd.AddCommand(serverCmd)
	// Here you will define your flags and configuration settings.
	serverCmd.Flags().StringVarP(&env.Port, "port", "p", "1324", "Server Port")
	serverCmd.Flags().BoolVarP(&env.Debug, "debug", "d", false, "Start Debug Mode")
	serverCmd.Flags().StringVarP(&env.Build, "build", "b", "0.0.0.0", "Server Build IP")
	serverCmd.Flags().BoolVar(&env.Pprof, "pprof", false, "Server Pprof IP")
	serverCmd.Flags().StringVarP(&env.TGToken, "tgtoken", "t", "0000", "TG Bot Token")
	serverCmd.Flags().Int64VarP(&env.LogchatID, "logchat", "l", 0, "Log Chat ID")
	serverCmd.Flags().Int64VarP(&env.ChatID, "chat", "c", 0, "Chat ID")
}

func server() {
	env.Load()
	// Create Echo server and hide extra banner info and port
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	e.Validator = &Validator{v: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	router.Router(e)

	controller.OnBot()

	// go func() {
	// 	c := wsclient.New(&env.CInfo)
	// 	c.Connect()
	// }()
	go func() {
		// Start server
		zap.L().Info("http server started on "+env.Build+":"+env.Port, zap.String("version", env.Version))
		if err := e.Start(env.Build + ":" + env.Port); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("start server failed", zap.Error(err))
		}
		// start client to table source(tSrc)
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("server is shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		zap.L().Fatal("shutdown server failed", zap.Error(err))
	}
}

func startPProf() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			zap.L().Error("PPROF error: " + err.Error())
		} else {
			zap.L().Info("PPROF enabled on port 6060")
		}
	}()
}

// Validator 驗證器
type Validator struct {
	v *validator.Validate
}

// Validate 驗證
func (v *Validator) Validate(i interface{}) error {
	return v.v.Struct(i)
}
