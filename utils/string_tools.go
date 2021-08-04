package utils

import (
	"encoding/json"
	"time"
)

func JsonToMap(jsonStr string) (map[string]interface{}, error) {
	/*
		example：
			fmt.Println(fmt.Sprintf("%s", res[key]))
	*/
	var res map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func JsonListToSlice(jsonStr string) ([]map[string]interface{}, error) {
	/*
		example：
			for res := result {
				fmt.Println(fmt.Sprintf("%s", res[key]))
			}
	*/
	var res []map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &res)
	if err != nil {
		return res, err
	}
	return res, nil
}

func StringToTime(timeStr string) (time.Time, error) {
	// RFC3999 time string to time.Time
	t, err := time.Parse(
		time.RFC3339,
		timeStr,
	)
	if err != nil {
		return t, err
	}
	return t, nil
}

func StringToTimestamp(timeStr string) (int64, error) {
	// RFC3999 time string to timestamp
	t, err := time.Parse(
		time.RFC3339,
		timeStr,
	)
	if err != nil {
		return int64(0), err
	}
	return int64(t.UTC().Unix()), nil
}
