package cache

import (
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

func getGeoIPInfoByIP(ip string) (string, error) {
	client := GetDBHandler()
	score, err := ipToScore(ip)
	if err != nil {
		return "", err
	}
	results, err := redis.Strings(client.Do("ZREVRANGEBYSCORE", "ip2cityid:", score, 0, "WIHTSCORES", "Limit", 0, 1))
	if err != nil {
		return "", err
	}
	return strings.Join(results, ""), nil
}
