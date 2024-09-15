<script lang="ts">
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
	let errorMessage: string = "";
	let successMessage: string = "";
	let input: string = "";
	let data: Data[] = [];
	let sqlStyle = false;

	function generateNewULID(): void {
		data = [];
		errorMessage = "";
		successMessage = "";

		GenerateNewULID()
			.then((result) => {
				input = input + result.ULID + "\n";
				toast("New ULID generated successfully.");
			})
			.finally(() => {
				parseInput();
			})
			.catch((error) => alert(error));
	}

	function parseInput(): void {
		data = [];
		errorMessage = "";
		successMessage = "";
		console.log(data);

		if (input === "" || input === undefined) {
			alert("Please enter a ULID or ULID as Hex to get started.");
			return;
		}
		ParseInput(input)
			.then((result) => {
				console.log(result);
				if (result.Errors && result.Errors.length > 0) {
					errorMessage = result.Errors.join("\n");
					alert(errorMessage);
					return;
				}
				result.ULIDs.forEach((entry) => {
					data = [
						...data,
						{
							ulid: entry.ULID,
							ulidAsHex: entry.ULIDHex,
							ulidAsHexPrefixed: entry.ULIDHexPrefixed,
							ulidTimeComponent: entry.ULIDTimeComponent,
						},
					];
				});
				if (!errorMessage) {
					toast("ULID(s) parsed successfully.");
				}
			})
			.catch((error) => {
				alert(error);
			});
	}

	function toast(message: string): void {
		errorMessage = "";
		successMessage = message;
		setTimeout(() => {
			successMessage = "";
		}, 5000);
	}

	function alert(message: string): void {
		successMessage = "";
		errorMessage = message;
		setTimeout(() => {
			errorMessage = "";
		}, 5000);
	}

	function clear(): void {
		errorMessage = "";
		successMessage = "";
		input = "";
		data = [];
	}
</script>

<main>
	<section class="content">
		<form id="form" on:submit|preventDefault={parseInput}>
			<textarea
				class="code-editor"
				name="input"
				id="input"
				cols="30"
				rows="10"
				bind:value={input}
			></textarea>
		</form>
		<button class="btn" on:click={generateNewULID}>Generate new ULID</button>
		<button class="btn" type="submit" form="form">Submit Input</button>
		<button class="btn-red" on:click={clear}>Clear</button>
		<label for="sql-style" class="checkbox-label">
			<input
				type="checkbox"
				name="sql-style"
				id="sql-style"
				bind:checked={sqlStyle}
			/>
			<span class="checkbox-custom"></span>
			SQL Style Timestamps
		</label>
		{#if errorMessage}
			<Error message={errorMessage} />
		{/if}
		{#if successMessage}
			<Success message={successMessage} />
		{/if}
		<hr />
		<Ulid {data} bind:successMessage bind:sqlStyle />
	</section>
	<footer>&copy; 2023 ULID Tools</footer>
</main>

<style>
	.content {
		padding: 20px;
	}

	footer {
		background-color: #333;
		color: white;
		text-align: center;
		padding: 10px 20px;
		position: fixed;
		width: 100%;
		bottom: 0;
	}

	.btn {
		background-color: #28a745;
		color: white;
		border: none;
		padding: 8px 16px;
		cursor: pointer;
		border-radius: 4px;
		margin: 5px;
		transition: background-color 0.3s ease;
	}

	.btn:hover {
		background-color: #218838;
	}

	.btn-red {
		background-color: #dc3545;
		color: white;
		border: none;
		padding: 8px 16px;
		cursor: pointer;
		border-radius: 4px;
		margin: 5px;
		transition: background-color 0.3s ease;
	}

	.btn-red:hover {
		background-color: #c82333;
	}

	.code-editor {
		font-family: "Courier New", Courier, monospace;
		font-size: 14px;
		line-height: 1.5;
		padding: 10px;
		border: 1px solid #ccc;
		border-radius: 0 4px 4px 0;
		background-color: #2e3440;
		color: #d8dee9;
		width: 100%;
		box-sizing: border-box;
		overflow: auto;
		height: 200px;
		resize: vertical;
	}

	.code-editor::selection {
		background: #4c566a;
	}
	.checkbox-label {
		display: flex;
		align-items: center;
		cursor: pointer;
		font-family: "Courier New", Courier, monospace;
		font-size: 0.9em;
		font-weight: bold;
		margin: 5px;
	}

	.checkbox-label input[type="checkbox"] {
		display: none;
	}

	.checkbox-custom {
		width: 16px;
		height: 16px;
		background-color: #28a745;
		border-radius: 4px;
		margin-right: 8px;
		transition: background-color 0.3s ease;
		position: relative;
	}

	.checkbox-label input[type="checkbox"]:checked + .checkbox-custom {
		background-color: #218838;
	}

	.checkbox-label input[type="checkbox"]:checked + .checkbox-custom::after {
		content: "";
		position: absolute;
		top: 2px;
		left: 5px;
		width: 4px;
		height: 8px;
		border: solid white;
		border-width: 0 2px 2px 0;
		transform: rotate(45deg);
	}

	.checkbox-label input[type="checkbox"]:focus + .checkbox-custom {
		outline: 2px solid #218838;
	}

	.checkbox-label input[type="checkbox"]:hover + .checkbox-custom {
		background-color: #218838;
	}
</style>
