package valueobject

// Pair is currency pair
type Pair int

// BtcJpy is Bitcoin & Japanese Yen
const (
	BtcJpy Pair = iota
)

type User struct {
	ID       uint64
	Username string
}

// Ticker is basic data type of trade
type Ticker struct {
	Sell      string
	Buy       string
	High      string
	Low       string
	Last      string
	Vol       string
	Timestamp string
}

// Payment is payment info of domain.Order
type Payment struct {
}
