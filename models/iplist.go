package models

type Iplist struct {
	*Model
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	Protocol    string `json:"protocol"`
	Cms         string `json:"cms"`
	Language    string `json:"language"`
	Portnum     int    `json:"portnum"`
	Url         string `json:"url"`
	Loginurl    string `json:"loginurl"`
	Title       string `json:"title"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func GetIplistBrute(port int, protocol string) (iplist []Iplist) {
	db.Select("ip,port").Where("port = ?", port).Or("protocol = ?", protocol).Find(&iplist)
	return
}

//根据关键字获取分页查询结果
func GetIplist(pageNum int, pageSize int, maps interface{}) (iplist []Iplist) {
	dbTmp := db
	querys := maps.(map[string]interface{})

	if querys["protocol"] != nil {
		dbTmp = dbTmp.Where("protocol LIKE ?", "%"+querys["protocol"].(string)+"%")
	}


	if querys["ip"] != nil {
		dbTmp = dbTmp.Where("ip LIKE ?", "%"+querys["ip"].(string)+"%")
	}

	if querys["port"] != nil {
		dbTmp = dbTmp.Where("port LIKE ?", "%"+querys["port"].(string)+"%")
	}

	if querys["title"] != nil {
		dbTmp = dbTmp.Where("title LIKE ?", "%"+querys["title"].(string)+"%")
	}

	if querys["finger"] != nil {
		dbTmp = dbTmp.Where("cms LIKE ?", "%"+querys["finger"].(string)+"%")
	}

	dbTmp.Offset(pageNum).Limit(pageSize).Order("updated_time  desc").Find(&iplist)
	return
}



func GetIplistTotal(maps interface{}) (count int) {
	dbTmp := db
	querys := maps.(map[string]interface{})
	if querys["protocol"] != nil {
		dbTmp = dbTmp.Where("protocol LIKE ?", "%"+querys["protocol"].(string)+"%")
	}

	if querys["ip"] != nil {
		dbTmp = dbTmp.Where("ip LIKE ?", "%"+querys["ip"].(string)+"%")
	}

	if querys["port"] != nil {
		dbTmp = dbTmp.Where("port LIKE ?", "%"+querys["port"].(string)+"%")
	}

	if querys["title"] != nil {
		dbTmp = dbTmp.Where("title LIKE ?", "%"+querys["title"].(string)+"%")
	}

	if querys["finger"] != nil {
		dbTmp = dbTmp.Where("cms LIKE ?", "%"+querys["finger"].(string)+"%")
	}

	dbTmp.Model(&Iplist{}).Count(&count)
	return
}
