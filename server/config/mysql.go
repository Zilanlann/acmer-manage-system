package config

type MySQL struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *MySQL) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Path + ":" + m.Port + ")/" + m.Dbname + "?" + m.Config
}

func (m *MySQL) GetLogMode() string {
	return m.LogMode
}
