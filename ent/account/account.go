// Code generated by ent, DO NOT EDIT.

package account

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldPort holds the string denoting the port field in the database.
	FieldPort = "port"
	// FieldProtocol holds the string denoting the protocol field in the database.
	FieldProtocol = "protocol"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldPrivateKey holds the string denoting the private_key field in the database.
	FieldPrivateKey = "private_key"
	// FieldPrivateKeyPassword holds the string denoting the private_key_password field in the database.
	FieldPrivateKeyPassword = "private_key_password"
	// FieldDeviceID holds the string denoting the device_id field in the database.
	FieldDeviceID = "device_id"
	// FieldDepartmentID holds the string denoting the department_id field in the database.
	FieldDepartmentID = "department_id"
	// EdgeDevices holds the string denoting the devices edge name in mutations.
	EdgeDevices = "devices"
	// EdgeDepartments holds the string denoting the departments edge name in mutations.
	EdgeDepartments = "departments"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// DevicesTable is the table that holds the devices relation/edge.
	DevicesTable = "accounts"
	// DevicesInverseTable is the table name for the Device entity.
	// It exists in this package in order to avoid circular dependency with the "device" package.
	DevicesInverseTable = "devices"
	// DevicesColumn is the table column denoting the devices relation/edge.
	DevicesColumn = "device_id"
	// DepartmentsTable is the table that holds the departments relation/edge.
	DepartmentsTable = "accounts"
	// DepartmentsInverseTable is the table name for the Department entity.
	// It exists in this package in order to avoid circular dependency with the "department" package.
	DepartmentsInverseTable = "departments"
	// DepartmentsColumn is the table column denoting the departments relation/edge.
	DepartmentsColumn = "department_id"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldUsername,
	FieldPort,
	FieldProtocol,
	FieldPassword,
	FieldPrivateKey,
	FieldPrivateKeyPassword,
	FieldDeviceID,
	FieldDepartmentID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// PortValidator is a validator for the "port" field. It is called by the builders before save.
	PortValidator func(uint32) error
	// ProtocolValidator is a validator for the "protocol" field. It is called by the builders before save.
	ProtocolValidator func(uint8) error
	// PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	PasswordValidator func(string) error
	// PrivateKeyValidator is a validator for the "private_key" field. It is called by the builders before save.
	PrivateKeyValidator func(string) error
	// PrivateKeyPasswordValidator is a validator for the "private_key_password" field. It is called by the builders before save.
	PrivateKeyPasswordValidator func(string) error
	// DeviceIDValidator is a validator for the "device_id" field. It is called by the builders before save.
	DeviceIDValidator func(uint64) error
	// DepartmentIDValidator is a validator for the "department_id" field. It is called by the builders before save.
	DepartmentIDValidator func(uint64) error
)

// OrderOption defines the ordering options for the Account queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUsername orders the results by the username field.
func ByUsername(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUsername, opts...).ToFunc()
}

// ByPort orders the results by the port field.
func ByPort(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPort, opts...).ToFunc()
}

// ByProtocol orders the results by the protocol field.
func ByProtocol(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProtocol, opts...).ToFunc()
}

// ByPassword orders the results by the password field.
func ByPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPassword, opts...).ToFunc()
}

// ByPrivateKey orders the results by the private_key field.
func ByPrivateKey(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivateKey, opts...).ToFunc()
}

// ByPrivateKeyPassword orders the results by the private_key_password field.
func ByPrivateKeyPassword(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrivateKeyPassword, opts...).ToFunc()
}

// ByDeviceID orders the results by the device_id field.
func ByDeviceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeviceID, opts...).ToFunc()
}

// ByDepartmentID orders the results by the department_id field.
func ByDepartmentID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDepartmentID, opts...).ToFunc()
}

// ByDevicesField orders the results by devices field.
func ByDevicesField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDevicesStep(), sql.OrderByField(field, opts...))
	}
}

// ByDepartmentsField orders the results by departments field.
func ByDepartmentsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newDepartmentsStep(), sql.OrderByField(field, opts...))
	}
}
func newDevicesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DevicesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DevicesTable, DevicesColumn),
	)
}
func newDepartmentsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(DepartmentsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, DepartmentsTable, DepartmentsColumn),
	)
}
