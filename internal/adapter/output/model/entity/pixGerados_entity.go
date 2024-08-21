package entity

type PixGerado struct {
	ID        int
	Value     float64
	TxID      string
	SessionID int
	Session   Session
}
