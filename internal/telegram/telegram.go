package telegram

import (
	"context"
	"fmt"
	"github.com/celestix/gotgproto"
	"github.com/guisecreator/pintebot/internal/config"
	"github.com/guisecreator/pintebot/internal/types"
	"github.com/mymmrac/telego"
	th "github.com/mymmrac/telego/telegohandler"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type TgBotService struct {
	Bot         *telego.Bot
	BotServices types.BotServices
	Client      gotgproto.Client
	Handlers    *th.BotHandler
	Logger      logrus.Logger
	stop        chan struct{}
	done        chan struct{}
}

func NewTelegram(
	cfg config.Config,
	botServices types.BotServices,
	token string,
	log *logrus.Logger,
) (*TgBotService, error) {
	apiBotService := &TgBotService{}

	bot, err := telego.NewBot(
		token,
		telego.WithDefaultDebugLogger(),
	)
	if err != nil {
		return nil, fmt.Errorf("bot: %v", err)
	}

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return nil, fmt.Errorf("updates: %v", err)
	}

	botHandler, err := th.NewBotHandler(
		bot,
		updates,
		th.WithStopTimeout(5*time.Second),
	)
	if err != nil {
		return nil, fmt.Errorf("bh: %v", err)
	}

	botHandler.Group()

	//Middleware here
	botHandler.Use(func(bot *telego.Bot, update telego.Update, next th.Handler) {
		next(bot, update)
	})

	botServices = types.BotServices{
		Config: &cfg,
	}

	apiBotService = &TgBotService{
		Bot:         bot,
		BotServices: botServices,
		Handlers:    botHandler,
		Logger:      *log,
	}

	return apiBotService, nil
}

func (service *TgBotService) StartService(ctx context.Context) error {
	service.handleStopSignal()

	service.handlersInit()

	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		<-signals
		service.Stop()
		service.
			Logger.
			Info("Telegram service stopped")
	}()

	go service.Handlers.Start()
	service.
		Logger.
		Info("Telegram service started")

	if !service.Handlers.IsRunning() {
		return fmt.Errorf("service is not running")
	}

	<-service.done

	return nil
}

func (service *TgBotService) Stop() {
	go func() {
		service.stop <- struct{}{}
	}()
}

func (service *TgBotService) handleStopSignal() {
	go func() {
		<-service.stop

		service.
			Logger.
			Info("Stopping...")

		service.Bot.StopLongPolling()
		service.
			Logger.
			Info("Long polling done")

		service.Handlers.Stop()
		service.
			Logger.
			Info("Bot handler done")

		service.done <- struct{}{}
	}()
}
