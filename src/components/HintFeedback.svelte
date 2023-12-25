<script lang="ts">
	import { onDestroy } from 'svelte';
	import { fade } from 'svelte/transition';
	import { hintFeedback } from '../stores/HintFeedback';

	let message: string;
	let show = false;
	const BASE_SPEED = 1;
	const delay = 1000;
	const durationFactor = 2;

	function typewriter(node: Node) {
		const valid = node.childNodes.length === 1 && node.childNodes[0].nodeType === Node.TEXT_NODE;

		if (!valid) {
			throw new Error(`This transition only works on elements with a single text node child`);
		}

		const text = node.textContent;
		if (text != undefined) {
			const duration = text.length / (BASE_SPEED * 0.01);
			return {
				duration,
				tick: (t: number) => {
					const i = Math.trunc(text.length * t);
					node.textContent = text.slice(0, i);
				}
			};
		}
		return {};
	}

	export const print = (m: string) => {
		message = m;
		const t1 = setTimeout(() => {
			show = true;
			clearTimeout(t1);
		});

		const typeDuration = m.length / (BASE_SPEED * 0.01);
		const duration = typeDuration + delay;

		const t2 = setTimeout(() => {
			show = false;
			clearTimeout(t2);
		}, duration);

		// returns duration of printing the message and set another timeout
		return duration * durationFactor;
	};

	const unsubscribe = hintFeedback.message.subscribe((message: { message: string }) => {
		if (message) {
			const duration = print(message.message);

			hintFeedback.waitTimeWritable.set({ time: duration });
		}
	});
	onDestroy(unsubscribe);
</script>

{#if show}
	<div out:fade>
		<i transition:typewriter>{message}</i>
	</div>
{/if}
