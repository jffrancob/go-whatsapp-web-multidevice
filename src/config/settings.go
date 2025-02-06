package config

import (
	"go.mau.fi/whatsmeow/proto/waCompanionReg"
)

var (
	AppVersion             = "v5.0.0"
	AppPort                = "3000"
	AppDebug               = false
	AppOs                  = "AldinoKemal"
	AppPlatform            = waCompanionReg.DeviceProps_PlatformType(1)
	AppBasicAuthCredential []string

	PathQrCode      = "statics/qrcode"
	PathSendItems   = "statics/senditems"
	PathMedia       = "statics/media"
	PathStorages    = "storages"
	PathChatStorage = "storages/chat.txt"

	DBURI = "file:storages/whatsapp.db?_foreign_keys=off"

	WhatsappAutoReplyMessage       string
	WhatsappWebhook                []string
	WhatsappWebhookSecret                = "secret"
	WhatsappLogLevel                     = "ERROR"
	WhatsappSettingMaxFileSize     int64 = 50000000  // 50MB
	WhatsappSettingMaxVideoSize    int64 = 100000000 // 100MB
	WhatsappSettingMaxDownloadSize int64 = 500000000 // 500MB
	WhatsappTypeUser                     = "@s.whatsapp.net"
	WhatsappTypeGroup                    = "@g.us"
	WhatsappAccountValidation            = true
)
