package predict

import (
	"github.com/spf13/cobra"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "predict",
		Short: "",
	}

	cmd.AddCommand(LineCommand())
	return cmd
}
