package av

import (
	"fmt"
	"sort"
	"strconv"
)

// Param is a struct for the parameters of a GET call to Alpha Vantage
type Param struct {
	Function   string
	Symbol     string
	Interval   string
	Outputsize string // optional
	Datatype   string // optional
	APIKey     string
}

// TimeSeriesDaily is a type meant to receive an unmarshalled json from Alpha Vantage
// containing a series of data with 1 minute intervals.
type TimeSeriesDaily struct {
	Metadata struct {
		Information   string `json:"1. Information"`
		Symbol        string `json:"2. Symbol"`
		LastRefreshed string `json:"3. Last Refreshed"`
		Interval      string `json:"4. Interval"`
		OutputSize    string `json:"5. Output Size"`
		TimeZone      string `json:"6. Time Zone"`
	} `json:"Meta Data"`
	Series InterfaceMap `json:"Time Series (Daily)"`
}

// MapSorter is an interface for types that have the MapSort method.
type MapSorter interface {
	MapSort()
}

// InterfaceMap
type InterfaceMap map[string]interface{}

// FloatMap
type FloatMap map[string]float32

// StringMap is a map type whose keys are strings. Implements MapSorter.
type StringMap map[string]string

// InterfaceSlice is a
type InterfaceSlice []interface{}

// MapSort is a method which takes a map and returns a sorted slice of the keys.
func (m StringMap) MapSort() []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// MapSort for InterfaceMaps
func (m InterfaceMap) MapSort() []string {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// ConvertToStrings takes an InterfaceSlice and converts it to a slice of strings
func (s InterfaceSlice) ConvertToStrings() []string {
	stringSlice := make([]string, len(s))
	for i, v := range s {
		stringSlice[i] = fmt.Sprint(v)
	}
	return stringSlice
}

func (m InterfaceMap) ConvertToFloats() FloatMap {
	floatMap := make(map[string]float32, len(m))
	for k, v := range m {
		// cant just cast a string as a float
		float, err := strconv.ParseFloat(v.(string), 32)
		if err != nil {
			panic(err)
		}
		floatMap[k] = float32(float)
	}
	return floatMap
}
