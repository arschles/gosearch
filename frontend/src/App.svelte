<script>
	import { onMount } from 'svelte';
	let term;
	let results = [];
	let errorMsg;
	async function doSearch() {
		console.log("doSearch");
		if(term == "") {
			errorMsg = "Select a term";
			return; // TODO: error somewhere?
		}
		let url = `/api/search?term=${term}`
		let rawResponse = await fetch(url);
		let jsonResponse = await rawResponse.json();
		results = jsonResponse["results"];
	}
</script>

<main>
	<h1>WORKING TITLE??? We're not ready to commit</h1>

	{#if errorMsg }
	<div style="color:red">{errorMsg}</div>
	{/if}

	<form on:submit|preventDefault={doSearch}>
	<input type="text" bind:value={term} />
	<!--
		TODO: need to make it do the search if you also press
		enter in the search box
	-->

	<button type="submit">I'm feeling search-ey</button>
	</form>
	<div id="results" class="list-of-results">
	{#each results as result, i}
		<div class="result" id="result{i}">
			<a href="{result.url}">{result.name}</a>
		</div>
	{/each}
	</div>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>