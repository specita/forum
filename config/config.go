package config

var Config = struct {
	Debug bool `default:"false"`

	APPName string `default:"forum"`

	Address string `default:"127.0.0.1:5000"`

	Database struct {
		DSN string `default:""`
	}

	Logger struct {
		Level         string `default:"info"`
		Format        string `default:"json"`
		Out           string `default:"stdout"`
		FluentdEnable bool   `default:"false"`
		FluentdHost   string `default:""`
		FluentdPort   int    `default:""`
		FluentdTag    string `default:""`
	}
}{}
