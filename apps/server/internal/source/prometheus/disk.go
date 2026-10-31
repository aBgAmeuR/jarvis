package prometheus

import "jarvis/server/internal/domain/supervision"

func newDisk(device string, labels map[string]string, values map[string]float64) supervision.Disk {
	disk := supervision.Disk{
		Device:       device,
		Model:        labels["model"],
		Serial:       labels["serial"],
		Protocol:     labels["protocol"],
		Rotational:   labels["rotational"] == "1",
		HealthPassed: values["smart_health_passed"] == 1,
		Standby:      values["smart_device_standby"] == 1,
	}
	if celsius, ok := values["smart_temperature_celsius"]; ok {
		disk.TemperatureCelsius = &celsius
	}
	if hours, ok := values["smart_power_on_hours"]; ok {
		n := int64(hours)
		disk.PowerOnHours = &n
	}
	if used, ok := values["smart_nvme_percentage_used"]; ok {
		disk.NVMe = &supervision.NVMeHealth{
			PercentageUsed:        used,
			AvailableSparePercent: values["smart_nvme_available_spare_percent"],
			CriticalWarning:       int64(values["smart_nvme_critical_warning"]),
			MediaErrors:           int64(values["smart_nvme_media_errors"]),
		}
	}
	if _, ok := values["smart_ata_reallocated_sectors"]; ok {
		disk.ATA = &supervision.ATAHealth{
			ReallocatedSectors:   int64(values["smart_ata_reallocated_sectors"]),
			PendingSectors:       int64(values["smart_ata_pending_sectors"]),
			UncorrectableSectors: int64(values["smart_ata_offline_uncorrectable"]),
			UDMACRCErrors:        int64(values["smart_ata_udma_crc_errors"]),
		}
	}
	return disk
}
