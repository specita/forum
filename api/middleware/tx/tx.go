//基于gorm事务支持
package tx

import (
	"forum/model"
	"github.com/go-martini/martini"
)

func Transaction() martini.Handler {
	return func(c martini.Context) {
		tx := model.DB.Begin()
		//tx type *gorm.DB
		c.Map(tx)
		c.Next()

		if tx.Error != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}
}
