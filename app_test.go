package main

import "testing"

func TestApp_ParseInput(t *testing.T) {
	app := NewApp()

	const knownInput = "01FCB79DXEHPQJRD8BWC6G7PVH"
	const expectedHex = "0x017B1674B7AE8DAF2C350BE30D03DB71"
	const expectedDatabaseHex = "017B1674B7AE8DAF2C350BE30D03DB71"

	out, err := app.ParseInput(knownInput)
	if err != nil {
		t.Fatalf("ParseInput(%s) returned error: %s", knownInput, err)
	}

	if out.ULID != knownInput {
		t.Errorf("ParseInput(%s) returned ULID %s, expected %s", knownInput, out.ULID, knownInput)
	}

	if out.ULIDHexPrefixed != expectedHex {
		t.Errorf("ParseInput(%s) returned ULIDHex %s, expected %s", knownInput, out.ULIDHexPrefixed, expectedHex)
	}

	if out.ULIDHex != expectedDatabaseHex {
		t.Errorf("ParseInput(%s) returned ULIDDatabaseHex %s, expected %s", knownInput, out.ULIDHex, expectedDatabaseHex)
	}
}
