package bot

// MongoConfig db connection info
type MongoConfig struct {
	URL      string `mapstructure:"url"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

// Messages to send
type Messages struct {
	Ready string `mapstructure:"ready"`
}

// ServiceConfig for the application
type ServiceConfig struct {
	Token       string      `mapstructure:"token"`
	Prefix      string      `mapstructure:"prefix"`
	MongoConfig MongoConfig `mapstructure:"mongo"`
	Msg         Messages    `mapstructure:"messages"`
}
