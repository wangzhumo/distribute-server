package conf

const DriverName = "mysql"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Psd      string
	Database string
	Running  bool
}

var DbConfigList = []DbConfig{
	{
		Host:     "127.0.0.1",
		Port:     "3306",
		User:     "root",
		Psd:      "wzmmysql",
		Database: "distribute",
		Running:  true,
	},
}

var DbMaster DbConfig = DbConfigList[0]
