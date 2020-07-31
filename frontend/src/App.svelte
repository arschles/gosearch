<script>
	let term;
	async function doSearch() {
		if(term == "") {
			return; // TODO: error somewhere?
		}
		let url = `/api/search?term=${term}`
		let rawResponse = await fetch(url);
		// let dbPromise = connectToDB()
		// // this code executes even before the DB is connected
		// dbPromise.then(() => {
		// 	//this code executes only after the DB is connected
		// }
		let jsonResponse = rawResponse.json();
		let results = jsonResponse["results"];
		for(result in results) {
			// TODO: what should we do now?
		}
		// 1. get the search term from the input box.
		// 		I think there's some way to bind the valuje of the
		//		input box to a JS var
		// 2. make the call to /api/search?term
		// 3. render the results - use a new svelte component
	}
</script>

<main>
	<h1>WORKING TITLE??? We're not ready to commit</h1>
	<input type="text" bind:value={term} />
	<!--
		TODO: need to make it do the search if you also press
		enter in the search box
	-->
	<button on:click={doSearch}>I'm feeling search-ey</button>
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