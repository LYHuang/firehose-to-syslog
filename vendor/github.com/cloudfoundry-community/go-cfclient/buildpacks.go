package cfclient

import (
	"encoding/json"
	"io/ioutil"

	"fmt"
	"github.com/pkg/errors"
	"net/http"
)

type BuildpackResponse struct {
	Count     int                 `json:"total_results"`
	Pages     int                 `json:"total_pages"`
	NextUrl   string              `json:"next_url"`
	Resources []BuildpackResource `json:"resources"`
}

type BuildpackResource struct {
	Meta   Meta      `json:"metadata"`
	Entity Buildpack `json:"entity"`
}

type Buildpack struct {
	Guid      string `json:"guid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Name      string `json:"name"`
	Enabled   bool   `json:"enabled"`
	Locked    bool   `json:"locked"`
	Filename  string `json:"filename"`
	c         *Client
}

func (c *Client) ListBuildpacks() ([]Buildpack, error) {
	var buildpacks []Buildpack
	requestUrl := "/v2/buildpacks"
	for {
		buildpackResp, err := c.getBuildpackResponse(requestUrl)
		if err != nil {
			return []Buildpack{}, err
		}
		for _, buildpack := range buildpackResp.Resources {
			buildpacks = append(buildpacks, c.mergeBuildpackResource(buildpack))
		}
		requestUrl = buildpackResp.NextUrl
		if requestUrl == "" {
			break
		}
	}
	return buildpacks, nil
}

func (c *Client) getBuildpackResponse(requestUrl string) (BuildpackResponse, error) {
	var buildpackResp BuildpackResponse
	r := c.NewRequest("GET", requestUrl)
	resp, err := c.DoRequest(r)
	if err != nil {
		return BuildpackResponse{}, errors.Wrap(err, "Error requesting buildpacks")
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return BuildpackResponse{}, errors.Wrap(err, "Error reading buildpack request")
	}
	err = json.Unmarshal(resBody, &buildpackResp)
	if err != nil {
		return BuildpackResponse{}, errors.Wrap(err, "Error unmarshalling buildpack")
	}
	return buildpackResp, nil
}

func (c *Client) mergeBuildpackResource(buildpack BuildpackResource) Buildpack {
	buildpack.Entity.Guid = buildpack.Meta.Guid
	buildpack.Entity.CreatedAt = buildpack.Meta.CreatedAt
	buildpack.Entity.UpdatedAt = buildpack.Meta.UpdatedAt
	buildpack.Entity.c = c
	return buildpack.Entity
}

func (c *Client) GetBuildpackByGuid(buidpackGUID string) (Buildpack, error) {
	requestUrl := fmt.Sprintf("/v2/buildpacks/%s", buidpackGUID)
	r := c.NewRequest("GET", requestUrl)
	resp, err := c.DoRequest(r)
	if err != nil {
		return Buildpack{}, errors.Wrap(err, "Error requestion buildpack info")
	}
	return c.handleBuildpackResp(resp)
}

func (c *Client) handleBuildpackResp(resp *http.Response) (Buildpack, error) {
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return Buildpack{}, err
	}
	var buildpackResource BuildpackResource
	if err := json.Unmarshal(body, &buildpackResource); err != nil {
		return Buildpack{}, err
	}
	return c.mergeBuildpackResource(buildpackResource), nil
}
