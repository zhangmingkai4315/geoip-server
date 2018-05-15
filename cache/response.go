package cache

// IPInfo hold all data geoinfo
type IPInfo struct {
	IP          string `json:"ip"`
	Continent   string `json:"continent"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
	Subdivision string `json:"subdivision"`
	City        string `json:"city"`
	Metro       string `json:"metro"`
}

// GetInfo get one ip information from cache
func (ipinfo *IPInfo) GetInfo(lang string) error {
	ip := ipinfo.IP
	info, err := getGeoIPInfoByIP(ip, lang)
	ipinfo.Continent = info[1]
	ipinfo.Country = info[3]
	ipinfo.CountryCode = info[2]
	ipinfo.City = info[6]
	ipinfo.Subdivision = info[5]
	ipinfo.Metro = info[7]

	if err != nil {
		return err
	}
	return nil
}

// IPInfoList hold all ip list info
type IPInfoList []IPInfo
