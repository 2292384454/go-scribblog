package log

type Options struct {
	FilePath   string `yaml:"file_path"`
	FileName   string `yaml:"file_name"`
	MaxSize    int    `yaml:"max_size"` //单个日志文件最大空间占用，单位MB
	MaxBackups int    `yaml:"max_backups"`
	MaxAge     int    `yaml:"max_age"`   //日志过期天数，过期日志会被滚动清理
	LogLevel   Level  `yaml:"log_level"` //日志等级
	Console    bool   `yaml:"console"`   //是否在stdout打印日志
}
