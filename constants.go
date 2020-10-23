package dwolla

const (
	Deffered  PaymentStatus = "deferred"
	Pending   PaymentStatus = "pending"
	Cancelled PaymentStatus = "cancelled"
)

type PaymentStatus string
