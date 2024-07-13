<script lang="ts">
	export let data;
	const { project } = data;
</script>

<div class="root">
	<h1>{project.name}</h1>

	<div class="grid grid-cols-1 sm:grid-cols-2 gap-4">
		<div>
			{#if project.url}
				<a href={project.url}>URL</a>
			{:else}
				<p>Is not available publically</p>
			{/if}

			<p>You can find sources <a href={project.src_url}>here</a></p>

			{#if project.descriptionHTML != undefined}
				{@html project.descriptionHTML}
			{:else}
				<p>{project.description}</p>
			{/if}
		</div>
		<div>
			<img class="mx-auto min-h-2 rounded-sm border-2 border-fg-2" src="/{project.logo}" alt="{project.name} project logo" />
			{#if project.technologies}
				<h4>Built with:</h4>
				<div class="technologiesGrid">
					{#each project.technologies as p}
						<div class="technologyItem">
							<img src={'/' + p.img} alt={p.alt} />
							<p>{p.name}</p>
						</div>
					{/each}
				</div>
			{/if}
		</div>
	</div>

</div>

<style>
	div.root {
		margin: 0px 3%;
		padding: 1px 2%;
		/* background-color: var(--bg-3); */
		border-radius: var(--border-radius);
		border: 1px solid var(--fg-3);
	}
	img {
		width: 30%;
		/* widows: 100px;  */
	}
	div.technologiesGrid {
		display: grid;
		grid-auto-flow: column;
		align-items: center;
		justify-content: space-evenly;
	}
	div.technologyItem {
		text-align: center;
	}
</style>
