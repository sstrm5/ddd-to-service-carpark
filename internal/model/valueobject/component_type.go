package valueobject

import errs "dddcarpark/internal/errors"

type ComponentType string

var validComponentTypes = make(map[ComponentType]struct{})

func register(ct ComponentType) ComponentType {
	validComponentTypes[ct] = struct{}{}
	return ct
}

var (
	OilFilter      ComponentType = register(ComponentType("oil_filter"))       // масляный фильтр
	AirFilter      ComponentType = register(ComponentType("air_filter"))       // воздушный фильтр
	CabinFilter    ComponentType = register(ComponentType("cabin_filter"))     // салонный фильтр
	SparkPlug      ComponentType = register(ComponentType("spark_plug"))       // свеча зажигания
	TimingBelt     ComponentType = register(ComponentType("timing_belt"))      // ремень ГРМ
	BrakePadsFront ComponentType = register(ComponentType("brake_pads_front")) // передние тормозные колодки
	BrakePadsRear  ComponentType = register(ComponentType("brake_pads_rear"))  // задние тормозные колодки
	BrakeDiscFront ComponentType = register(ComponentType("brake_disc_front")) // передние тормозные диски
	BrakeDiscRear  ComponentType = register(ComponentType("brake_disc_rear"))  // задние тормозные диски
)

func NewComponentType(value string) (ComponentType, error) {
	componentType := ComponentType(value)
	if _, ok := validComponentTypes[componentType]; !ok {
		return "", errs.ErrInvalidComponentType
	}
	return componentType, nil
}
