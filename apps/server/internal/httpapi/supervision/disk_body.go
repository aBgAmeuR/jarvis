package supervision

import domainsupervision "jarvis/server/internal/domain/supervision"

type nvmeHealthBody struct {
	PercentageUsed        float64 `json:"percentageUsed" doc:"Wear reported by the drive, in percent." example:"1"`
	AvailableSparePercent float64 `json:"availableSparePercent" doc:"Spare capacity left, in percent." example:"100"`
	CriticalWarning       int64   `json:"criticalWarning" doc:"Critical warning bit field, 0 when none." example:"0"`
	MediaErrors           int64   `json:"mediaErrors" doc:"Unrecovered data integrity errors." example:"0"`
}

type ataHealthBody struct {
	ReallocatedSectors   int64 `json:"reallocatedSectors" doc:"Reallocated sectors." example:"0"`
	PendingSectors       int64 `json:"pendingSectors" doc:"Sectors waiting to be reallocated." example:"0"`
	UncorrectableSectors int64 `json:"uncorrectableSectors" doc:"Uncorrectable sectors found offline." example:"0"`
	UDMACRCErrors        int64 `json:"udmaCrcErrors" doc:"Interface CRC errors." example:"0"`
}

type diskBody struct {
	Device             string          `json:"device" doc:"Device path." example:"/dev/sda"`
	Model              string          `json:"model" doc:"Drive model." example:"WDC WD40EFPX-68C6CN0"`
	Serial             string          `json:"serial" doc:"Serial number." example:"WD-WX32D855YNY2"`
	Protocol           string          `json:"protocol" doc:"Interface protocol." example:"ATA"`
	Rotational         bool            `json:"rotational" doc:"Whether the drive has spinning platters." example:"true"`
	HealthPassed       bool            `json:"healthPassed" doc:"Result of the SMART overall health test." example:"true"`
	Standby            bool            `json:"standby" doc:"Whether the drive is asleep." example:"false"`
	TemperatureCelsius *float64        `json:"temperatureCelsius,omitempty" doc:"Drive temperature. Absent when not reported." example:"26"`
	PowerOnHours       *int64          `json:"powerOnHours,omitempty" doc:"Hours powered on. Absent when not reported." example:"7533"`
	NVMe               *nvmeHealthBody `json:"nvme,omitempty" doc:"NVMe health figures. Absent on other drives."`
	ATA                *ataHealthBody  `json:"ata,omitempty" doc:"ATA health figures. Absent on other drives."`
}

func newDiskBody(d domainsupervision.Disk) diskBody {
	body := diskBody{
		Device:             d.Device,
		Model:              d.Model,
		Serial:             d.Serial,
		Protocol:           d.Protocol,
		Rotational:         d.Rotational,
		HealthPassed:       d.HealthPassed,
		Standby:            d.Standby,
		TemperatureCelsius: d.TemperatureCelsius,
		PowerOnHours:       d.PowerOnHours,
	}
	if d.NVMe != nil {
		body.NVMe = &nvmeHealthBody{
			PercentageUsed:        d.NVMe.PercentageUsed,
			AvailableSparePercent: d.NVMe.AvailableSparePercent,
			CriticalWarning:       d.NVMe.CriticalWarning,
			MediaErrors:           d.NVMe.MediaErrors,
		}
	}
	if d.ATA != nil {
		body.ATA = &ataHealthBody{
			ReallocatedSectors:   d.ATA.ReallocatedSectors,
			PendingSectors:       d.ATA.PendingSectors,
			UncorrectableSectors: d.ATA.UncorrectableSectors,
			UDMACRCErrors:        d.ATA.UDMACRCErrors,
		}
	}
	return body
}
