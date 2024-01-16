<script lang="ts">
	import { fly } from 'svelte/transition';
	import { elasticOut } from 'svelte/easing';
	import { hoveredHint, hintVisibility, hintShown } from '../stores/HintFeedback';
	import { RevealedStoreKeysE, handleReveal } from '../stores/RevealedStore';

	const PopinDuration = 5000;
	let radius = 60;
	let show = false;
	let frozen = false;

	const t = setTimeout(() => {
		if (!$hoveredHint) {
			show = true;
		}
		clearTimeout(t);
	});

	const handleHide = () => {
		handleReveal(RevealedStoreKeysE.HINTBADGE);
		$hoveredHint = true;
		if (!frozen) {
			show = false;
			frozen = true;
			const t = setTimeout(() => {
				frozen = false;
				clearTimeout(t);
			}, PopinDuration);
		}
	};
	const handleShow = () => {
		hintShown();
		show = true;
	};

	const spin = (node: Node, { duration }: { duration: number }) => {
		return {
			duration,
			css: (t: number) => {
				const eased = elasticOut(t);

				return `
          transform: scale(${eased * 0.2 + 0.8});
					fill: hsl(
						${Math.trunc(t * 360)},
						${Math.min(50, 1000 * (1 - t))}%,
						${Math.min(50, 500 * (1 - t))}%
					);`;
			}
		};
	};
</script>

{#if show}
	<div in:fly={{ duration: 1000, x: -300, opacity: 0.1 }} out:fly={{ x: -300 }}>
		<p on:mouseenter={handleHide}><i>Sometimes it is helpful to hover over "..."</i></p>
		<svg width="180" height="180">
			<circle class="fill-bg-3 dark:d-bg-3 opacity-50" in:spin={{ duration: PopinDuration }} cx="60" cy="70" r={radius} />
			<circle class="fill-bg-1 dark:d-bg-1 opacity-60" cx="60" cy="70" r="50" />
		</svg>
	</div>
{:else}
	<div class="text-center rotate-90" style="left:{$hintVisibility}px;">
		<button class='text-fg-2 dark:text-d-fg-2 bg-bg-3 dark:bg-d-bg-3'
			on:mouseenter={handleShow}
		>hint</button>
	</div>
{/if}

<style>
	svg {
		position: absolute;
		top: -35px;
		left: -10px;
		z-index: -1;
	}
	div {
		z-index: 1;
		font-size: smaller;
		position: absolute;
		width: 100px;
		top: 220px;
	}
	p {
		margin: 0px;
	}
</style>
