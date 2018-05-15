package cache

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/garyburd/redigo/redis"
)

func ipToScore(ipaddress string) (score int, err error) {
	iplist := strings.Split(ipaddress, ".")
	for _, num := range iplist {
		i, err := strconv.Atoi(num)
		if err != nil {
			return 0, err
		}
		score = score*256 + i
	}
	return
}

func uint8ListToString(ulist []uint8) string {
	bappend := make([]byte, 0, len(ulist))
	for _, b := range ulist {
		bappend = append(bappend, byte(b))
	}
	return string(bappend)

}

func getGeoIPInfoByIP(ipaddress string, lang string) (results []string, err error) {
	client := GetDBHandler()
	score, err := ipToScore(ipaddress)
	if err != nil {
		return
	}
	value1, err := redis.Values(client.Do("ZREVRANGEBYSCORE", "ip2cityid:", score, 0, "withscores", "limit", 0, 1))
	if err != nil {
		return
	}
	var _score string
	var _cityID string
	_, err = redis.Scan(value1, &_cityID, &_score)
	if err != nil {
		return
	}

	cityID, err := strconv.Atoi(strings.Split(_cityID, "_")[0])
	if err != nil {
		return
	}
	value2, err := client.Do("HGET", "cityid2city:"+lang+":", cityID)
	if err != nil {
		return
	}
	err = json.Unmarshal(value2.([]uint8), &results)

	if err != nil {
		return
	}
	return results, nil
}
