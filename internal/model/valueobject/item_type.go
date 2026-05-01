package vo

import errs "dddcarpark/internal/errors"

type ItemType string

const (
	Work      ItemType = "work"
	Component ItemType = "component"
)

var validItemTypes = map[ItemType]struct{}{
	Work:      {},
	Component: {},
}

func NewItemType(value string) (ItemType, error) {
	itemType := ItemType(value)
	if _, ok := validItemTypes[itemType]; !ok {
		return "", errs.ErrInvalidCurrency
	}
	return itemType, nil
}
