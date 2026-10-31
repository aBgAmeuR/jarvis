package supervision

type Disk struct {
	Device       string
	Model        string
	Serial       string
	Protocol     string
	Rotational   bool
	HealthPassed bool
	Standby      bool
	// The figures below are nil when the disk does not report them.
	TemperatureCelsius *float64
	PowerOnHours       *int64
	NVMe               *NVMeHealth
	ATA                *ATAHealth
}

type NVMeHealth struct {
	PercentageUsed        float64
	AvailableSparePercent float64
	CriticalWarning       int64
	MediaErrors           int64
}

type ATAHealth struct {
	ReallocatedSectors   int64
	PendingSectors       int64
	UncorrectableSectors int64
	UDMACRCErrors        int64
}
