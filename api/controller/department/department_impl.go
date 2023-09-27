package department

import (
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"sectran/api/common"
	"sectran/api/model"
	"strconv"
)

func listDepartmentImpl(c *gin.Context) (error, []model.DepartmentModel, int) {
	var tableList []model.DepartmentModel
	var total int
	Db := common.Db
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	//分页
	if page > 0 && pageSize > 0 {
		Db = Db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
	err := Db.Find(&tableList).Offset(-1).Limit(-1).Count(&total).Error
	return err, tableList, total
}

func addDepartmentImpl(p departmentParameter) error {
	Db := common.Db
	fmt.Printf("%v\n", Db)

	var B = model.DepartmentModel{
		Id:       uuid.Must(uuid.NewV4(), nil).String(),
		Name:     p.Name,
		Describe: p.Describe,
	}
	err := Db.Create(&B).Error
	return err
}

func redactDepartmentImpl(p RedactDepartmentParameter) error {
	Db := common.Db
	var tableList []model.DepartmentModel
	fmt.Printf("%v\n", Db)
	err := Db.Model(&tableList).Where("id = ?", p.Id).Updates(map[string]interface{}{"name": p.Name, "describe": p.Describe}).Error
	return err
}

func deleteDepartmentImpl(p common.DeleteDto) error {
	Db := common.Db
	var tableList []model.DepartmentModel
	err := Db.Where("id = ?", p.Id).Delete(&tableList).Error
	return err
}
