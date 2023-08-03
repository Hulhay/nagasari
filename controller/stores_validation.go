package controller

import "github.com/hulhay/nagasari/lib/utils"

func validateGetStoresRequest(keyword string) error {
	if keyword != "" && len(keyword) < 3 {
		return utils.ErrKeywordLT3
	}

	return nil
}
