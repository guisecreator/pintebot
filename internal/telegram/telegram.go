package telegram

import (
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
	logg *logrus.Logger,
) (*TgBotService, error) {
	apiBotService := &TgBotService{}

	bot, err := telego.NewBot(
		token,
		telego.WithDefaultDebugLogger(),
	)
	if err != nil {
		return nil, err
	}

	updates, err := bot.UpdatesViaLongPolling(nil)
	if err != nil {
		return nil, err
	}

	botHandler, err := th.NewBotHandler(
		bot,
		updates,
		th.WithStopTimeout(5*time.Second),
	)
	if err != nil {
		return nil, err
	}

	botHandler.Group()

	done := make(chan struct{}, 1)
	stop := make(chan struct{}, 1)

	//middleware here
	botHandler.Use(
		func(bot *telego.Bot, update telego.Update, next th.Handler) {
			go func() {
				defer func() {
					if r := recover(); r != nil {
						logg.Error(r)
					}
				}()
				next(bot, update)
			}()
		},
	)

	botServices = types.BotServices{
		Config: &cfg,
	}

	apiBotService = &TgBotService{
		Bot:         bot,
		BotServices: botServices,
		Handlers:    botHandler,
		Logger:      *logg,
		stop:        stop,
		done:        done,
	}

	return apiBotService, nil
}

func (service *TgBotService) StartService() error {
	service.handleStopSignal()

	service.handlersInit()

	signals := make(chan os.Signal, 1)
	signal.Notify(
		signals,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go service.Handlers.Start()
	service.
		Logger.
		Info("Telegram service started")

	if !service.Handlers.IsRunning() {
		service.
			Logger.
			Fatal("Telegram service crashed or could not start")
		return nil
	}

	//stopping the bot
	go func() {
		<-signals

		go func() {
			service.stop <- struct{}{}
			service.
				Logger.
				Info("Telegram service stopped")
		}()
	}()

	<-service.done

	return nil
}

func (service *TgBotService) handleStopSignal() {
	go func() {
		<-service.stop

		service.
			Logger.
			Info("Stopping the rest...")

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
