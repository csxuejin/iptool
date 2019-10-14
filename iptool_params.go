package iptool

type GetIPGeoOutput struct {
	Code int             `json:"code"`
	Data TaobaoIPGeoInfo `json:"data"`
}

type TaobaoIPGeoInfo struct {
	IP        string `json:"ip"`
	Country   string `json:"country"`
	Area      string `json:"area"`
	Region    string `json:"region"`
	City      string `json:"city"`
	County    string `json:"county"`
	ISP       string `json:"isp"`
	CountryID string `json:"country_id"`
	AreaID    string `json:"area_id"`
	RegionID  string `json:"region_id"`
	CityID    string `json:"city_id"`
	CountyID  string `json:"county_id"`
	IspID     string `json:"isp_id"`
}
