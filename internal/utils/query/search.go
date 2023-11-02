package utils

import (
	"fmt"
	"net/http"
	"strings"

	rest_goswagger_api "github.com/ramadoiranedar/rest-goswagger-api"
	"github.com/ramadoiranedar/rest-goswagger-api/internal/utils/constants"
	"gorm.io/gorm"
)

func ApplySearch(rt *rest_goswagger_api.Runtime, db *gorm.DB, searchBy, searchKeyword *string) error {
	if searchKeyword != nil {
		if searchBy == nil {
			return rt.SetError(http.StatusUnprocessableEntity, constants.MSG_REQUIRED_SEARCHBY)
		}
		db = db.Where(fmt.Sprintf("%s LIKE ?", *searchBy), strings.Join([]string{"%", *searchKeyword, "%"}, ""))
	}

	return nil
}
