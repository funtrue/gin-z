package config

// APP   项目启动配置对象，定义了项目的基础信息
type APP struct {
	Env     string `mapstructure:"env" json:"env" yaml:"env"`                // 环境名称
	Port    string `mapstructure:"port" json:"port" yaml:"port"`             //服务监听端口号
	AppName string `mapstructure:"app_name" json:"app_name" yaml:"app_name"` //应用名称
	AppUrl  string `mapstructure:"app_url" json:"app_url" yaml:"app_url"`    //应用域名
}

// Log   日志对象，定义了日志的基础信息
type Log struct {
	Level      string `mapstructure:"level" json:"level" yaml:"level"`                   //日志等级
	RootDir    string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`          //日志根目录
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`          //日志文件名称
	Format     string `mapstructure:"format" json:"format" yaml:"format"`                //写入格式 可选json
	ShowLine   bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`       //是否显示调用行
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"` //旧文件的最大个数
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`          // 日志文件最大大小（MB）
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`             // 旧文件的最大保留天数
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress"`          //是否压缩
}

// Database   Database对象，定义了数据库的基础信息
type Database struct {
	Driver              string `mapstructure:"driver" json:"driver" yaml:"driver"`                                                 //数据库驱动
	Host                string `mapstructure:"host" json:"host" yaml:"host"`                                                       //域名
	Port                int    `mapstructure:"port" json:"port" yaml:"port"`                                                       //端口号
	Database            string `mapstructure:"database" json:"database" yaml:"database"`                                           //数据库名称
	UserName            string `mapstructure:"username" json:"username" yaml:"username"`                                           //用户名
	Password            string `mapstructure:"password" json:"password" yaml:"password"`                                           //密码
	Charset             string `mapstructure:"charset" json:"charset" yaml:"charset"`                                              //编码格式
	MaxIdleConns        int    `mapstructure:"max_idle_conns" json:"max_idle_conns" yaml:"max_idle_conns"`                         //空闲连接池中连接的最大数量
	MaxOpenConns        int    `mapstructure:"max_open_conns" json:"max_open_conns" yaml:"max_open_conns"`                         //打开数据库连接的最大数量
	LogMode             string `mapstructure:"log_mode" json:"log_mode" yaml:"log_mode"`                                           //日志级别
	EnableFileLogWriter bool   `mapstructure:"enable_file_log_writer" json:"enable_file_log_writer" yaml:"enable_file_log_writer"` //是否启用日志文件
	LogFilename         string `mapstructure:"log_filename" json:"log_filename" yaml:"log_filename"`                               //日志文件名称
}

// Jwt   Jwt对象，定义了Jwt的基础信息
type Jwt struct {
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtTtl                  int64  `mapstructure:"jwt_ttl" json:"jwt_ttl" yaml:"jwt_ttl"`                                                          // token 有效期（秒）
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"` // 黑名单宽限时间（秒）
	RefreshGracePeriod      int64  `mapstructure:"refresh_grace_period" json:"refresh_grace_period" yaml:"refresh_grace_period"`                   // token 自动刷新宽限时间（秒）
}

// Redis   Redis对象，定义了Redis基础信息
type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             //redis地址
	Port     string `mapstructure:"port" json:"port" yaml:"port"`             //redis端口
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   //redis库号
	Password string `mapstructure:"password" json:"password" yaml:"password"` //redis密码
}

// Configuration   项目对象，定义了项目的基础信息
type Configuration struct {
	App   APP      `mapstructure:"app" json:"app" yaml:"app"`
	Log   Log      `mapstructure:"log" json:"log" yaml:"log"`
	Db    Database `mapstructure:"database" json:"database" yaml:"database"`
	Jwt   Jwt      `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis Redis    `mapstructure:"redis" json:"redis" yaml:"redis"`
}
