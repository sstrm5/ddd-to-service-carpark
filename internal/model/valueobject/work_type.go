package vo

import errs "dddcarpark/internal/errors"

type WorkType string

const (
	EngineOilAndFilterChange    = WorkType("engineOilAndFilterChange")    // замена моторного масла и фильтров
	SuspensionDiagnostics       = WorkType("suspensionDiagnostics")       // диагностика ходовой части
	BrakePadAndRotorReplacement = WorkType("brakePadAndRotorReplacement") // замена тормозных колодок и дисков
	WheelAlignment              = WorkType("wheelAlignment")              // развал-схождение (регулировка углов установки колес)
	TimingBeltChainReplacement  = WorkType("TimingBeltChainReplacement")  // замена ремня или цепи ГРМ
	EngineComputerDiagnostics   = WorkType("engineComputerDiagnostics")   // компьютерная диагностика двигателя
	SparkPlugReplacement        = WorkType("sparkPlugReplacement")        // замена свечей зажигания
	AcSystemService             = WorkType("acSystemService")             // обслуживание системы кондиционирования
	ShockAbsorberReplacement    = WorkType("shockAbsorberReplacement")    // замена амортизаторов
	TireMountingAndBalancing    = WorkType("tireMountingAndBalancing")    // шиномонтаж и балансировка
)

var validWorkTypes = map[WorkType]struct{}{
	EngineOilAndFilterChange:    {},
	SuspensionDiagnostics:       {},
	BrakePadAndRotorReplacement: {},
	WheelAlignment:              {},
	TimingBeltChainReplacement:  {},
	EngineComputerDiagnostics:   {},
	SparkPlugReplacement:        {},
	AcSystemService:             {},
	ShockAbsorberReplacement:    {},
	TireMountingAndBalancing:    {},
}

func NewWorkType(value string) (WorkType, error) {
	workType := WorkType(value)
	if _, ok := validWorkTypes[workType]; !ok {
		return "", errs.ErrInvalidWorkType
	}
	return workType, nil
}
