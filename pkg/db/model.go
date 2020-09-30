package db

// 插件
type MPlugin struct {
	Id          int64
	Name        string `gorm:"column:name;type:varchar(255);default:'';not null;unique_index:idx_plugin_name"`
	Info        string `gorm:"column:info;type:varchar(1024);default:''"`
	InstallPath string `gorm:"column:install_path;type:varchar(255);default:'/';not null"`
}

// 插件版本
type MPluginVersion struct {
	Id          int64
	PluginKeyId int64  `gorm:"column:plugin_key_id;not null;unique_index:idx_ver"`
	Version     string `gorm:"column:version;type:varchar(64);unique_index:idx_ver"`     // 版本号
	BuildTime   string `gorm:"column:build_time;type:varchar(255);unique_index:idx_ver"` // 编译时间
	ChangeInfo  string `gorm:"column:change_info;type:varchar(255);default:''"`          // 更新说明
	PublishTime int64  `gorm:"column:publish_time"`                                      // 发布时间
	PublishFlag int    `gorm:"column:publish_flag;type:int;default:0"`                   // 1 正式版, 2 测试版
	FileId      string `gorm:"column:file_id;type:varchar(255);default:'';not null"`     // 文件位置
}
