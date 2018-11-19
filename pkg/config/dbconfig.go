package config

//Real Time Filtering by device/alertid/or other tags

// InfluxCfg is the main configuration for any InfluxDB TSDB
type InfluxCfg struct {
	ID                 string `xorm:"'id' unique" binding:"Required"`
	Host               string `xorm:"host" binding:"Required"`
	Port               int    `xorm:"port" binding:"Required;IntegerNotZero"`
	DB                 string `xorm:"db" binding:"Required"`
	RWUser             string `xorm:"rwuser" binding:"Required"`
	RWPassword         string `xorm:"rwpassword" binding:"Required"`
	Retention          string `xorm:"'retention' default 'autogen'" binding:"Required"`
	RetentionTime      int    `xorm:"retentiontime"`
	ShardingTime       int    `xorm:"sharingtime" `
	RDUser             string `xorm:"rduser"`
	RDPassword         string `xorm:"rdpassword"`
	Precision          string `xorm:"'precision' default 's'" binding:"Default(s);OmitEmpty;In(h,m,s,ms,u,ns)"` //posible values [h,m,s,ms,u,ns] default seconds for the nature of data
	Timeout            int    `xorm:"'timeout' default 30" binding:"Default(30);IntegerNotZero"`
	UserAgent          string `xorm:"useragent" binding:"Default(ipashome)"`
	EnableSSL          bool   `xorm:"enable_ssl"`
	SSLCA              string `xorm:"ssl_ca"`
	SSLCert            string `xorm:"ssl_cert"`
	SSLKey             string `xorm:"ssl_key"`
	InsecureSkipVerify bool   `xorm:"insecure_skip_verify"`
	Description        string `xorm:"description"`
}

// ServiceCfg has all the Platform Device
type ServiceCfg struct {
	ID          string `xorm:"'id' unique" binding:"Required"`
	Label       string `xorm:"label" binding:"Required"` //title
	HeaderIcon  string `xorm:"header_icon"`
	LinearColor string `xorm:"linear_color"`
	FootContent string `xorm:"foot_content"`
	FooterIcon  string `xorm:"footer_icon"`
	Link        string `xorm:"link"`
	Description string `xorm:"description"`
	//--status
	StatusMode            string `xorm:"status_mode"`
	StatusURL             string `xorm:"status_url"`
	StatusContentType     string `xorm:"status_content_type"`
	StatusValidationMode  string `xorm:"status_validation_mode"`
	StatusValidationValue string `xorm:"status_validation_value"`
	//--Credentials
	AdmUser   string `xorm:"adm_user"`
	AdmPasswd string `xorm:"adm_passwd"`
} // ServiceCfg has all the Platform Device

// ProductDBMap a map for products
type ProductDBMap struct {
	ID          string `xorm:"'id' unique" binding:"Required"`
	Database    string `xorm:"database"`
	ProductTags string `xorm:"product_tags"`
}

// TableName go-xorm way to set the Table name to something different to "product_d_b_map"
func (ProductDBMap) TableName() string {
	return "product_db_map"
}

// DBConfig
type DBConfig struct {
	Services     map[string]*ServiceCfg
	Influxdb     map[string]*InfluxCfg
	ProductDbMap map[string]*ProductDBMap
}

// Init initialices the DB
func Init(cfg *DBConfig) error {

	log.Debug("--------------------Initializing Config-------------------")

	log.Debug("-----------------------END Config metrics----------------------")
	return nil
}
