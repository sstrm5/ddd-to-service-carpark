package entity

import vo "dddcarpark/internal/model/valueobject"

type OrderItem struct {
	ID          string
	Description vo.Description
	Type        vo.ItemType
	Cost        vo.Money
}
