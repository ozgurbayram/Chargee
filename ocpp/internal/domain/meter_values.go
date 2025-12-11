package domain

type MeterValuesRequest struct {
	ConnectorId   int          `json:"connectorId"`
	TransactionId int          `json:"transactionId,omitempty"`
	MeterValue    []MeterValue `json:"meterValue"`
}

type MeterValue struct {
	Timestamp    string         `json:"timestamp"`
	SampledValue []SampledValue `json:"sampledValue"`
}

type SampledValue struct {
	Value     string         `json:"value"`
	Context   ReadingContext `json:"context,omitempty"`
	Format    ValueFormat    `json:"format,omitempty"`
	Measurand Measurand      `json:"measurand,omitempty"`
	Phase     Phase          `json:"phase,omitempty"`
	Location  Location       `json:"location,omitempty"`
	Unit      UnitOfMeasure  `json:"unit,omitempty"`
}

type ReadingContext string

const (
	ReadingContextInterruptionBegin ReadingContext = "Interruption.Begin"
	ReadingContextInterruptionEnd   ReadingContext = "Interruption.End"
	ReadingContextSampleClock       ReadingContext = "Sample.Clock"
	ReadingContextSamplePeriodic    ReadingContext = "Sample.Periodic"
	ReadingContextTransactionBegin  ReadingContext = "Transaction.Begin"
	ReadingContextTransactionEnd    ReadingContext = "Transaction.End"
	ReadingContextTrigger           ReadingContext = "Trigger"
	ReadingContextOther             ReadingContext = "Other"
)

type ValueFormat string

const (
	ValueFormatRaw        ValueFormat = "Raw"
	ValueFormatSignedData ValueFormat = "SignedData"
)

type Measurand string

const (
	MeasurandEnergyActiveExportRegister   Measurand = "Energy.Active.Export.Register"
	MeasurandEnergyActiveImportRegister   Measurand = "Energy.Active.Import.Register"
	MeasurandEnergyReactiveExportRegister Measurand = "Energy.Reactive.Export.Register"
	MeasurandEnergyReactiveImportRegister Measurand = "Energy.Reactive.Import.Register"
	MeasurandEnergyActiveExportInterval   Measurand = "Energy.Active.Export.Interval"
	MeasurandEnergyActiveImportInterval   Measurand = "Energy.Active.Import.Interval"
	MeasurandEnergyReactiveExportInterval Measurand = "Energy.Reactive.Export.Interval"
	MeasurandEnergyReactiveImportInterval Measurand = "Energy.Reactive.Import.Interval"
	MeasurandPowerActiveExport            Measurand = "Power.Active.Export"
	MeasurandPowerActiveImport            Measurand = "Power.Active.Import"
	MeasurandPowerOffered                 Measurand = "Power.Offered"
	MeasurandPowerReactiveExport          Measurand = "Power.Reactive.Export"
	MeasurandPowerReactiveImport          Measurand = "Power.Reactive.Import"
	MeasurandPowerFactor                  Measurand = "Power.Factor"
	MeasurandCurrentImport                Measurand = "Current.Import"
	MeasurandCurrentExport                Measurand = "Current.Export"
	MeasurandCurrentOffered               Measurand = "Current.Offered"
	MeasurandVoltage                      Measurand = "Voltage"
	MeasurandFrequency                    Measurand = "Frequency"
	MeasurandTemperature                  Measurand = "Temperature"
	MeasurandSoC                          Measurand = "SoC"
	MeasurandRPM                          Measurand = "RPM"
)

type Phase string

const (
	PhaseL1   Phase = "L1"
	PhaseL2   Phase = "L2"
	PhaseL3   Phase = "L3"
	PhaseN    Phase = "N"
	PhaseL1N  Phase = "L1-N"
	PhaseL2N  Phase = "L2-N"
	PhaseL3N  Phase = "L3-N"
	PhaseL1L2 Phase = "L1-L2"
	PhaseL2L3 Phase = "L2-L3"
	PhaseL3L1 Phase = "L3-L1"
)

type Location string

const (
	LocationCable  Location = "Cable"
	LocationEV     Location = "EV"
	LocationInlet  Location = "Inlet"
	LocationOutlet Location = "Outlet"
	LocationBody   Location = "Body"
)

type UnitOfMeasure string

const (
	UnitWh         UnitOfMeasure = "Wh"
	UnitKWh        UnitOfMeasure = "kWh"
	UnitVarh       UnitOfMeasure = "varh"
	UnitKVarh      UnitOfMeasure = "kvarh"
	UnitW          UnitOfMeasure = "W"
	UnitKW         UnitOfMeasure = "kW"
	UnitVA         UnitOfMeasure = "VA"
	UnitKVA        UnitOfMeasure = "kVA"
	UnitVar        UnitOfMeasure = "var"
	UnitKVar       UnitOfMeasure = "kvar"
	UnitA          UnitOfMeasure = "A"
	UnitV          UnitOfMeasure = "V"
	UnitK          UnitOfMeasure = "K"
	UnitCelsius    UnitOfMeasure = "Celsius"
	UnitFahrenheit UnitOfMeasure = "Fahrenheit"
	UnitPercent    UnitOfMeasure = "Percent"
)

type MeterValuesResponse struct {
}
