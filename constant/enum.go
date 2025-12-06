package constant

type RefTable string
type TransactionType int64

type WsType string
type EventStatus string

const (
	// Ref Table
	REF_TABLE_ITEM        RefTable = "item"
	REF_TABLE_ITEMVARIANT RefTable = "itemvariant"
	REF_TABLE_ADDON       RefTable = "addon"
	REF_TABLE_USER        RefTable = "user"
)

const (
	TRANSACTION_TYPE_DEBIT  TransactionType = 1
	TRANSACTION_TYPE_KREDIT TransactionType = -1
)

const (
	WS_TYPE_GET_EVENT  WsType = "GET_EVENT"
	WS_TYPE_DATA_EVENT WsType = "DATA_EVENT"
	WS_TYPE_REFETCH    WsType = "REFETCH"
)

const (
	EVENT_STATUS_HOLD    EventStatus = "HOLD"
	EVENT_STATUS_CONFIRM EventStatus = "CONFIRM"
)
