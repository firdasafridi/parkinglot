package log

const (
	DBG = "debug"
	ERR = "error"
	INF = "info"
	FTL = "fatal"
	PNC = "panic"
	NL  = "no_level"
)

var (
	defaultLog = &Config{IsJson: true}
)

type Config struct {
	IsJson bool `json:"is_json" yaml:"is_json"`
}

func New(cfg *Config) {
	if cfg == nil {
		return
	}
	defaultLog = cfg
}

func (cfg *Config) Debug() *Event {
	return NewEvent(DBG)
}

func (cfg *Config) Error() *Event {
	return NewEvent(ERR)
}

func (cfg *Config) Info() *Event {
	return NewEvent(INF)
}

func (cfg *Config) Fatal() *Event {
	return NewEvent(FTL)
}

func (cfg *Config) Panic() *Event {
	return NewEvent(PNC)
}

func (cfg *Config) NoLevel() *Event {
	return NewEvent(NL)
}
