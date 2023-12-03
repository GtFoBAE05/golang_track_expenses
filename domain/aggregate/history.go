package aggregate

import "golang_track_expense/domain/entity"

type History struct {
	user entity.User

	trxs []entity.Transaction
}

func NewHistory(user entity.User, trx entity.Transaction) (history History) {
	history.user = user
	history.trxs = append(history.trxs, trx)
	return
}

func (h *History) AddTrx(trx entity.Transaction) *History {
	h.trxs = append(h.trxs, trx)
	return h
}
