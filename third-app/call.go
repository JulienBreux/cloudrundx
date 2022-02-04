package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"google.golang.org/api/idtoken"
)

func get(targetURL string) (string, error) {
	ctx := context.Background()

	client, err := idtoken.NewClient(ctx, targetURL)
	if err != nil {
		return "", fmt.Errorf("idtoken.NewClient: %v", err)
	}

	resp, err := client.Get(targetURL)
	if err != nil {
		return "", fmt.Errorf("client.Get: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %v", err)
	}

	sb := string(body)

	return sb, nil
}
