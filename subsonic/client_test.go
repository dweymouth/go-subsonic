package subsonic

import "testing"

func runClientTests(client Client, t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		if _, err := client.Ping(); err != nil {
			t.Error("Ping failed (somehow)")
		}
	})
	t.Run("License", func(t *testing.T) {
		license, err := client.GetLicense()
		if err != nil {
			t.Error(err)
		}
		if !license.Valid {
			t.Errorf("Invalid license returned- %#v\n", license)
		}
	})
}
