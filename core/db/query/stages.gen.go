// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"slyfox-tails/db/models"
)

func newStage(db *gorm.DB, opts ...gen.DOOption) stage {
	_stage := stage{}

	_stage.stageDo.UseDB(db, opts...)
	_stage.stageDo.UseModel(&models.Stage{})

	tableName := _stage.stageDo.TableName()
	_stage.ALL = field.NewAsterisk(tableName)
	_stage.ID = field.NewUint64(tableName, "id")
	_stage.Title = field.NewString(tableName, "title")
	_stage.CreatorID = field.NewUint64(tableName, "creator_id")
	_stage.JobID = field.NewUint64(tableName, "job_id")
	_stage.StartedAt = field.NewTime(tableName, "started_at")
	_stage.CreatedAt = field.NewTime(tableName, "created_at")
	_stage.UpdatedAt = field.NewTime(tableName, "updated_at")
	_stage.DeletedAt = field.NewField(tableName, "deleted_at")
	_stage.Points = stageManyToManyPoints{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Points", "models.Point"),
		Stages: struct {
			field.RelationField
			Points struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Points.Stages", "models.Stage"),
			Points: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Points.Stages.Points", "models.Point"),
			},
		},
	}

	_stage.fillFieldMap()

	return _stage
}

type stage struct {
	stageDo

	ALL       field.Asterisk
	ID        field.Uint64
	Title     field.String
	CreatorID field.Uint64
	JobID     field.Uint64
	StartedAt field.Time
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Points    stageManyToManyPoints

	fieldMap map[string]field.Expr
}

func (s stage) Table(newTableName string) *stage {
	s.stageDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s stage) As(alias string) *stage {
	s.stageDo.DO = *(s.stageDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *stage) updateTableName(table string) *stage {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewUint64(table, "id")
	s.Title = field.NewString(table, "title")
	s.CreatorID = field.NewUint64(table, "creator_id")
	s.JobID = field.NewUint64(table, "job_id")
	s.StartedAt = field.NewTime(table, "started_at")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")

	s.fillFieldMap()

	return s
}

func (s *stage) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *stage) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 9)
	s.fieldMap["id"] = s.ID
	s.fieldMap["title"] = s.Title
	s.fieldMap["creator_id"] = s.CreatorID
	s.fieldMap["job_id"] = s.JobID
	s.fieldMap["started_at"] = s.StartedAt
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt

}

func (s stage) clone(db *gorm.DB) stage {
	s.stageDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s stage) replaceDB(db *gorm.DB) stage {
	s.stageDo.ReplaceDB(db)
	return s
}

type stageManyToManyPoints struct {
	db *gorm.DB

	field.RelationField

	Stages struct {
		field.RelationField
		Points struct {
			field.RelationField
		}
	}
}

func (a stageManyToManyPoints) Where(conds ...field.Expr) *stageManyToManyPoints {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a stageManyToManyPoints) WithContext(ctx context.Context) *stageManyToManyPoints {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a stageManyToManyPoints) Session(session *gorm.Session) *stageManyToManyPoints {
	a.db = a.db.Session(session)
	return &a
}

func (a stageManyToManyPoints) Model(m *models.Stage) *stageManyToManyPointsTx {
	return &stageManyToManyPointsTx{a.db.Model(m).Association(a.Name())}
}

type stageManyToManyPointsTx struct{ tx *gorm.Association }

func (a stageManyToManyPointsTx) Find() (result []*models.Point, err error) {
	return result, a.tx.Find(&result)
}

func (a stageManyToManyPointsTx) Append(values ...*models.Point) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a stageManyToManyPointsTx) Replace(values ...*models.Point) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a stageManyToManyPointsTx) Delete(values ...*models.Point) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a stageManyToManyPointsTx) Clear() error {
	return a.tx.Clear()
}

func (a stageManyToManyPointsTx) Count() int64 {
	return a.tx.Count()
}

type stageDo struct{ gen.DO }

type IStageDo interface {
	gen.SubQuery
	Debug() IStageDo
	WithContext(ctx context.Context) IStageDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IStageDo
	WriteDB() IStageDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IStageDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStageDo
	Not(conds ...gen.Condition) IStageDo
	Or(conds ...gen.Condition) IStageDo
	Select(conds ...field.Expr) IStageDo
	Where(conds ...gen.Condition) IStageDo
	Order(conds ...field.Expr) IStageDo
	Distinct(cols ...field.Expr) IStageDo
	Omit(cols ...field.Expr) IStageDo
	Join(table schema.Tabler, on ...field.Expr) IStageDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStageDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStageDo
	Group(cols ...field.Expr) IStageDo
	Having(conds ...gen.Condition) IStageDo
	Limit(limit int) IStageDo
	Offset(offset int) IStageDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStageDo
	Unscoped() IStageDo
	Create(values ...*models.Stage) error
	CreateInBatches(values []*models.Stage, batchSize int) error
	Save(values ...*models.Stage) error
	First() (*models.Stage, error)
	Take() (*models.Stage, error)
	Last() (*models.Stage, error)
	Find() ([]*models.Stage, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Stage, err error)
	FindInBatches(result *[]*models.Stage, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Stage) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStageDo
	Assign(attrs ...field.AssignExpr) IStageDo
	Joins(fields ...field.RelationField) IStageDo
	Preload(fields ...field.RelationField) IStageDo
	FirstOrInit() (*models.Stage, error)
	FirstOrCreate() (*models.Stage, error)
	FindByPage(offset int, limit int) (result []*models.Stage, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStageDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s stageDo) Debug() IStageDo {
	return s.withDO(s.DO.Debug())
}

func (s stageDo) WithContext(ctx context.Context) IStageDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s stageDo) ReadDB() IStageDo {
	return s.Clauses(dbresolver.Read)
}

func (s stageDo) WriteDB() IStageDo {
	return s.Clauses(dbresolver.Write)
}

func (s stageDo) Session(config *gorm.Session) IStageDo {
	return s.withDO(s.DO.Session(config))
}

func (s stageDo) Clauses(conds ...clause.Expression) IStageDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s stageDo) Returning(value interface{}, columns ...string) IStageDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s stageDo) Not(conds ...gen.Condition) IStageDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s stageDo) Or(conds ...gen.Condition) IStageDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s stageDo) Select(conds ...field.Expr) IStageDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s stageDo) Where(conds ...gen.Condition) IStageDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s stageDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IStageDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s stageDo) Order(conds ...field.Expr) IStageDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s stageDo) Distinct(cols ...field.Expr) IStageDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s stageDo) Omit(cols ...field.Expr) IStageDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s stageDo) Join(table schema.Tabler, on ...field.Expr) IStageDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s stageDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStageDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s stageDo) RightJoin(table schema.Tabler, on ...field.Expr) IStageDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s stageDo) Group(cols ...field.Expr) IStageDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s stageDo) Having(conds ...gen.Condition) IStageDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s stageDo) Limit(limit int) IStageDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s stageDo) Offset(offset int) IStageDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s stageDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStageDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s stageDo) Unscoped() IStageDo {
	return s.withDO(s.DO.Unscoped())
}

func (s stageDo) Create(values ...*models.Stage) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s stageDo) CreateInBatches(values []*models.Stage, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s stageDo) Save(values ...*models.Stage) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s stageDo) First() (*models.Stage, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Stage), nil
	}
}

func (s stageDo) Take() (*models.Stage, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Stage), nil
	}
}

func (s stageDo) Last() (*models.Stage, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Stage), nil
	}
}

func (s stageDo) Find() ([]*models.Stage, error) {
	result, err := s.DO.Find()
	return result.([]*models.Stage), err
}

func (s stageDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Stage, err error) {
	buf := make([]*models.Stage, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s stageDo) FindInBatches(result *[]*models.Stage, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s stageDo) Attrs(attrs ...field.AssignExpr) IStageDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s stageDo) Assign(attrs ...field.AssignExpr) IStageDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s stageDo) Joins(fields ...field.RelationField) IStageDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s stageDo) Preload(fields ...field.RelationField) IStageDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s stageDo) FirstOrInit() (*models.Stage, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Stage), nil
	}
}

func (s stageDo) FirstOrCreate() (*models.Stage, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Stage), nil
	}
}

func (s stageDo) FindByPage(offset int, limit int) (result []*models.Stage, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s stageDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s stageDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s stageDo) Delete(models ...*models.Stage) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *stageDo) withDO(do gen.Dao) *stageDo {
	s.DO = *do.(*gen.DO)
	return s
}
