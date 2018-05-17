package cache

// IPInfo hold all data geoinfo
type IPInfo struct {
	IP          string `json:"ip"`
	Continent   string `json:"continent,omitempty" `
	Country     string `json:"country,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Subdivision string `json:"subdivision,omitempty"`
	City        string `json:"city,omitempty"`
	Metro       string `json:"metro,omitempty"`
	ASN         string `json:"asn,omitempty"`
	Orgnazation string `json:"orgnazation,omitempty"`
	Error       string `json:"error,omitempty"`
}

// GetIPInfo get one ip information from cache
func (ipinfo *IPInfo) GetIPInfo(lang string) error {
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

// GetASNInfo get one ip asn information from cache
func (ipinfo *IPInfo) GetASNInfo() error {
	ip := ipinfo.IP
	info, err := getGeoIPASNInfoByIP(ip)
	if err != nil {
		return err
	}
	ipinfo.ASN = info[0]
	ipinfo.Orgnazation = info[1]
	return nil
}

// GetAllInfo will gather all asn data from cache database
// based ip list, if validate error then skip
func (ipinfo *IPInfo) GetAllInfo() error {
	err := ipinfo.GetIPInfo("en")
	if err != nil {
		return err
	}
	err = ipinfo.GetASNInfo()
	if err != nil {
		return err
	}
	return nil
}

// IPInfoList hold all ip list info
type IPInfoList []*IPInfo

// GetIPInfo will gather all ipinfo data from cache database
// based ipinfo list, if validate error then skip
func (ipinfoList IPInfoList) GetIPInfo(lang string) error {
	for _, ipinfo := range ipinfoList {
		if ipinfo.Error != "" {
			continue
		}
		info, err := getGeoIPInfoByIP(ipinfo.IP, lang)
		if err != nil {
			ipinfo.Error = err.Error()
			continue
		}
		ipinfo.Continent = info[1]
		ipinfo.Country = info[3]
		ipinfo.CountryCode = info[2]
		ipinfo.City = info[6]
		ipinfo.Subdivision = info[5]
		ipinfo.Metro = info[7]
	}
	return nil
}

// GetASNInfo will gather all asn data from cache database
// based ip list, if validate error then skip
func (ipinfoList IPInfoList) GetASNInfo() error {
	for _, ipinfo := range ipinfoList {
		if ipinfo.Error != "" {
			continue
		}
		info, err := getGeoIPASNInfoByIP(ipinfo.IP)
		if err != nil {
			ipinfo.Error = err.Error()
			continue
		}
		ipinfo.ASN = info[0]
		ipinfo.Orgnazation = info[1]
	}
	return nil
}

// GetAllInfo will gather all asn data from cache database
// based ip list, if validate error then skip
func (ipinfoList IPInfoList) GetAllInfo() error {
	err := ipinfoList.GetIPInfo("en")
	if err != nil {
		return err
	}
	err = ipinfoList.GetASNInfo()
	if err != nil {
		return err
	}
	return nil
}
