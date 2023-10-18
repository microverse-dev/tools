package alchemy

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	rpcUrl string
}

func NewClient(rpcUrl string) *Client {
	return &Client{rpcUrl: rpcUrl}
}

type GetNFTsRequest struct {
	Owner           string
	ContractAddress string
	// pageKey         string
}

type GetNFTsResponse struct {
	OwnedNfts []struct {
		Id struct {
			TokenId string `json:"tokenId"`
		} `json:"id"`
		Title string `json:"title"`
	} `json:"ownedNfts"`
}

func (c *Client) GetNFTs(ctx context.Context, payload GetNFTsRequest) (*GetNFTsResponse, error) {
	url, err := url.Parse(c.rpcUrl)
	if err != nil {
		return nil, err
	}

	url.Path = path.Join(url.Path, "getNFTs")

	q := url.Query()
	q.Add("owner", payload.Owner)
	q.Add("contractAddresses[]", payload.ContractAddress)
	q.Add("withMetadata", "true")

	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("access", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var r GetNFTsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		return nil, err
	}

	return &r, nil
}
