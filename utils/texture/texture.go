package texture

import "encoding/json"

func To(src interface{}, dst interface{}) error {
	orderJson, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(orderJson, dst)
	if err != nil {
		return err
	}

	return nil
}
