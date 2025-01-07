package etc

type Configuration struct {
	Db     db     `toml:"db"`
	Web    web    `toml:"web"`
	Redis  redis  `toml:"redis"`
	Token  token  `toml:"token"`
	Notify notify `toml:"notify"`
	Ttl    ttl    `toml:"ttl"`
}

type web struct {
	Listen string `toml:"listen"`
}

type redis struct {
	Enable   bool
	Addr     string
	Password string
	Db       int
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}

type token struct {
	Enable bool
	Issuer string
	Key    string
	Ttl    int
}
type notify struct {
	Host   string
	AppKey string
	Secret string
}
type ttl struct {
	Time     int64
	TimeType string
}
