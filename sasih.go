/*
Copyright © Portalnesia <support@portalnesia.com>
*/
package saka

type SasihData struct {
	ID    int64
	RefID int64
	Name  string
	ref   *SasihData
}

func (s SasihData) Reference() SasihData {
	if s.ref == nil {
		temp := newSasihSlice()[s.RefID]
		s.ref = &temp
	}
	return *s.ref
}

type Sasih struct {
	Kasa         SasihData
	Karo         SasihData
	Katiga       SasihData
	Kapat        SasihData
	Kalima       SasihData
	Kanem        SasihData
	Kapitu       SasihData
	Kawolu       SasihData
	Kasanga      SasihData
	Kadasa       SasihData
	Destha       SasihData
	Sadha        SasihData
	MalaDestha   SasihData
	MalaSadha    SasihData
	NampihDestha SasihData
	NampihKatiga SasihData
	NampihKasa   SasihData
	NampihKadasa SasihData
	NampihKaro   SasihData
	NampihSadha  SasihData
}

type SasihDayInfoData struct {
	ID    int64
	RefID int64
	Name  string
	ref   *SasihDayInfoData
}

func (s SasihDayInfoData) Reference() SasihDayInfoData {
	if s.ref == nil {
		temp := newSasihDayInfoSlice()[s.RefID]
		s.ref = &temp
	}
	return *s.ref
}

type SasihDayInfo struct {
	Penanggal SasihDayInfoData
	Pangelong SasihDayInfoData
	Purnama   SasihDayInfoData
	Tilem     SasihDayInfoData
}

func newSasih() Sasih {
	return Sasih{
		Kasa:         SasihData{0, 0, "Kasa", nil},
		Karo:         SasihData{1, 1, "Karo", nil},
		Katiga:       SasihData{2, 2, "Katiga", nil},
		Kapat:        SasihData{3, 3, "Kapat", nil},
		Kalima:       SasihData{4, 4, "Kalima", nil},
		Kanem:        SasihData{5, 5, "Kanem", nil},
		Kapitu:       SasihData{6, 6, "Kapitu", nil},
		Kawolu:       SasihData{7, 7, "Kawolu", nil},
		Kasanga:      SasihData{8, 8, "Kasanga", nil},
		Kadasa:       SasihData{9, 9, "Kadasa", nil},
		Destha:       SasihData{10, 10, "Destha", nil},
		Sadha:        SasihData{11, 11, "Sadha", nil},
		MalaDestha:   SasihData{12, 10, "Mala Destha", nil},
		MalaSadha:    SasihData{13, 11, "Mala Sadha", nil},
		NampihDestha: SasihData{14, 10, "Nampih Destha", nil},
		NampihKatiga: SasihData{15, 2, "Nampih Katiga", nil},
		NampihKasa:   SasihData{16, 0, "Nampih Kasa", nil},
		NampihKadasa: SasihData{17, 9, "Nampih Kadasa", nil},
		NampihKaro:   SasihData{18, 1, "Nampih Karo", nil},
		NampihSadha:  SasihData{19, 11, "Nampih Sadha", nil},
	}
}

func newSasihDayInfo() SasihDayInfo {
	return SasihDayInfo{
		Penanggal: SasihDayInfoData{0, 0, "Penanggal", nil},
		Pangelong: SasihDayInfoData{1, 1, "Pangelong", nil},
		Purnama:   SasihDayInfoData{2, 0, "Purnama", nil},
		Tilem:     SasihDayInfoData{3, 1, "Tilem", nil},
	}
}

func newSasihSlice() []SasihData {
	data := newSasih()
	return []SasihData{
		data.Kasa,
		data.Karo,
		data.Katiga,
		data.Kapat,
		data.Kalima,
		data.Kanem,
		data.Kapitu,
		data.Kawolu,
		data.Kasanga,
		data.Kadasa,
		data.Destha,
		data.Sadha,
		data.MalaDestha,
		data.MalaSadha,
		data.NampihDestha,
		data.NampihKatiga,
		data.NampihKasa,
		data.NampihKadasa,
		data.NampihKaro,
		data.NampihSadha,
	}
}

func newSasihDayInfoSlice() []SasihDayInfoData {
	data := newSasihDayInfo()
	return []SasihDayInfoData{
		data.Penanggal,
		data.Pangelong,
		data.Purnama,
		data.Tilem,
	}
}
