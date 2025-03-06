package bank

import (
	"context"
	"database/sql"
	"errors"
	"goApp/internal/dao"
	"goApp/internal/model"
	"goApp/internal/model/entity"
	"strings"

	"goApp/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

var (
	cols = dao.Bank.Columns()
	/*cTable = dao.Bank.Table()*/
)

type sBank struct{}

func New() *sBank { return &sBank{} }

func init() { service.RegisterBank(New()) }

func (c *sBank) Create(ctx context.Context, in *model.BankCreateInput) (err error) {
	return
}

func (c *sBank) Update(ctx context.Context, id int64, in *model.BankUpdateInput) (err error) {
	return
}

func (c *sBank) Delete(ctx context.Context, id int64) (err error) {
	return
}

func (c *sBank) GetList(ctx context.Context, in *model.BankGetListInput) (out *model.BankGetListOutput, err error) {
	search := strings.TrimSpace(in.Search)
	if search != "" {
	}
	return
}

func (c *sBank) GetOne(ctx context.Context, id int64) (out *model.BankGetOneOutput, err error) {
	bank := new(entity.Bank)
	if err = dao.Bank.Ctx(ctx).Where(cols.Id, id).Scan(bank); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, gerror.New("银行信息不存在")
		}
		return nil, gerror.New("查询银行信息失败")
	}
	return &model.BankGetOneOutput{Bank: bank}, nil
}
