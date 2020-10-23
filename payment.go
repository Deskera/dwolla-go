package dwolla

import (
	"encoding/json"
)

type payment struct {
	authHandler *auth
	baseURL     string
}

func PaymentHandler(paymentConfig *customer) *customer {
	return paymentConfig
}

func (p *payment) InitiateMassPayment(idempotencyKey string, massPaymentReq *MassPaymentRequest) (*MassPaymentRequest, error) {
	url := p.baseURL + "/mass-payments"

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	header := &Header{
		IdempotencyKey: idempotencyKey,
	}

	resp, err := makePostRequest(url, header, massPaymentReq, token)
	if err != nil {
		return nil, err
	}

	massPaymentLocation := resp.Header.Get(location)
	massPaymentReq.Location = massPaymentLocation

	return massPaymentReq, nil
}

func (p *payment) GetMassPaymentById(massPaymentLink string) (*MassPaymentResponse, error) {
	url := massPaymentLink

	token, err := p.authHandler.GetToken()
	if err != nil {
		return nil, err
	}

	resp, err := makeGetRequest(url, token)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var massPaymentResp MassPaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&massPaymentResp); err != nil {
		return nil, err
	}

	return &massPaymentResp, nil
}

func (p *payment) UpdateMassPaymentStatus(massPaymentLink string, status PaymentStatus) error {
	url := massPaymentLink

	token, err := p.authHandler.GetToken()
	if err != nil {
		return err
	}

	statusReq := &UpdateMassPayment{
		Status: status,
	}

	resp, err := makePostRequest(url, nil, statusReq, token)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil

}
