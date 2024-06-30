package dto

import "time"

type CreateMemberRequest struct {
	Email    string
	Name     string
	Age      int
	Gender   string
	Password string
}

type CreateTransactionRequest struct {
	MemberIDSender   string
	MemberIDReceiver string
	SentAt           time.Time
	Amount           int
}
