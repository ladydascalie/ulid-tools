<script lang="ts">
  // import { errorIcon } from "./assets/images/error.svg";
  import { GenerateNewULID, ParseInput } from "../wailsjs/go/main/App.js";
  import Error from "./Error.svelte";
  import Success from "./Success.svelte";
  import Ulid from "./ULID.svelte";

  interface Data {
    ulid: string;
    ulidAsHex: string;
    ulidAsHexPrefixed: string;
    ulidTimeComponent: string;
  }

  let errorMessage: string;
  let successMessage: string;

  let input: string;

  let data: Data;

  function generateNewULID(): void {
    errorMessage = "";

    GenerateNewULID()
      .then((result) => {
        data = { ulid: result } as Data;
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
        data = {
          ulid: result.ULID,
          ulidAsHex: result.ULIDHex,
          ulidAsHexPrefixed: result.ULIDHexPrefixed,
          ulidTimeComponent: result.ULIDTimeComponent,
        };
      })
      .catch((error) => {
        errorMessage = error;
        setTimeout(() => {
          errorMessage = "";
        }, 5000);
      });
  }

  function copy(e: MouseEvent): void {
    console.log(e);
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
    <Error message={errorMessage} />
  {/if}
  {#if successMessage}
    <Success message={successMessage} />
  {/if}
  <Ulid {...data} bind:successMessage />
</main>

<style>
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
</style>
