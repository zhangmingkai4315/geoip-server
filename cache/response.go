package cache

// IPInfo hold all data geoinfo
type IPInfo struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	Region  string `json:"region"`
	State   string `json:"state"`
	City    string `json:"city"`
}

// GetInfo get one ip information from cache
func (ipinfo *IPInfo) GetInfo() (string, error) {
	ip := ipinfo.IP
	info, err := getGeoIPInfoByIP(ip)
	if err != nil {
		return "", err
	}
	return info, nil
}

// IPInfoList hold all ip list info
type IPInfoList []IPInfo
