package core

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/oslokommune/okctl/pkg/api"
)

func decodeCreateIdentityPool(_ context.Context, r *http.Request) (interface{}, error) {
	var opts api.CreateIdentityPoolOpts

	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func decodeCreateIdentityPoolClient(_ context.Context, r *http.Request) (interface{}, error) {
	var opts api.CreateIdentityPoolClientOpts

	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func decodeCreateIdentityPoolUser(_ context.Context, r *http.Request) (interface{}, error) {
	var opts api.CreateIdentityPoolUserOpts

	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func decodeDeleteIdentityPoolUser(_ context.Context, r *http.Request) (interface{}, error) {
	var opts api.DeleteIdentityPoolUserOpts

	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func decodeDeleteIdentityPool(_ context.Context, r *http.Request) (interface{}, error) {
	var opts api.DeleteIdentityPoolOpts

	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}

func decodeDeleteIdentityPoolClient(_ context.Context, r *http.Request) (interface{}, error) {
	var opts api.DeleteIdentityPoolClientOpts

	err := json.NewDecoder(r.Body).Decode(&opts)
	if err != nil {
		return nil, err
	}

	return opts, nil
}
