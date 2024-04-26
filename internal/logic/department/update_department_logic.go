package department

import (
	"context"

	"sectran_admin/ent"
	"sectran_admin/ent/department"
	"sectran_admin/internal/svc"
	"sectran_admin/internal/types"
	"sectran_admin/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDepartmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDepartmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDepartmentLogic {
	return &UpdateDepartmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDepartmentLogic) UpdateDepartment(req *types.DepartmentInfo) (*types.BaseMsgResp, error) {
	//查询当前主体的部门、获取到他父亲部门的部门前缀
	domain := l.ctx.Value("request_domain").((*ent.User))
	domainParentDepartments, err := l.svcCtx.DB.Department.Query().
		Where(department.ID(domain.DepartmentID)).
		Select(department.FieldParentDepartments).String(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//查询目标的部门
	data, err := l.svcCtx.DB.Department.Get(l.ctx, *req.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	//判断当前账号是否对待操作部门存在访问权限
	if _, err := DomainDeptAccessed(l.ctx, l.svcCtx, domainParentDepartments, data.ParentDepartments); err != nil {
		return nil, err
	}

	err = l.svcCtx.DB.Department.UpdateOneID(*req.Id).
		SetNotNilName(req.Name).
		SetNotNilArea(req.Area).
		SetNotNilDescription(req.Description).
		// 不允许修改部门的上级部门
		// SetNotNilParentDepartmentID(req.ParentDepartmentId).
		// SetNotNilParentDepartments(req.ParentDepartments).
		Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, req)
	}

	return &types.BaseMsgResp{Msg: l.svcCtx.Trans.Trans(l.ctx, i18n.UpdateSuccess)}, nil
}
