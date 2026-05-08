package entity

import (
	errs "dddcarpark/internal/errors"
	vo "dddcarpark/internal/model/valueobject"
)

type OrderLine struct {
	ID          string
	Type        vo.ItemType
	Cost        vo.Money
	Description *vo.Description

	// для Type="component"
	Component *InstalledComponent

	// для Type="work"
	WorkType *vo.WorkType
}

func NewOrderLine(
	itemType vo.ItemType,
	cost vo.Money,
	description *vo.Description,
	component *InstalledComponent,
	workType *vo.WorkType,
) (OrderLine, error) {
	if itemType == vo.Work && component != nil {
		return OrderLine{}, errs.ErrMismatchTypeOrderLine
	}
	if itemType == vo.Component && component == nil {
		return OrderLine{}, errs.ErrMismatchTypeOrderLine
	}
	return OrderLine{
		Type:        itemType,
		Cost:        cost,
		Description: description,
		Component:   component,
	}, nil
}
