package cns

const (
	AppDevelopmentEnv = "development"
	AppProductionEnv  = "production"
)

const (
	NilString = ""
	Space     = ""
	NewLine   = "\n"
)

var (
	ByteNewLine = []byte{'\n'}
	ByteSpace   = []byte{' '}
)

// Postgres constants

const (
	TodoItemsTable    = "todo_items"
	TodoItemsID       = "id"
	TodoItemsTitle    = "title"
	TodoItemsDesc     = "description"
	TodoItemDone      = "done"
	TodoItemCreatedAt = "created_at"
	TodoItemDeletedAt = "deleted_at"

	TodoListTable = "todo_list"
	TodoListID    = "id"
	TodoListTitle = "title"
	TodoListDesc  = "description"

	ListItemsTable  = "list_items"
	ListItemsListID = "list_id"
	ListItemsItemID = "item_id"
)
