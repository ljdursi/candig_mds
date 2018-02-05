package repos

import (
	"database/sql"
	"encoding/json"
	"log"

	database "github.com/Bio-core/jtree/database"
	"github.com/Bio-core/jtree/models"
)

var starSample = "sample_id, facility, test_requested, se_num, date_collected, date_received, sample_type, material_received, material_received_num, material_received_other, volume_of_blood_marrow, surgical_num , tumor_site, historical_diagnosis, tumor_percnt_of_total, tumor_percnt_of_circled, reviewed_by, h_e_slide_location, non_uhn_id, name_of_requestor, dna_concentration, dna_volume, dna_location, rna_concentration, rna_volume, rna_location, wbc_location, plasma_location, cf_plasma_location, pb_bm_location, rna_lysate_location, sample_size, study_id, sample_name, date_submitted, container_type, container_name, container_id, container_well, copath_num, other_identifier, has_sample_files, dna_sample_barcode, dna_extraction_date, dna_quality, ffpe_qc_date, delta_ct_Value, comments, rnase_p_date, dna_quality_by_rnase_p, rna_quality, rna_extraction_date"

//GetAllSamples gets all and results a list of samples
func GetAllSamples() []*models.Sample {
	var list []*models.Sample
	rows, err := database.DB.Query("SELECT " + starSample + " FROM Samples")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		sample := getSampleRow(rows)
		list = append(list, &sample)
	}
	return list
}

//GetManySamplesByString gets all and results a list of samples
func GetManySamplesByString(field, value string) []*models.Sample {
	var query = "SELECT " + starSample + " FROM Samples Where " + field + "= '" + value + "'"
	var list []*models.Sample
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		sample := getSampleRow(rows)
		list = append(list, &sample)
	}
	return list
}

//GetOneSampleByString gets a bio sample and returns it based on the strings provided
func GetOneSampleByString(field, value string) *models.Sample {
	var sample models.Sample
	var query = "SELECT " + starSample + " FROM Samples Where " + field + "= '" + value + "' LIMIT 1"
	rows, err := database.DB.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		sample = getSampleRow(rows)
	}
	return &sample
}

//InsertSample allows users to add generic objects to a collection in the database
func InsertSample(sample *models.Sample) bool {
	j, _ := json.Marshal(&sample)
	var intersample interface{}
	json.Unmarshal(j, &intersample)
	//Implement here
	return true
}

//RemoveAllSamples will empty a collection
func RemoveAllSamples() bool {
	//Implement here
	return true
}

func getSampleRow(rows *sql.Rows) models.Sample {
	var sample models.Sample
	err := rows.Scan(
		&sample.SampleID,
		&sample.Facility,
		&sample.TestRequested,
		&sample.SeNum,
		&sample.DateCollected,
		&sample.DateReceived,
		&sample.SampleType,
		&sample.MaterialReceived,
		&sample.MaterialReceivedNum,
		&sample.MaterialReceivedOther,
		&sample.VolumeOfBloodMarrow,
		&sample.SurgicalNum,
		&sample.TumorSite,
		&sample.HistoricalDiagnosis,
		&sample.TumorPercntOfTotal,
		&sample.TumorPercntOfCircled,
		&sample.ReviewedBy,
		&sample.HESlideLocation,
		&sample.NonUhnID,
		&sample.NameOfRequestor,
		&sample.DnaConcentration,
		&sample.DnaVolume,
		&sample.DnaLocation,
		&sample.RnaConcentration,
		&sample.RnaVolume,
		&sample.RnaLocation,
		&sample.WbcLocation,
		&sample.PlasmaLocation,
		&sample.CfPlasmaLocation,
		&sample.PbBmLocation,
		&sample.RnaLysateLocation,
		&sample.SampleSize,
		&sample.StudyID,
		&sample.SampleName,
		&sample.DateSubmitted,
		&sample.ContainerType,
		&sample.ContainerName,
		&sample.ContainerID,
		&sample.ContainerWell,
		&sample.CopathNum,
		&sample.OtherIdentifier,
		&sample.HasSampleFiles,
		&sample.DnaSampleBarcode,
		&sample.DnaExtractionDate,
		&sample.DnaQuality,
		&sample.FfpeQcDate,
		&sample.DeltaCtValue,
		&sample.Comments,
		&sample.RnasePDate,
		&sample.DnaQualityByRnaseP,
		&sample.RnaQuality,
		&sample.RnaExtractionDate)
	if err != nil {
		log.Fatal(err)
	}
	return sample
}