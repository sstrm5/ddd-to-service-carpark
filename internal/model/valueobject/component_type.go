package vo

import errs "dddcarpark/internal/errors"

type ComponentType string

const (
	OilFilter      ComponentType = ("oil_filter")       // масляный фильтр
	AirFilter      ComponentType = ("air_filter")       // воздушный фильтр
	CabinFilter    ComponentType = ("cabin_filter")     // салонный фильтр
	SparkPlug      ComponentType = ("spark_plug")       // свеча зажигания
	TimingBelt     ComponentType = ("timing_belt")      // ремень ГРМ
	BrakePadsFront ComponentType = ("brake_pads_front") // передние тормозные колодки
	BrakePadsRear  ComponentType = ("brake_pads_rear")  // задние тормозные колодки
	BrakeDiscFront ComponentType = ("brake_disc_front") // передние тормозные диски
	BrakeDiscRear  ComponentType = ("brake_disc_rear")  // задние тормозные диски
)

var validComponentTypes = map[ComponentType]struct{}{
	OilFilter:      {},
	AirFilter:      {},
	CabinFilter:    {},
	SparkPlug:      {},
	TimingBelt:     {},
	BrakePadsFront: {},
	BrakePadsRear:  {},
	BrakeDiscFront: {},
	BrakeDiscRear:  {},
}

func NewComponentType(value string) (ComponentType, error) {
	componentType := ComponentType(value)
	if _, ok := validComponentTypes[componentType]; !ok {
		return "", errs.ErrInvalidComponentType
	}
	return componentType, nil
}
