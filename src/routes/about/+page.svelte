<script lang="ts">
	import {
		RevealedStore,
		RevealedPhrases,
		RevealedStoreKeysE,
		handleReveal
	} from '../../stores/RevealedStore';
	import Dots from '../../components/Dots.svelte';
</script>

<div class="max-w-xl mx-auto">
<h1>About page</h1>

<h3>I think that I have provided enough information on the <a href="/">home page</a>.</h3>
<h3 on:mouseenter={() => handleReveal(RevealedStoreKeysE.EMAIL)}>
	But if you are looking for my humanreadable email:
	{#if $RevealedStore.EMAIL}
		<address><a href="mailto:{RevealedPhrases.EMAIL}">{RevealedPhrases.EMAIL}</a></address>
	{:else}
		<Dots />
	{/if}
</h3>
<h3 on:mouseenter={() => handleReveal(RevealedStoreKeysE.GITHUB)}>
	Some other most recent project that I am working on could be found on
	{#if $RevealedStore.GITHUB}
		<a href={RevealedPhrases.GITHUB} target="_blank">github</a>
	{:else}
		<Dots/>
	{/if}
</h3>
<h3 on:mouseenter={() => handleReveal(RevealedStoreKeysE.LINKEDIN)}>
	But I also have
	{#if $RevealedStore.LINKEDIN}
		<a href={RevealedPhrases.LINKEDIN} target="_blank">linkedin</a>
	{:else}
		linkedin
	{/if}
	(please hover over one more time)
</h3>
</div>

<style>
	/* code {
		display: block;
		border-radius: var(--border-radius);
		font-family: Consolas, 'courier new';
		background-color: var(--bg-3);
		padding: 2px;
		font-size: 105%;
	} */
	address {
		margin: 1%;
	}
</style>
