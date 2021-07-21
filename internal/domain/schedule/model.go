package schedule

type Schedule struct {
	Quota            int    `json:"quota"`
	Date             string `json:"tanggal_vaksin"`
	FaskesScheduleID int    `json:"id_jadwal_faskes"`
	FaskesDataID     int    `json:"id_data_faskes"`
	Faskes           string `json:"faskes"`
	QuotaLeft        int    `json:"sisa"`
}

type Schedules []Schedule
