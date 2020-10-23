package dwolla

const (
	Deffered  PaymentStatus = "deferred"
	Pending   PaymentStatus = "pending"
	Cancelled PaymentStatus = "cancelled"
)

type PaymentStatus string

const (
	location = "Location"
)

type BusinessType string

const (
	LLC                BusinessType = "llc"
	Patnership         BusinessType = "partnership"
	Corporation        BusinessType = "corporation"
	SoleProprietorship BusinessType = "soleProprietorship"
)

type CustomerType string

const (
	Perosnal CustomerType = "personal"
	Business CustomerType = "business"
)
