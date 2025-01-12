// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "username", Type: field.TypeString, Size: 16, Comment: "account username|账号名称"},
		{Name: "port", Type: field.TypeUint32, Comment: "account port|端口"},
		{Name: "protocol", Type: field.TypeUint8, Comment: "protocol of the this account.|账号协议"},
		{Name: "password", Type: field.TypeString, Nullable: true, Size: 128, Comment: "account password|账号密码"},
		{Name: "private_key", Type: field.TypeString, Nullable: true, Size: 4096, Comment: "private_key of the this account.|账号私钥"},
		{Name: "private_key_password", Type: field.TypeString, Nullable: true, Size: 4096, Comment: "private_key password of the this account.|私钥口令"},
		{Name: "device_id", Type: field.TypeUint64, Comment: "account belong to|账号所属设备"},
		{Name: "department_id", Type: field.TypeUint64, Comment: "account belong to|账号所属部门"},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "accounts_devices_devices",
				Columns:    []*schema.Column{AccountsColumns[9]},
				RefColumns: []*schema.Column{DevicesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "accounts_departments_departments",
				Columns:    []*schema.Column{AccountsColumns[10]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// DepartmentsColumns holds the columns for the "departments" table.
	DepartmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 64, Comment: "The name of the department.|部门名称"},
		{Name: "area", Type: field.TypeString, Size: 128, Comment: "The area where the department is located.|部门所在地区"},
		{Name: "description", Type: field.TypeString, Size: 128, Comment: "Description of the department.|部门描述"},
		{Name: "parent_department_id", Type: field.TypeUint64, Comment: "parent department ID.|父亲部门id"},
		{Name: "parent_departments", Type: field.TypeString, Comment: "Comma-separated list of parent department IDs in ascending order.|上级部门集合逗号分隔升序排列"},
	}
	// DepartmentsTable holds the schema information for the "departments" table.
	DepartmentsTable = &schema.Table{
		Name:       "departments",
		Columns:    DepartmentsColumns,
		PrimaryKey: []*schema.Column{DepartmentsColumns[0]},
	}
	// DevicesColumns holds the columns for the "devices" table.
	DevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 128, Comment: "The name of the device.|设备名称"},
		{Name: "host", Type: field.TypeString, Size: 64, Comment: "login host|设备地址"},
		{Name: "type", Type: field.TypeString, Size: 64, Comment: "type of the device.|设备类型"},
		{Name: "description", Type: field.TypeString, Size: 128, Comment: "Description of the device.|设备描述"},
		{Name: "department_id", Type: field.TypeUint64, Comment: "ID of the device's department.|设备所属部门"},
	}
	// DevicesTable holds the schema information for the "devices" table.
	DevicesTable = &schema.Table{
		Name:       "devices",
		Columns:    DevicesColumns,
		PrimaryKey: []*schema.Column{DevicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "devices_departments_departments",
				Columns:    []*schema.Column{DevicesColumns[7]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LableTreesColumns holds the columns for the "lable_trees" table.
	LableTreesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Size: 64, Comment: "lable name|标签名称"},
		{Name: "type", Type: field.TypeUint, Comment: "lable type|标签类型（分组标签、控制标签、授权标签）"},
		{Name: "icon", Type: field.TypeString, Size: 32, Comment: "lable icon|标签图标"},
		{Name: "content", Type: field.TypeString, Size: 1024, Comment: "lable content|标签内容"},
		{Name: "ownership", Type: field.TypeUint8, Comment: "lable ownership Level (Department Level/User Level)|标签所有权级别（部门级别/用户级别）"},
		{Name: "owner_id", Type: field.TypeUint64, Comment: "lable owner,user ID,dept ID|标签所属者,用户ID,部门ID"},
		{Name: "parent_id", Type: field.TypeUint64, Comment: "parent lable id|父标签id"},
		{Name: "description", Type: field.TypeString, Size: 256, Comment: "label description|标签描述"},
		{Name: "target_type", Type: field.TypeUint16, Comment: "lable target type|标签目标类型"},
		{Name: "parents", Type: field.TypeString, Comment: "parent lables id,split by ',',lable tree deep cannot gather than 32|父标签id集合升序排列,逗号分隔,限制最多不可超过64级"},
		{Name: "inherit", Type: field.TypeBool, Comment: "child lable can inherit parents|标签是否可以继承"},
	}
	// LableTreesTable holds the schema information for the "lable_trees" table.
	LableTreesTable = &schema.Table{
		Name:       "lable_trees",
		Columns:    LableTreesColumns,
		PrimaryKey: []*schema.Column{LableTreesColumns[0]},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString, Unique: true, Size: 32, Comment: "The name of the role.|角色名称"},
		{Name: "weight", Type: field.TypeInt, Comment: "The weight of the role. Smaller values indicate higher priority.|角色优先级，值越小优先级越高"},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "account", Type: field.TypeString, Unique: true, Size: 64, Comment: "User account.|用户账号"},
		{Name: "name", Type: field.TypeString, Size: 64, Comment: "User name.|用户姓名"},
		{Name: "password", Type: field.TypeString, Size: 128, Comment: "User password.|用户密码"},
		{Name: "status", Type: field.TypeBool, Comment: "User status (enabled(true) or disabled(false)).|用户账号状态", Default: true},
		{Name: "description", Type: field.TypeString, Nullable: true, Size: 128, Comment: "User description.|用户账号描述"},
		{Name: "email", Type: field.TypeString, Nullable: true, Size: 64, Comment: "User email.|用户邮箱"},
		{Name: "phone_number", Type: field.TypeString, Nullable: true, Size: 32, Comment: "User phone number.|用户手机号码"},
		{Name: "department_id", Type: field.TypeUint64, Nullable: true, Comment: "ID of the user's department.|用户所属部门"},
		{Name: "role_id", Type: field.TypeUint64, Nullable: true, Comment: "ID of the user's role.|用户所属角色"},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "users_departments_departments",
				Columns:    []*schema.Column{UsersColumns[10]},
				RefColumns: []*schema.Column{DepartmentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "users_roles_roles",
				Columns:    []*schema.Column{UsersColumns[11]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		DepartmentsTable,
		DevicesTable,
		LableTreesTable,
		RolesTable,
		UsersTable,
	}
)

func init() {
	AccountsTable.ForeignKeys[0].RefTable = DevicesTable
	AccountsTable.ForeignKeys[1].RefTable = DepartmentsTable
	DevicesTable.ForeignKeys[0].RefTable = DepartmentsTable
	UsersTable.ForeignKeys[0].RefTable = DepartmentsTable
	UsersTable.ForeignKeys[1].RefTable = RolesTable
}
