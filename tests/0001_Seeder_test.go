package tests

import (
	"testing"

	"github.com/ArchishmanSengupta/expense-tracker/cmd"
	"github.com/ArchishmanSengupta/expense-tracker/utils"
)

func TestSeeder(t *testing.T) {
	t.Run("Seed Transactions", func(t *testing.T) {
		err := utils.SeedTransactions(cmd.DbConn)
		if err != nil {
			t.Error(err)
		}
	})
}
