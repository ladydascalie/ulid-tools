<script lang="ts">
	import Copier from "./Copier.svelte";
	interface Data {
		ulid: string;
		ulidAsHex: string;
		ulidAsHexPrefixed: string;
		uuid: string;
		ulidTimeComponent: string;
	}
	export let data: Data[] = [];
	export let sqlStyle = false;
	export let hexStyle = true;
	export let successMessage: string;
	$: formattedData = data.map((entry) => {
		console.log(sqlStyle);
		return {
			...entry,
			formattedTimeComponent: sqlStyle
				? formatTimeComponentAsSQL(entry.ulidTimeComponent)
				: entry.ulidTimeComponent,
			ulidAsHex: hexStyle ? entry.ulidAsHexPrefixed : entry.ulidAsHex,
		};
	});

	function formatTimeComponentAsSQL(isoString) {
		const date = new Date(isoString);
		return `'${date.toISOString().slice(0, 23).replace("T", " ")}'`;
	}
</script>

<ulid>
	<table>
		{#if data.length === 0}
			<tr>
				<td>No ULID data available.</td>
			</tr>
		{:else}
			<tr>
				<th>ULID</th>
				<th>ULID as Hex</th>
				<th>UUID</th>
				<th>ULID Time Component</th>
			</tr>
			<!-- {#each data as { ulid, ulidAsHex, ulidAsHexPrefixed, ulidTimeComponent }} -->
			{#each formattedData as item}
				<tr>
					<td>
						<Copier message={item.ulid} bind:successMessage />
					</td>
					<td>
						<Copier message={item.ulidAsHex} bind:successMessage />
					</td>
					<td>
						<Copier message={item.uuid} bind:successMessage />
					</td>
					<td>
						<Copier message={item.formattedTimeComponent} bind:successMessage />
					</td>
				</tr>
			{/each}
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
</style>
