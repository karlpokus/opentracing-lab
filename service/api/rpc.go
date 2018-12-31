package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func findOnePet(petName string) ([]byte, error) {
	url := "http://localhost:9113/api/pet"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Set("name", petName)
	req.URL.RawQuery = q.Encode()

	res, err := do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func findAllPets() ([]byte, error) {
	url := "http://localhost:9113/api/pets"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func getPetTypes() ([]byte, error) {
	url := "http://localhost:9113/api/pets/types"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func checkUser(authHeader string) error {
	url := "http://localhost:9112"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", authHeader)
	_, err = do(req)
	if err != nil {
		return err
	}
	return nil
}

func do(req *http.Request) ([]byte, error) {
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("%d %s", res.StatusCode, body)
	}

	return body, nil
}
