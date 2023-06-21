package podmanextension

import (
	"fmt"

	"github.com/adrianriobo/goax/pkg/util/logging"
)

func cleanup() error {
	logging.InitLogrus("", "", "")
	return fmt.Errorf("not implemented yet")
}

func installer(userPassword string) error {
	return fmt.Errorf("not implemented yet")
}
