package utils

import "k8s.io/apimachinery/pkg/util/json"

func Struct2map(s interface{}) map[string]string {
	bytes, _ := json.Marshal(s)
	m := make(map[string]string)
	json.Unmarshal(bytes, &m)
	return m
}
