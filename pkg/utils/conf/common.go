package conf

type Server struct {
	Addr       string
	EnableCors bool
	EnableCsrf bool
	Logger     bool
}

type Gorm struct {
	Debug           bool
	Dsn             string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int64
	TablePrefix     string
}

type Jwt struct {
	Exp     int64
	Iss     string
	Signing string
}
