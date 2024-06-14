package entity

type Model interface {
	TableName() string
}

type Suffix interface {
	ToString() string
}
