package services

import (
	"encoding/json"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/lib/utils"
	"github.com/gravitational/trace"
)

const PAMConfigSchema = `{
	"type": "object",
	"additionalProperties": false,
	"properties": {
		"enabled": {
			"type": "boolean"
		},
		"service_name": {
			"type": "string"
		},
		"use_pam_auth": {
			"type": "boolean"
		},
		"environment": {
			"type": "object"
		},
	}
}`

// UnmarshalPAMConfig unmarshals JSON into a PAMConfig resource.
func UnmarshalPAMConfig(bytes []byte, opts ...MarshalOption) (types.PAMConfig, error) {
	var config types.PAMConfigV3

	if len(bytes) == 0 {
		return nil, trace.BadParameter("missing resource data")
	}

	cfg, err := CollectOptions(opts)
	if err != nil {
		return nil, trace.Wrap(err)
	}

	if cfg.SkipValidation {
		if err := utils.FastUnmarshal(bytes, &config); err != nil {
			return nil, trace.BadParameter(err.Error())
		}
	} else {
		err := utils.UnmarshalWithSchema(PAMConfigSchema, &config, bytes)
		if err != nil {
			return nil, trace.BadParameter(err.Error())
		}
	}

	if cfg.ID != 0 {
		config.SetResourceID(cfg.ID)
	}
	if !cfg.Expires.IsZero() {
		config.SetExpiry(cfg.Expires)
	}

	return &config, nil
}

// MarshalPAMConfig marshals the PAMConfig resource to JSON.
func MarshalPAMConfig(c types.PAMConfig, opts ...MarshalOption) ([]byte, error) {
	return json.Marshal(c)
}
