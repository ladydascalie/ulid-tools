<script lang="ts">
  // import { errorIcon } from "./assets/images/error.svg";
  import { GenerateNewULID, ParseInput } from "../wailsjs/go/main/App.js";

  let errorMessage: string;
  let input: string;

  let ulid: string;
  let ulidAsHexPrefixed: string;
  let ulidAsHex: string;
  let ulidTimeComponent: string;

  function generateNewULID(): void {
    errorMessage = "";

    GenerateNewULID()
      .then((result) => {
        ulid = result;
        input = result;
      })
      .finally(() => {
        parseInput();
      })
      .catch((error) => (errorMessage = error));
  }

  function parseInput(): void {
    errorMessage = "";

    if (input === "" || input === undefined) {
      errorMessage = "Please enter a ULID or ULID as Hex to get started.";
      return;
    }

    ParseInput(input)
      .then((result) => {
        ulid = result.ULID;
        ulidAsHexPrefixed = result.ULIDHexPrefixed;
        ulidAsHex = result.ULIDHex;
        ulidTimeComponent = result.ULIDTimeComponent;
      })
      .catch((error) => (errorMessage = error));
  }
</script>

<main>
  <div class="center">
    <button class="btn" on:click={generateNewULID}>Generate new ULID</button>
    <span class="or">or...</span>
    <form id="form" on:submit|preventDefault={parseInput}>
      <input type="text" id="input" bind:value={input} />
      <button type="submit">submit</button>
    </form>
  </div>
  {#if errorMessage}
    <div class="error-box">
      <span class="error-icon">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 24 24"
          width="24"
          height="24"
        >
          <circle cx="12" cy="12" r="10.5" fill="#FF5F5F" />
          <path
            d="M 8.5 8.5 L 15.5 15.5 M 15.5 8.5 L 8.5 15.5"
            stroke="#FFFFFF"
            stroke-width="2"
          />
        </svg>
      </span>
      <span class="error-message">{errorMessage}</span>
    </div>
  {/if}
  <table>
    {#if ulid}
      <tr>
        <td>ULID</td>
        <td><code>{ulid}</code></td>
      </tr>
    {/if}
    {#if ulidAsHex}
      <tr>
        <td>ULID as Hex</td>
        <td><code>{ulidAsHex}</code></td>
      </tr>
    {/if}
    {#if ulidAsHexPrefixed}
      <tr>
        <td>ULID as Hex (prefixed)</td>
        <td><code>{ulidAsHexPrefixed}</code></td>
      </tr>
    {/if}
    {#if ulidTimeComponent}
      <tr>
        <td>ULID time component</td>
        <td><code>{ulidTimeComponent}</code></td>
      </tr>
    {/if}
  </table>
</main>

<style>
  table {
    text-align: left;
    margin: 0 auto;
  }
  td {
    padding: 0.5rem;
  }
  code {
    font-size: 1.5rem;
  }
  .center {
    padding: 15px;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100%;
  }
  .or {
    margin: 0 1rem;
  }
  .error-box {
    display: flex;
    align-items: center;
    justify-content: center;
    background-color: #ffc7c7;
    border: 1px solid #ff5f5f;
    border-radius: 5px;
    color: #b40404;
    font-size: 16px;
    padding: 10px;
    margin: 10px;
  }

  .error-icon {
    margin-right: 10px;
    font-size: 24px;
  }

  .error-message {
    font-weight: bold;
  }
</style>
