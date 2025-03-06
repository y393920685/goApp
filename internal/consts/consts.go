package consts

type Status string
type Active string
type BankRecordType string
type BankRecordTypeText string

const (
	StatusEnable  Status = "1"
	StatusDisable Status = "0"
)

const (
	ActiveNo  Active = "0"
	ActiveYes Active = "1"
)

const (
	BankRecordTypeQCYE BankRecordType = "QCYE" //期初余额
	BankRecordTypeCG   BankRecordType = "CG"   //采购
)
const (
	BankRecordTypeTextQCYE BankRecordTypeText = "期初余额"
	BankRecordTypeTextCG   BankRecordTypeText = "采购"
)
