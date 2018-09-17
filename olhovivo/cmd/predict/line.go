package predict

import (
	"fmt"
	"os"
	"strconv"

	olhovivo "github.com/nikolvs/olhovivo-go"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func LineCommand() *cobra.Command {
	// flags variables
	var stop string

	// command setup
	cmd := &cobra.Command{
		Use:   "line",
		Short: "",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			client := olhovivo.New(viper.GetString("token"))
			lines, err := client.QueryLines(args[0])

			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"Linha", "Destino", "Parada", "Previsão"})

			for _, line := range lines {
				previsions, err := client.LinePrevisions(line.Cl)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				lineName := line.Tp
				if line.Sl == 2 {
					lineName = line.Ts
				}

				if len(previsions.Ps) != 0 {
					for _, stop := range previsions.Ps {
						if len(stop.Vs) != 0 {
							for _, localization := range stop.Vs {
								// hadouken
								table.Append([]string{
									fmt.Sprintf("%s-%s", line.Lt, strconv.Itoa(line.Tl)),
									lineName,
									stop.Np,
									localization.T,
								})
							}
						}
					}
				}

			}

			table.Render()
		},
	}

	// set flags
	cmd.Flags().StringVar(&stop, "stop", "", "Display results only for a specific bus stop.")

	return cmd
}
