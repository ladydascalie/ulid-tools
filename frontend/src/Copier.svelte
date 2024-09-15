<script lang="ts">
	export let message: string;
	export let successMessage: string = "";
	function copy(e: MouseEvent) {
		e.preventDefault();

		let target = e.target as HTMLElement;
		let text = target.innerText;

		navigator.clipboard.writeText(text);
		successMessage = "Copied to clipboard!";
	}
</script>

<button on:click={copy}>
	{message}
</button>

<style>
	button {
		background-color: #2d2d2d;
		color: #f1f1f1;
		border: none;
		border-radius: 6px;
		box-shadow: 0px 5px 10px rgba(0, 0, 0, 0.4);
		display: flex;
		align-items: center;
		justify-content: center;
		padding: 0.3em 0.8em; /* Reduced padding */
		font-family: "Courier New", Courier, monospace;
		font-size: 1em;
		font-weight: bold;
		cursor: pointer;
		outline: none;
		transition:
			transform 0.1s ease-in-out,
			box-shadow 0.1s ease-in-out;
	}

	button:hover {
		transform: translateY(-2px);
	}

	button:active {
		transform: translateY(2x);
		box-shadow: none;
	}
</style>
