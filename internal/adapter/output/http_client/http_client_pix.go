package http_client

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/converter"
	"github.com/ThalesMonteir0/care-central/internal/adapter/output/model/response"
	"github.com/ThalesMonteir0/care-central/internal/application/domain"
	"github.com/ThalesMonteir0/care-central/internal/application/port/output"
	"github.com/ThalesMonteir0/care-central/pkg/rest_err"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	EFI_CLIENT_ID_KEY     = "EFI_CLIENT_ID_KEY"
	EFI_CLIENT_SECRET_KEY = "EFI_CLIENT_SECRET_KEY"
	EFI_URL               = "EFI_URL"
	IS_PRD                = "IS_PRD"
)

type httpClientPix struct {
}

func NewHttpClientPix() output.HttpClientPix {
	return &httpClientPix{}
}

func (h *httpClientPix) CreatePixCob(domain domain.CreatePixDomain, token string) (response.CreatePixEFIResponse, *rest_err.RestErr) {
	url := getUrlEFI()
	method := "POST"
	CreatePixSendObject := converter.ConvertDomainCreatePixToSendPixEFI(domain)
	bodyBytes, err := json.Marshal(CreatePixSendObject)
	if err != nil {
		return response.CreatePixEFIResponse{}, rest_err.NewInternalServerError("unable create body to create pix")
	}
	body := strings.NewReader(string(bodyBytes))
	var pixResponse response.CreatePixEFIResponse

	_, cert := loadCertificate()

	client := &http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/v2/cob", url), body)
	if err != nil {
		return response.CreatePixEFIResponse{}, rest_err.NewInternalServerError("unable create request to send pix create")
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	res, err := client.Do(req)
	if err != nil {
		return response.CreatePixEFIResponse{}, rest_err.NewInternalServerError("unable request create pix")
	}

	defer res.Body.Close()

	bodyResponse, err := io.ReadAll(res.Body)
	if err != nil {
		return response.CreatePixEFIResponse{}, rest_err.NewInternalServerError("unable read body")
	}

	if err := json.Unmarshal(bodyResponse, &pixResponse); err != nil {
		return response.CreatePixEFIResponse{}, rest_err.NewInternalServerError("unable parse body to response object")
	}

	return pixResponse, nil
}

func (h *httpClientPix) AuthPix() (string, *rest_err.RestErr) {
	url := getUrlEFI()
	clientID := getClientIDKey()
	secretClient := getClientSecretKey()
	method := "POST"
	bodyString := strings.NewReader(`{\"grant_type\": \"client_credentials\"}`)
	var pixOauth response.PixOauthResponse

	_, cert := loadCertificate()

	client := http.Client{
		Timeout: 60 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	req, err := http.NewRequest(method, fmt.Sprintf("%s/oauth/token", url), bodyString)
	if err != nil {
		return "", rest_err.NewInternalServerError("unable create request to EFI oauth")
	}

	req.SetBasicAuth(clientID, secretClient)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return "", rest_err.NewInternalServerError("Unable authentication to EFI")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", rest_err.NewInternalServerError("Unable read body")
	}

	if err := json.Unmarshal(body, &pixOauth); err != nil {
		return "", rest_err.NewInternalServerError("Unable unmarshl body to struct")
	}

	return pixOauth.AccessToken, nil
}

func getClientIDKey() string {
	return os.Getenv(EFI_CLIENT_ID_KEY)
}

func getClientSecretKey() string {
	return os.Getenv(EFI_CLIENT_SECRET_KEY)
}

func getUrlEFI() string {
	return os.Getenv(EFI_URL)
}

func loadCertificate() (*rest_err.RestErr, tls.Certificate) {
	var cert tls.Certificate
	isPrdBool := os.Getenv(IS_PRD)
	isPrd, err := strconv.ParseBool(isPrdBool)
	if err != nil {
		return rest_err.NewInternalServerError("error parse string to bool"), tls.Certificate{}
	}

	if !isPrd {
		cert, _ = tls.LoadX509KeyPair("internal/certificados/homolog/care-central-certificado-homolog.crt.pem", "internal/certificados/homolog/private-key-care-central.key.pem")
	} else {
		cert, _ = tls.LoadX509KeyPair("", "")
	}

	return nil, cert
}
