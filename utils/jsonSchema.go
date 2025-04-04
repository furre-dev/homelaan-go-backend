package utils

var JsonSchema = map[string]interface{}{
	"$schema": "https://json-schema.org/draft/2020-12/schema",
	"$id":     "https://github.com/furre-dev/homelaan-go-backend/graph/model/investor-profile",
	"properties": map[string]interface{}{
		"geo_info": map[string]interface{}{
			"properties": map[string]interface{}{
				"location": map[string]interface{}{
					"properties": map[string]interface{}{
						"city":    map[string]interface{}{"type": []interface{}{"string", "null"}},
						"country": map[string]interface{}{"type": []interface{}{"string", "null"}},
					},
					"additionalProperties": false,
					"type":                 "object",
				},
				"years_abroad": map[string]interface{}{"type": []interface{}{"integer", "null"}},
				"will_stay_long_term_in_current_location": map[string]interface{}{"type": []interface{}{"boolean", "null"}},
				"has_family_abroad":                       map[string]interface{}{"type": []interface{}{"boolean", "null"}},
				"residency_status":                        map[string]interface{}{"type": []interface{}{"string", "null"}},
			},
			"additionalProperties": false,
			"type":                 "object",
		},
		"professional_background": map[string]interface{}{
			"properties": map[string]interface{}{
				"current_role":          map[string]interface{}{"type": []interface{}{"string", "null"}},
				"organization":          map[string]interface{}{"type": []interface{}{"string", "null"}},
				"industry":              map[string]interface{}{"type": []interface{}{"string", "null"}},
				"career_overview":       map[string]interface{}{"type": []interface{}{"string", "null"}},
				"expertise_areas":       map[string]interface{}{"items": map[string]interface{}{"type": []interface{}{"string", "null"}}, "type": "array"},
				"seniority_level":       map[string]interface{}{"type": []interface{}{"string", "null"}},
				"decision_making_scope": map[string]interface{}{"type": []interface{}{"string", "null"}},
			},
			"additionalProperties": false,
			"type":                 "object",
		},
		"network_market_access": map[string]interface{}{
			"properties": map[string]interface{}{
				"industries_with_strong_network": map[string]interface{}{"items": map[string]interface{}{"type": []interface{}{"string", "null"}}, "type": "array"},
				"professional_contacts":          map[string]interface{}{"type": []interface{}{"string", "null"}},
				"local_business_diaspora_groups": map[string]interface{}{"items": map[string]interface{}{"type": []interface{}{"string", "null"}}, "type": "array"},
			},
			"additionalProperties": false,
			"type":                 "object",
		},
		// Include the rest of the properties here similarly...
	},
	"additionalProperties": false,
	"type":                 "object",
}
