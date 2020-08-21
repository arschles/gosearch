<script>
	import { onMount } from 'svelte';
	import SearchResult from './SearchResult.svelte'
	let term;
	let results = [];
	let errorMsg;
	let inProgress = false;
	async function doSearch() {
		console.log("doSearch");
		if(!term) {
			errorMsg = "Select a term";
			return;
		}
		inProgress = true;
		let url = `/api/search?term=${term}`
		let rawResponse = await fetch(url);
		let jsonResponse = await rawResponse.json();
		results = jsonResponse["results"];
		inProgress = false;
	}
</script>

<main>
	<img
		alt="gopher"
		height="200"
		src="https://i.etsystatic.com/15149849/r/il/c0da64/2032115144/il_570xN.2032115144_gp7m.jpg"/>

	<img height="0" alt="lycos dog" src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcQGzidEMqeAt1arMT-zFUTgNSJMmcNFMKhe7Q&usqp=CAU" />

	{#if errorMsg }
	<div style="color:red">{errorMsg}</div>
	{/if}

	<form on:submit|preventDefault={doSearch}>
		<input id="term" type="text" bind:value={term} />
		<button type="submit">Binggo!</button>
	</form>
	<div id="results" class="list-of-results">
	{#if inProgress }
		Working on it
	{:else}
		{#each results as result, i}
			<SearchResult
				idx={i}
				url={result.url}
				name={result.name}></SearchResult>
		{/each}
	{/if}
	</div>
</main>

<style>
	#results {
		margin-top:1em;
	}
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

	#term{
		/* rockerBOO: vw: 100 = 100% viewport width */
		width: 50%; 
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>