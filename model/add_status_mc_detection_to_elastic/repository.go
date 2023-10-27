package add_status_mc_detection_to_elastic

import (
	"strconv"
	"strings"

	esv7 "github.com/elastic/go-elasticsearch/v7"
)

type Repository interface {
	CreateElasticStatusMcDetection(elasticStatusMcDetection ElasticStatusMcDetection) (ElasticStatusMcDetection, error)
}

type repository struct {
	elasticsearch *esv7.Client
}

func NewRepository(elasticsearch *esv7.Client) *repository {
	return &repository{elasticsearch}
}

func (r *repository) CreateElasticStatusMcDetection(elasticStatusMcDetection ElasticStatusMcDetection) (ElasticStatusMcDetection, error) {

	// Menggunakan library "github.com/elastic/go-elasticsearch" untuk melakukan operasi penyimpanan
	// Gantilah `indexName` dengan nama index Elasticsearch yang sesuai
	indexName := "status_mc_detection_index"

	// Anda dapat membuat body dokumen yang akan disimpan di Elasticsearch
	// Misalnya, jika Anda ingin menyimpan data deteksi manusia yang diberikan sebagai JSON:
	body := []byte(`{
		"id": "` + elasticStatusMcDetection.ID + `",
		"tid_id": "` + strconv.Itoa(*elasticStatusMcDetection.TidID) + `",
		"date_time": "` + elasticStatusMcDetection.DateTime + `",
		"status_signal": "` + elasticStatusMcDetection.StatusSignal + `",
		"status_storage": "` + elasticStatusMcDetection.StatusStorage + `",
		"status_ram": "` + elasticStatusMcDetection.StatusRam + `",
		"status_cpu": "` + elasticStatusMcDetection.StatusCpu + `"
		
	}`)

	// Mengirimkan data ke Elasticsearch untuk disimpan
	_, err := r.elasticsearch.Index(indexName, strings.NewReader(string(body)))
	if err != nil {
		return elasticStatusMcDetection, err
	}

	// Jika operasi berhasil, Anda dapat mengembalikan data yang sama yang Anda terima sebagai argumen.
	return elasticStatusMcDetection, nil

}
