package main

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/go-resty/resty/v2"
)

type Validator struct {
	OperatorAddress string `json:"operator_address"`
}

type Pagination struct {
	NextKey string `json:"next_key"`
	Total   int    `json:"total,string"`
}
type GetActiveSetResponse struct {
	Validators []Validator `json:"validators"`
	Pagination Pagination  `json:"pagination"`
}

type RESTNodeClientConfig struct {
	Hosts []string
}

type RESTNodeClient struct {
	client *resty.Client
	config RESTNodeClientConfig
}

func New(cfg RESTNodeClientConfig) (c *RESTNodeClient, e error) {
	client := resty.New()

	c = &RESTNodeClient{
		client: client,
		config: cfg,
	}

	return c, nil
}

func (c *RESTNodeClient) GetActiveSet() ([]Validator, error) {
	// TODO: to get limit from config
	limit := 40
	total := 0
	received := 0
	next := ""

	result := []Validator{}

	for received == 0 || received < total || next != "" {
		params := url.Values{}
		params.Add("status", "BOND_STATUS_BONDED")
		params.Add("pagination.limit", fmt.Sprint(limit))
		if next != "" {
			params.Add("pagination.key", next)
		} else {
			params.Add("pagination.offset", fmt.Sprint(received))
		}

		//fmt.Println(c.config.Host + "/cosmos/staking/v1beta1/validators?" + params.Encode())
		resp, err := c.client.R().Get(
			c.config.Hosts[0] + "/cosmos/staking/v1beta1/validators?" + params.Encode(),
		)
		if err != nil {
			return nil, err
		}

		r := GetActiveSetResponse{}
		//fmt.Println("resp.Body(): ", string(resp.Body()))
		e := json.Unmarshal(resp.Body(), &r)
		if e != nil {
			return nil, e
		}
		//fmt.Println("r: ", len(r.Validators))
		//fmt.Println("Pagination", r.Pagination)

		received = received + len(r.Validators)
		result = append(result, r.Validators...)
		total = int(r.Pagination.Total)
		next = r.Pagination.NextKey
		//fmt.Println(received, total, next)
	}

	return result, nil
}
