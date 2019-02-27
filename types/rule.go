package types

import (
	"errors"
	"github.com/Knetic/govaluate"
	"log"
)

type Rules []Rule

type Rule struct {
	Rule    string `yaml:"rule"`
	Block   bool   `yaml:"block"`
	Require bool   `yaml:"require"`
	Error   string `yaml:"error"`
}

func (rules Rules) ExecuteRules(parameters govaluate.MapParameters) error {
	for _, rule := range rules {
		expr, err := govaluate.NewEvaluableExpression(rule.Rule)
		if err != nil {
			log.Printf("Error while parsing rule '%s': %s", rule.Rule, err.Error())
			return ErrorConfiguration
		}

		var valueBoolean bool
		value, err := expr.Eval(parameters)

		//If no other errors, make sure this is a boolean.
		if err == nil {
			ok := false
			valueBoolean, ok = value.(bool)
			if !ok {
				log.Printf("Error while executing rule '%s': %s", rule.Rule, "rule does not resolve to a bool")
				return ErrorConfiguration
			}
		}
		if err != nil {
			log.Printf("Error while executing rule '%s': %s", rule.Rule, err.Error())
			return ErrorConfiguration
		}

		//If this is true, the rule has failed.
		if rule.Block && valueBoolean || rule.Require && !valueBoolean {
			if rule.Error != "" {
				return errors.New(rule.Error)
			} else {
				return ErrorNotAuthorized
			}
		}
	}
	return nil
}
