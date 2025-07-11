package testflight

import (
	"errors"
	"github.com/Laky-64/TestFlightTrackBot/internal/db/models"
	"github.com/Laky-64/TestFlightTrackBot/internal/tor"
	"github.com/Laky-64/TestFlightTrackBot/internal/utils"
	"sync"
)

func (c *Client) Check(links []models.Link) (map[string]Result, error) {
	transaction, err := c.TorClient.NewTransaction(len(links))
	if err != nil {
		return nil, err
	}
	defer transaction.Close()
	results := make(map[string]Result, len(links))
	var wg sync.WaitGroup
	var mu sync.Mutex
	pool := utils.NewPool(tor.MaxRequestsPerInstance * c.TorClient.InstanceCount())
	for _, link := range links {
		wg.Add(1)
		pool.Enqueue(func() {
			defer wg.Done()
			for i := 0; i < 3; i++ {
				res, err := transaction.ExecuteRequest(
					link.URL,
					pickRandomUserAgent(c.userAgents),
				)
				mu.Lock()
				if err != nil {
					if res != nil && res.StatusCode == 404 {
						results[link.URL] = Result{
							Status: models.StatusInvalid,
						}
					} else {
						results[link.URL] = Result{
							Error: err,
						}
					}
					mu.Unlock()
					return
				}
				bodyString := res.String()
				if len(bodyString) < 300 {
					mu.Unlock()
					continue
				}
				var appName string
				var status models.LinkStatus
				rawAppName := RegexAppName.FindAllStringSubmatch(bodyString, -1)
				if len(rawAppName) == 0 {
					status = models.StatusClosed
				} else {
					rawStatus := RegexStatus.FindAllStringSubmatch(bodyString, -1)
					if rawStatus[0][1] == "This beta is full." {
						status = models.StatusFull
					} else {
						status = models.StatusAvailable
					}
					appName = rawAppName[0][1]
				}
				results[link.URL] = Result{
					AppName: appName,
					Status:  status,
				}
				mu.Unlock()
				return
			}
			results[link.URL] = Result{
				Error: errors.New("failed to check link after 3 attempts"),
			}
		})
	}
	wg.Wait()
	return results, nil
}
