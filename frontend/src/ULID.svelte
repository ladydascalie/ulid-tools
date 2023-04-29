<script lang="ts">
  export let ulid: string;
  export let ulidAsHex: string;
  export let ulidAsHexPrefixed: string;
  export let ulidTimeComponent: string;

  export let successMessage: string;

  function copy(e: MouseEvent) {
    e.preventDefault();

    let target = e.target as HTMLElement;
    let text = target.innerText;

    navigator.clipboard.writeText(text);

    successMessage = "Copied to clipboard.";
    setTimeout(() => {
      successMessage = "";
    }, 5000);
  }
</script>

<ulid>
  <table>
    {#if ulid}
      <tr>
        <td>ULID</td>
        <td><button on:click={copy}>{ulid}</button></td>
      </tr>
    {/if}
    {#if ulidAsHex}
      <tr>
        <td>ULID as Hex</td>
        <td><button on:click={copy}>{ulidAsHex}</button></td>
      </tr>
    {/if}
    {#if ulidAsHexPrefixed}
      <tr>
        <td>ULID as Hex (prefixed)</td>
        <td><button on:click={copy}>{ulidAsHexPrefixed}</button></td>
      </tr>
    {/if}
    {#if ulidTimeComponent}
      <tr>
        <td>ULID time component</td>
        <td><button on:click={copy}>{ulidTimeComponent}</button></td>
      </tr>
    {/if}
  </table>
</ulid>

<style>
  table {
    text-align: left;
    margin: 0 auto;
  }
  td {
    padding: 0.5rem;
  }

  button {
    background-color: #2d2d2d;
    color: #f1f1f1;
    border: none;
    border-radius: 6px;
    box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.4);
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.5em 1em;
    font-family: "Courier New", Courier, monospace;
    font-size: 1em;
    font-weight: bold;
    cursor: pointer;
    outline: none;
    transition: transform 0.1s ease-in-out, box-shadow 0.1s ease-in-out;
  }

  button:hover {
    transform: translateY(-2px);
  }

  button:active {
    transform: translateY(2x);
    box-shadow: none;
  }
</style>
