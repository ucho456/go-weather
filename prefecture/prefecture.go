package prefecture

import "fmt"

var Map = map[string]string{
	"東京": "130000",
	"大阪": "270000",
}

func GetPrefectureCode(m string) (string, error) {
	code, ok := Map[m]
	if !ok {
		err := fmt.Errorf("error wrong prefecture: %s", m)
		return "", err
	}
	return code, nil
}
