<script lang="ts">
	import { hoveredProject } from '../stores/HoveredProject';

	export let id: string;
	export let name: string;
	export let url: string | undefined;
	export let src_url: string;
	export let technologies:
		| {
				name: string;
				alt: string;
				img: string;
		  }[]
		| undefined;

	let more = 'more';
	let interval: NodeJS.Timeout;

	function addDots() {
		let c = 0;
		interval = setInterval(() => {
			if (c < 3) {
				more += '.';
				c += 1;
			}
		}, 200);
		hoveredProject.update(() => id);
	}
	function removeDots() {
		more = 'more';
		clearInterval(interval);
		hoveredProject.update(() => '');
	}

	let hoveredProjectID: string;
	hoveredProject.subscribe((id) => (hoveredProjectID = id));
</script>

<div class="h-max p-2 rounded-lg border-2 border-fg-2">
	<h2>{name}</h2>
	<div class="grid mx-1 my-0">
		<div>
			<ul>
				<li>
					{#if url}
						<a class='text-link dark:text-d-link' href={url}>URL</a>
					{:else}
						<p>Didn't published yet</p>
					{/if}
				</li>
				<li>
					<a href={src_url}>sources</a>
				</li>
				<li>Techonologies:</li>
			</ul>
		</div>
		<div class="w-full grid grid-flow-col m-4">
			{#if technologies != undefined}
				{#each technologies as technology}
					<div class="text-center m-auto">
						<img class="max-h-12" src={technology.img} alt={technology.alt} />
					</div>
				{/each}
			{:else}
				<div>
					<p>Some technologies</p>
				</div>
			{/if}
		</div>
	</div>
	<a on:mouseenter={addDots} on:mouseleave={removeDots}
		class="block text-center m-auto bg-bg-2 dark:bg-d-bg-2 border-0 rounded-lg p-1"
		href="/projects/{id}">
		{hoveredProjectID === id ? more : hoveredProjectID === '' ? more : more + '?'}
	</a>
</div>

<style>
	p {
		margin: 0px;
	}
</style>
