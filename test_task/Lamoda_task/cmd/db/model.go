package db

type Product struct {
	Articular int `json:"articular" db:"articular"`
	Quantity  int `json:"quantity" db:"quantity"`
}

type CancelReservation struct {
	Articular int    `json:"articular" db:"articular"`
	Quantity  int    `json:"quantity" db:"quantity"`
	Storage   string `json:"storage" db:"storage"`
}
