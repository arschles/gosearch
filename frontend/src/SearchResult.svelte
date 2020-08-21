<script>
export let url;
export let name;
export let idx;
console.log(`url = ${url}, name = ${name}, idx = ${idx}`)
</script>

<div class="result" id="result{idx}">
    <a href="{url}">{name}</a>
    <p><small>{url}</small></p>
</div>