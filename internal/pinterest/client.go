package pinterest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

const baseURL = "https://api.pinterest.com/v5"

type Client struct {
	httpClient         *http.Client
	accessToken        string
	clientID           string
	clientSecret       string
	baseURL            string
	DefaultContentType string
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Image struct {
	Url string `json:"url"`
}

type Media struct {
	Images Images `json:"images"`
}

type Images struct {
	Res150x150 Image `json:"150x150"`
	Res400x300 Image `json:"400x300"`
	Res600x    Image `json:"600x"`
	Res1200x   Image `json:"1200x"`
}

type BoardData struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type BoardsData struct {
	Items    []BoardData `json:"boards"`
	Bookmark string      `json:"bookmark"`
}

type PinData struct {
	Id            string `json:"id"`
	Title         string `json:"title"`
	DominantColor string `json:"dominant_color"`
	Media         Media  `json:"media"`
	BoardId       string `json:"board_id"`
	AltText       string `json:"alt_text"`
}

type PinsData struct {
	Items []PinData `json:"pins"`
	Image string    `json:"image"`
}

func NewClient(accessToken, clientSecret, clientID string) (*Client, error) {
	if accessToken == "" {
		return nil, errors.New("accessToken is empty")
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		accessToken:        accessToken,
		clientID:           clientID,
		clientSecret:       clientSecret,
		baseURL:            baseURL,
		DefaultContentType: "application/json",
	}, nil
}

func (c *Client) NewRequest(endpoint string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, c.baseURL+endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("Content-Type", c.DefaultContentType)

	response, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode > 200 {
		defer response.Body.Close()

		errorRes, errStatus := handleWrongStatusCode(response)
		if errStatus != nil {
			return nil, errStatus
		}

		return nil, errors.New(fmt.Sprintf(
			"Error response: ErrorCode: %d ErrorMessage: %s",
			errorRes.Code,
			errorRes.Message,
		))
	}

	responseBytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return responseBytes, nil
}

func handleWrongStatusCode(res *http.Response) (errorResponse, error) {
	errorRes := errorResponse{}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return errorRes, errors.New("unable to read response body while handleWrongStatus code")
	}

	err = json.Unmarshal(bytes, &errorRes)
	if err != nil {
		return errorRes, errors.New("unable to unmarshal response body while handleWrongStatus code")
	}

	return errorRes, nil
}

func (c *Client) GetPinById(pinId string) (*PinsData, error) {
	url := "/pins/"

	if len(pinId) > 0 {
		url += "&image=" + pinId
	}

	responseBytes, err := c.NewRequest(url + pinId + "/")
	if err != nil {
		return nil, err
	}

	var pin PinsData
	err = json.Unmarshal(responseBytes, &pin)
	if err != nil {
		return nil, err
	}

	return &pin, nil
}

func (c *Client) GetPinsById(pinId string) ([]PinData, error) {
	pins, err := c.GetPinById(pinId)
	if err != nil {
		return nil, err
	}

	var responseBytes []PinData

	responseBytes = append(responseBytes, pins.Items...)

	for len(pins.Items) > 0 {
		pins, err = c.GetPinById(pins.Items[len(pins.Items)-1].Id)
		if err != nil {
			return nil, err
		}

		responseBytes = append(responseBytes, pins.Items...)
	}

	return responseBytes, nil
}

func (c *Client) GetBoard(bookmark string) (*BoardsData, error) {
	url := "/boards/?page_size=100"
	if len(bookmark) > 0 {
		url += "&bookmark=" + bookmark
	}

	bytes, err := c.NewRequest(url)
	if err != nil {
		return nil, err
	}

	var board = new(BoardsData)
	unmarshalErr := json.Unmarshal(bytes, &board)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}

	return board, nil
}

func (c *Client) GetBoards() ([]BoardData, error) {
	var resultBoards []BoardData

	boards, err := c.GetBoard("")
	if err != nil {
		return nil, err
	}

	resultBoards = append(resultBoards, boards.Items...)

	for len(boards.Bookmark) > 0 {
		boards, err = c.GetBoard(boards.Bookmark)
		if err != nil {
			return nil, err
		}

		resultBoards = append(resultBoards, boards.Items...)
	}

	return resultBoards, nil
}

func (c *Client) SearchPinsByaGivenSearchTerm(search string) (*PinData, error) {
	responseBytes, err := c.NewRequest("/search/partner/pins/" + search + "/")
	if err != nil {
		return nil, err
	}

	var pin PinData
	err = json.Unmarshal(responseBytes, &pin)
	if err != nil {
		return nil, err
	}

	return &pin, nil
}
