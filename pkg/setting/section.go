package setting

type Config struct {
	MySQL MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Url             string `mapstructure:"url"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifetime int    `mapstructure:"connMaxLifeTime"`
}
