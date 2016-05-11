package datagovsg

import (
	"encoding/json"
	"net/http"
	"sync"
)

// ClientResult contains the result from the HTTP request from Client
type ClientResult struct {
	Body interface{}
	Err  error
}

// Client is a special HTTP client that batches HTTP requests for the same URL, returning requests through channels
type Client struct {
	APIKey string

	listeners    map[string][]chan ClientResult
	listenerLock sync.RWMutex
}

// NewClient returns a new Client
func NewClient(apiKey string) *Client {
	return &Client{
		APIKey:       apiKey,
		listeners:    map[string][]chan ClientResult{},
		listenerLock: sync.RWMutex{},
	}
}

// broadcastOnce Broadcasts to all listeners and close channel immediately. No new listeners can register at this time.
func (c *Client) broadcastOnce(url string, result ClientResult) {
	c.listenerLock.Lock()
	listeners, _ := c.listeners[url]
	for _, listener := range listeners {
		listener <- result
		close(listener)
	}

	c.listeners[url] = nil
	delete(c.listeners, url)

	c.listenerLock.Unlock()
}

func (c *Client) register(url string) (ch chan ClientResult, alreadyExists bool) {
	ch = make(chan ClientResult)
	c.listenerLock.Lock()
	_, alreadyExists = c.listeners[url]
	if !alreadyExists {
		c.listeners[url] = []chan ClientResult{}
	}
	c.listeners[url] = append(c.listeners[url], ch)
	c.listenerLock.Unlock()
	return ch, alreadyExists
}

func (c *Client) request(method string, url string, target interface{}) chan ClientResult {

	ch, alreadyExists := c.register(url)
	if alreadyExists {
		return ch
	}

	// set up go-routine to make batched request
	go func(url string) {

		// create request
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			c.broadcastOnce(url, ClientResult{
				Err: err,
			})
			return
		}

		// set API Key
		req.Header.Set("api-key", c.APIKey)

		// make HTTP request
		client := &http.Client{}
		res, err := client.Do(req)
		if err != nil {
			c.broadcastOnce(url, ClientResult{
				Err: err,
			})
			return
		}
		defer res.Body.Close()

		// decode as JSON response
		err = json.NewDecoder(res.Body).Decode(target)

		c.broadcastOnce(url, ClientResult{
			Body: target,
			Err:  err,
		})

	}(url)
	return ch
}

// Get allows user to make a /GET HTTP request, getting it through a channel.
func (c *Client) Get(url string, target interface{}) chan ClientResult {
	return c.request("GET", url, target)
}
