package proposalbcommunity

import (
	"fmt"

	"github.com/mfranczy/proposal-b-core/pkg/plugins"
)

func init() {
	fmt.Println(plugins.Plugin)
	fmt.Println("I am alive")
}
