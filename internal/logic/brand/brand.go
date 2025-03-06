package brand

import (
	"context"
	"database/sql"
	"errors"
	"goApp/internal/dao"
	"goApp/internal/model"
	"goApp/internal/model/entity"
	"goApp/internal/service"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type sBrand struct{}

var (
	cols = dao.Brand.Columns()
)

func New() *sBrand { return &sBrand{} }
func init()        { service.RegisterBrand(New()) }

func (s *sBrand) Create(ctx context.Context, in *model.BrandCreateInput) (err error) {
	if cnt, err := dao.Brand.Ctx(ctx).Where(cols.Name, in.Name).Count(); err != nil || cnt > 0 {
		return gerror.Newf("银行名称: %s 已存在", in.Name)
	}
	_, err = dao.Brand.Ctx(ctx).Insert(in)
	if err != nil {
		return gerror.New("创建银行失败")
	}
	return
}
func (s *sBrand) Update(ctx context.Context, id int64, in *model.BrandUpdateInput) (err error) {
	if cnt, err := dao.Brand.Ctx(ctx).Where(cols.Name, in.Name).WhereNot(cols.Id, id).Count(); err != nil || cnt > 0 {
		return gerror.Newf("银行名称: %s 已存在", in.Name)
	}
	if _, err = dao.Brand.Ctx(ctx).Where(cols.Id, id).Update(in); err != nil {
		return gerror.New("更新银行失败")
	}
	return
}
func (s *sBrand) Delete(ctx context.Context, id int64) (err error) {
	if cnt, err := dao.Brand.Ctx(ctx).Where(cols.Id, id).Count(); err != nil || cnt == 0 {
		return gerror.New("银行不存在")
	}
	if _, err := dao.Brand.Ctx(ctx).Where(cols.Id, id).Delete(); err != nil {
		return gerror.New("删除银行失败")
	}
	return
}
func (s *sBrand) GetOne(ctx context.Context, id int64) (out *model.BrandGetOneOutput, err error) {
	brand := entity.Brand{}
	if err = dao.Brand.Ctx(ctx).Where(cols.Id, id).Scan(&brand); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.New("银行不存在")
		}
		return nil, gerror.New("获取银行失败")
	}
	return &model.BrandGetOneOutput{Brand: &brand}, nil
}
func (s *sBrand) GetList(ctx context.Context, in *model.BrandGetListInput) (out *model.BrandGetListOutput, err error) {
	brands := make([]*entity.Brand, 0)
	total := 0
	search := strings.TrimSpace(in.Search)
	m := dao.Brand.Ctx(ctx)
	if search != "" {
		if id := gconv.Int64(search); id > 0 {
			m = m.Where(cols.Id, id)
		}
		m = m.WhereOr(cols.Name, "%"+search+"%")
	}
	if err = m.Page(in.PageNum, in.PageSize).OrderAsc(cols.Id).ScanAndCount(&brands, &total, false); err != nil {
		return nil, gerror.New("获取银行列表失败")
	}
	return &model.BrandGetListOutput{List: brands, Total: total}, nil
}
