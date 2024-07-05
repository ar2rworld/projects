<script lang="ts">
	import Dots from '../components/Dots.svelte';
import {
		RevealedPhrases,
		RevealedStore,
		RevealedStoreKeysE,
		handleReveal
	} from '../stores/RevealedStore';
</script>

<div class="max-w-xl mx-auto">
	<h1>Hello!</h1>
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

