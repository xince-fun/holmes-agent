package main

import "github.com/spf13/cobra"

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "holmes",
		Short: "holmes is an observability collector based on eBPF",
	}

	return cmd
}
