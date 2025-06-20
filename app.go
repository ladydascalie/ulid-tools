package main

import (
	"bufio"
	"context"
	"encoding/hex"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GenerateNewULID() ParsedULID {
	// don't care about the error here, we know it's a valid ULID.
	parsed, _ := parse(ulid.Make().String())
	spew.Dump(parsed)
	return parsed
}

type ParsedULID struct {
	ULID              string
	ULIDHex           string
	ULIDHexPrefixed   string
	UUID              string
	ULIDTimeComponent time.Time
}

type ParseInputResponse struct {
	ULIDs  []ParsedULID
	Errors []string
}

func (a *App) ParseInput(input string) ParseInputResponse {
	// There could be a ULID of any format in the input, one per line.
	// We need to parse each line separately, and return the results as a list.
	var ulids []ParsedULID
	var errors []string

	line := 0

	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line++
		parsed, err := parse(scanner.Text())
		if err != nil {
			errors = append(errors, fmt.Sprintf("line %d: %s", line, err.Error()))
			continue
		}
		ulids = append(ulids, parsed)
	}

	if err := scanner.Err(); err != nil {
		return ParseInputResponse{}
	}

	ulids = slices.DeleteFunc(ulids, func(u ParsedULID) bool {
		return u == ParsedULID{}
	})

	return ParseInputResponse{
		ULIDs:  ulids,
		Errors: errors,
	}
}

func parse(input string) (ParsedULID, error) {
	// ignore empty lines.
	if input == "" {
		return ParsedULID{}, nil
	}

	// if it's quoted, remove the quotes
	input = strings.Trim(input, "\"'")

	// first try to parse it as a ULID
	id, err := ulid.ParseStrict(input)

	switch err {
	case nil:
		// it's a ulid, we're done here.
		h := hex.EncodeToString(id.Bytes())

		uid, err := uuid.FromBytes(id.Bytes())
		if err != nil {
			return ParsedULID{}, fmt.Errorf("failed converting ULID to UUID: %w", err)
		}

		return ParsedULID{
			ULID:              id.String(),
			ULIDHex:           strings.ToUpper(h),
			ULIDHexPrefixed:   "0x" + strings.ToUpper(h),
			UUID:              uid.String(),
			ULIDTimeComponent: ulid.Time(id.Time()),
		}, err
	default:
		// ok, it's not a ulid then, try to parse it as a hex string
		// always remove the "0x" prefix if present
		input = strings.TrimPrefix(input, "0x")

		// now, in case it's in UUID format, remove the dashes
		input = strings.ReplaceAll(input, "-", "")

		_, err := hex.DecodeString(input)
		if err != nil {
			// it's not a hex string either, return the error
			return ParsedULID{}, fmt.Errorf("failed decoding string, input is neither a ULID nor a hex string: %w", err)
		}

		var id ulid.ULID

		if _, err = hex.Decode(id[:], []byte(input)); err != nil {
			return ParsedULID{}, fmt.Errorf("failed decoding string, input is neither a ULID nor a hex string: %w", err)
		}

		// we're doing it again, but this time we know it's a ULID
		h := hex.EncodeToString(id.Bytes())

		uid, err := uuid.FromBytes(id.Bytes())
		if err != nil {
			return ParsedULID{}, fmt.Errorf("failed converting ULID to UUID: %w", err)
		}

		return ParsedULID{
			ULID:              id.String(),
			ULIDHex:           strings.ToUpper(h),
			ULIDHexPrefixed:   "0x" + strings.ToUpper(h),
			UUID:              uid.String(),
			ULIDTimeComponent: ulid.Time(id.Time()),
		}, nil
	}
}
