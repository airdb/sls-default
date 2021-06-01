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
