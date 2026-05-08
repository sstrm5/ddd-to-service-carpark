package entity

import (
	errs "dddcarpark/internal/errors"
	vo "dddcarpark/internal/model/valueobject"
	"time"
)

type MaintenanceOrder struct {
	ID             string
	VehicleID      string
	CreatedAt      time.Time
	ServiceMileage vo.Mileage
	Status         vo.StatusOrder
	Currency       vo.Currency
	OrderLines     []OrderLine
}

func NewMaintenanceOrder(
	vehicleID string,
	lastServiceMileage vo.Mileage,
	serviceMileage vo.Mileage,
	date time.Time,
	currency vo.Currency,
) (MaintenanceOrder, error) {
	if vehicleID == "" {
		return MaintenanceOrder{}, errs.ErrInvalidVehicleID
	}
	if serviceMileage.LessThan(lastServiceMileage) {
		return MaintenanceOrder{}, errs.ErrMileageCannotDecrease
	}
	if date.IsZero() || date.After(time.Now()) {
		return MaintenanceOrder{}, errs.ErrInvalidDate
	}
	return MaintenanceOrder{
		VehicleID:      vehicleID,
		CreatedAt:      date,
		ServiceMileage: serviceMileage,
		Status:         vo.Draft,
		Currency:       currency,
	}, nil
}

func (mo *MaintenanceOrder) AddOrderLine(orderLine OrderLine) error {
	if mo.Status != vo.Draft {
		return errs.ErrInvalidOrderStatus
	}

	if mo.Currency != orderLine.Cost.Currency {
		return errs.ErrInvalidCurrency
	}

	// TODO: убрать повторение
	if orderLine.Type == vo.Work {
		for _, o := range mo.OrderLines {
			if o.Type == vo.Work && o.WorkType == orderLine.WorkType {
				return errs.ErrOrderLineAlreadyExists
			}
		}
		if orderLine.WorkType == nil {
			return errs.ErrInvalidWorkType
		}
	}
	if orderLine.Type == vo.Component {
		for _, o := range mo.OrderLines {
			if o.Type == vo.Component && orderLine.Component == o.Component {
				return errs.ErrOrderLineAlreadyExists
			}
		}
		if orderLine.Component == nil {
			return errs.ErrInvalidComponent
		}
	}

	mo.OrderLines = append(mo.OrderLines, orderLine)
	return nil
}

func (mo *MaintenanceOrder) ReplaceOrderLine(newOrderLine OrderLine) error {
	if mo.Status != vo.Draft {
		return errs.ErrInvalidOrderStatus
	}

	if mo.Currency != newOrderLine.Cost.Currency {
		return errs.ErrInvalidCurrency
	}

	if newOrderLine.Type == vo.Work {
		if newOrderLine.WorkType == nil {
			return errs.ErrInvalidWorkType
		}
		for i, orderLine := range mo.OrderLines {
			if orderLine.Type == vo.Work && orderLine.WorkType == newOrderLine.WorkType {
				mo.OrderLines[i] = newOrderLine
				break
			}
		}
	}
	if newOrderLine.Type == vo.Component {
		if newOrderLine.Component == nil {
			return errs.ErrInvalidWorkType
		}
		for i, orderLine := range mo.OrderLines {
			if orderLine.Type == vo.Component && orderLine.Component.Type == newOrderLine.Component.Type {
				mo.OrderLines[i] = newOrderLine
			}
		}
	}
	return nil
}

func (mo *MaintenanceOrder) Complete(vehicle Vehicle) error {
	if mo.VehicleID != vehicle.ID {
		return errs.ErrInvalidVehicleID
	}

	for _, orderLine := range mo.OrderLines {
		if orderLine.Type == vo.Component {
			if err := vehicle.InstallComponent(*orderLine.Component); err != nil {
				return err
			}
		}
	}

	mo.Status = vo.Completed
	return nil
}
