<script lang="ts">
	import { fade } from 'svelte/transition';

	import Dots from '../components/Dots.svelte';
	import {
		RevealedPhrases,
		RevealedStore,
		RevealedStoreKeysE,
		handleReveal
	} from '../stores/RevealedStore';

	let showHint = true;
	const hideHint = () => {
		showHint = false;
	}	
</script>

<div class="max-w-xl mx-auto">
	<h1>Hello!</h1>
	{ #if showHint }
		<h2 class="bg-fg-2 rounded-sm" out:fade on:mouseenter={() => hideHint()}>Hey, this is interactive page, so try hovering over <Dots/> and see what happens</h2>
	{ /if }
	<h2>And welcome to the Artur Linnik's projects website</h2>
	<h2>My name is Artur Linnik</h2>
	<h2>But you can call me Artur</h2>
	<br />
	<h3 on:mouseenter={() => handleReveal(RevealedStoreKeysE.LIFE)}>
		I hope that you will have a good
		{ #if $RevealedStore.LIFE }
			{ RevealedPhrases.LIFE }
		{ :else }
			<Dots/>
		{ /if } 
	</h3>
	<h3 on:mouseenter={() => handleReveal(RevealedStoreKeysE.HOMEINFO)}>
		If you click on the "projects" tab you will find
		{ #if $RevealedStore.HOMEINFO }
			{ RevealedPhrases.HOMEINFO }
		{ :else }
			<Dots/>
		{ /if }
	</h3>
</div>

