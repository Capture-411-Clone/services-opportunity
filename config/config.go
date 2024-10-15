package config

import (
	"fmt"
	"os"
	"reflect"

	"github.com/Capture-411/services-opportunity/kit/exp"
	"github.com/spf13/viper"
)

type Config struct {
	Name                              string `mapstructure:"NAME"`
	Environment                       string `mapstructure:"ENVIRONMENT"`
	Protocol                          string `mapstructure:"PROTOCOL"`
	ExternalPort                      int    `mapstructure:"EXTERNAL_PORT"`
	InternalPort                      int    `mapstructure:"INTERNAL_PORT"`
	Version                           string `mapstructure:"VERSION"`
	HostUri                           string `mapstructure:"HOST_URI"`
	HostAddress                       string `mapstructure:"HOST_ADDRESS"`
	ApiPrefix                         string `mapstructure:"API_PREFIX"`
	ApiVersion1                       string `mapstructure:"API_VERSION_1"`
	StaticDirPath                     string `mapstructure:"STATIC_DIR_PATH"`
	AccessTokenSigningKey             string `mapstructure:"ACCESS_TOKEN_SIGNING_KEY"`
	AccessTokenTokenExpiration        int    `mapstructure:"ACCESS_TOKEN_EXPIRATION"`
	RefreshTokenSigningKey            string `mapstructure:"REFRESH_TOKEN_SIGNING_KEY"`
	RefreshTokenExpiration            int    `mapstructure:"REFRESH_TOKEN_EXPIRATION"`
	InternalAuthToken                 string `mapstructure:"INTERNAL_AUTH_TOKEN"`
	DatabasePort                      int    `mapstructure:"DATABASE_PORT"`
	DatabaseName                      string `mapstructure:"DATABASE_NAME"`
	DatabaseHost                      string `mapstructure:"DATABASE_HOST"`
	DatabaseUsername                  string `mapstructure:"DATABASE_USERNAME"`
	DatabasePassword                  string `mapstructure:"DATABASE_PASSWORD"`
	DatabaseSslMode                   string `mapstructure:"DATABASE_SSL_MODE"`
	DatabaseSslRootCert               string `mapstructure:"DATABASE_SSL_ROOT_CERT"`
	DatabaseDefaultName               string `mapstructure:"DATABASE_DEFAULT_NAME"`
	DatabaseTimezone                  string `mapstructure:"DATABASE_TIMEZONE"`
	DatabaseStandalone                bool   `mapstructure:"DATABASE_STANDALONE"`
	SMTPHost                          string `mapstructure:"SMTP_HOST"`
	SMTPPort                          int    `mapstructure:"SMTP_PORT"`
	SMTPUsername                      string `mapstructure:"SMTP_USERNAME"`
	SMTPPassword                      string `mapstructure:"SMTP_PASSWORD"`
	SMTPFrom                          string `mapstructure:"SMTP_FROM"`
	LogEnable                         bool   `mapstructure:"LOG_ENABLE"`
	MailEnable                        bool   `mapstructure:"MAIL_ENABLE"`
	SmsApiKey                         string `mapstructure:"SMS_API_KEY"`
	MaxLoginDeviceCount               int    `mapstructure:"MAX_LOGIN_DEVICE_COUNT"`
	RequiredSignUpInviteCode          bool   `mapstructure:"REQUIRED_SIGN_UP_invite_code"`
	AutoDeleteDevice                  bool   `mapstructure:"AUTO_DELETE_DEVICE"`
	EmailDriver                       string `mapstructure:"EMAIL_DRIVER"`
	SendEmail                         bool   `mapstructure:"SEND_EMAIL"`
	ChatGptApiKey                     string `mapstructure:"CHAT_GPT_API_KEY"`
	EntHunterURI                      string `mapstructure:"ENT_HUNTER_URI"`
	EntHunterInternalApiKey           string `mapstructure:"ENT_HUNTER_INTERNAL_API_KEY"`
	BytebaseInternalAuthKey           string `mapstructure:"BYTEBASE_INTERNAL_AUTH_KEY"`
	ByteBaseInternalDownloadBaseUri   string `mapstructure:"BYTEBASE_INTERNAL_DOWNLOAD_BASE_URI"`
	AwsRegion                         string `mapstructure:"AWS_REGION"`
	AwsAccessKeyId                    string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AwsSecretAccessKey                string `mapstructure:"AWS_SECRET_ACCESS_KEY"`
	ContributionThreshold             int    `mapstructure:"CONTRIBUTION_THRESHOLD"`
	ContributionRequestTimeLimitHours int    `mapstructure:"CONTRIBUTION_REQUEST_TIME_LIMIT_HOURS"`
	ContributionRequestSizeLimit      int    `mapstructure:"CONTRIBUTION_REQUEST_SIZE_LIMIT"`
	VerificationCodeExpireMinutes     int    `mapstructure:"VERIFICATION_CODE_EXPIRE_MINUTES"`
	FileSuppressorPrefix              string `mapstructure:"FILE_SUPPRESSOR_PREFIX"`
	SAMApiKey                         string `mapstructure:"SAM_API_KEY"`

	StripeSecretKey      string `mapstructure:"STRIPE_SECRET_KEY"`
	StripeRedirectDomain string `mapstructure:"STRIPE_REDIRECT_DOMAIN"`
	StripeWebhookSecret  string `mapstructure:"STRIPE_WEBHOOK_SECRET"`
}

var AppConfig *Config = &Config{}

func Load() {
	fmt.Println("OS ENVIRONMENT: ", os.Getenv("ENVIRONMENT"))
	if os.Getenv("ENVIRONMENT") == "" {
		// default to development
		fmt.Println("ENVIRONMENT is not set, defaulting to development")
		os.Setenv("ENVIRONMENT", "development")
	}

	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env" + exp.TerIf(os.Getenv("ENVIRONMENT") != "", "."+os.Getenv("ENVIRONMENT"), ""))
	viper.AutomaticEnv() // read in environment variables that match

	// read config from system environment
	elem := reflect.TypeOf(AppConfig).Elem()
	for i := 0; i < elem.NumField(); i++ {
		key := elem.Field(i).Tag.Get("mapstructure")
		value := os.Getenv(key)
		if value != "" {
			viper.Set(key, value)
		}
	}

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	AppConfig.FileSuppressorPrefix = "file@"

	err := viper.Unmarshal(&AppConfig)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Config unmarshal error: ", err)
	}
}
