///////////////////////////////////////////////////////////
// THIS FILE IS AUTO GENERATED by gormgen, DON'T EDIT IT //
//        ANY CHANGES DONE HERE WILL BE LOST             //
///////////////////////////////////////////////////////////

package authorized_repo

import (
	"fmt"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/api/repository/db_repo"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func NewModel() *Authorized {
	return new(Authorized)
}

func NewQueryBuilder() *authorizedRepoQueryBuilder {
	return new(authorizedRepoQueryBuilder)
}

func (t *Authorized) Create(db *gorm.DB) (id int32, err error) {
	if err = db.Create(t).Error; err != nil {
		return 0, errors.Wrap(err, "create err")
	}
	return t.Id, nil
}

type authorizedRepoQueryBuilder struct {
	order []string
	where []struct {
		prefix string
		value  interface{}
	}
	limit  int
	offset int
}

func (qb *authorizedRepoQueryBuilder) buildQuery(db *gorm.DB) *gorm.DB {
	ret := db
	for _, where := range qb.where {
		ret = ret.Where(where.prefix, where.value)
	}
	for _, order := range qb.order {
		ret = ret.Order(order)
	}
	ret = ret.Limit(qb.limit).Offset(qb.offset)
	return ret
}

func (qb *authorizedRepoQueryBuilder) Updates(db *gorm.DB, m map[string]interface{}) (err error) {
	db = db.Model(&Authorized{})

	for _, where := range qb.where {
		db.Where(where.prefix, where.value)
	}

	if err = db.Updates(m).Error; err != nil {
		return errors.Wrap(err, "updates err")
	}
	return nil
}

func (qb *authorizedRepoQueryBuilder) Delete(db *gorm.DB) (err error) {
	for _, where := range qb.where {
		db = db.Where(where.prefix, where.value)
	}

	if err = db.Delete(&Authorized{}).Error; err != nil {
		return errors.Wrap(err, "delete err")
	}
	return nil
}

func (qb *authorizedRepoQueryBuilder) Count(db *gorm.DB) (int64, error) {
	var c int64
	res := qb.buildQuery(db).Model(&Authorized{}).Count(&c)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		c = 0
	}
	return c, res.Error
}

func (qb *authorizedRepoQueryBuilder) First(db *gorm.DB) (*Authorized, error) {
	ret := &Authorized{}
	res := qb.buildQuery(db).First(ret)
	if res.Error != nil && res.Error == gorm.ErrRecordNotFound {
		ret = nil
	}
	return ret, res.Error
}

func (qb *authorizedRepoQueryBuilder) QueryOne(db *gorm.DB) (*Authorized, error) {
	qb.limit = 1
	ret, err := qb.QueryAll(db)
	if len(ret) > 0 {
		return ret[0], err
	}
	return nil, err
}

func (qb *authorizedRepoQueryBuilder) QueryAll(db *gorm.DB) ([]*Authorized, error) {
	var ret []*Authorized
	err := qb.buildQuery(db).Find(&ret).Error
	return ret, err
}

func (qb *authorizedRepoQueryBuilder) Limit(limit int) *authorizedRepoQueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *authorizedRepoQueryBuilder) Offset(offset int) *authorizedRepoQueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereId(p db_repo.Predicate, value int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIdIn(value []int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIdNotIn(value []int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "id", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderById(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "id "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessKey(p db_repo.Predicate, value string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessKeyIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessKeyNotIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_key", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByBusinessKey(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_key "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessSecret(p db_repo.Predicate, value string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessSecretIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessSecretNotIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_secret", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByBusinessSecret(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_secret "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessDeveloper(p db_repo.Predicate, value string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessDeveloperIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereBusinessDeveloperNotIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "business_developer", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByBusinessDeveloper(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "business_developer "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereRemark(p db_repo.Predicate, value string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereRemarkIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereRemarkNotIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "remark", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByRemark(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "remark "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIsUsed(p db_repo.Predicate, value int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIsUsedIn(value []int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIsUsedNotIn(value []int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_used", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByIsUsed(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_used "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIsDeleted(p db_repo.Predicate, value int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIsDeletedIn(value []int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereIsDeletedNotIn(value []int32) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "is_deleted", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByIsDeleted(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "is_deleted "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereCreatedAt(p db_repo.Predicate, value time.Time) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereCreatedAtIn(value []time.Time) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereCreatedAtNotIn(value []time.Time) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByCreatedAt(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_at "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereCreatedUser(p db_repo.Predicate, value string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereCreatedUserIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereCreatedUserNotIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "created_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByCreatedUser(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "created_user "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereUpdatedAt(p db_repo.Predicate, value time.Time) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereUpdatedAtIn(value []time.Time) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereUpdatedAtNotIn(value []time.Time) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_at", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByUpdatedAt(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_at "+order)
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereUpdatedUser(p db_repo.Predicate, value string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", p),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereUpdatedUserIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) WhereUpdatedUserNotIn(value []string) *authorizedRepoQueryBuilder {
	qb.where = append(qb.where, struct {
		prefix string
		value  interface{}
	}{
		fmt.Sprintf("%v %v ?", "updated_user", "NOT IN"),
		value,
	})
	return qb
}

func (qb *authorizedRepoQueryBuilder) OrderByUpdatedUser(asc bool) *authorizedRepoQueryBuilder {
	order := "DESC"
	if asc {
		order = "ASC"
	}

	qb.order = append(qb.order, "updated_user "+order)
	return qb
}