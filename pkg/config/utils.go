package config

import "strings"

func GetMapWithoutPrefix(set map[string]string,prefix string) map[string]string{
	m := make(map[string]string)

	for key,value := range set{
		if strings.HasPrefix(key,prefix){
			m[strings.TrimPrefix(key,prefix)] = value
		}
	}
	if len(m) == 0{
		return nil
	}

	return m
}
