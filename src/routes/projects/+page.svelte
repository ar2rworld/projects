<script>
	import ProjectCard from '../../components/ProjectCard.svelte';
	import { projects } from '../../stores/data';
</script>

<h1 class="p-1">Here should be some projects:</h1>

<div class="grid gap-5 projectsContainer">
	{#each projects as project}
		<ProjectCard
			id={project.id}
			name={project.name}
			url={project.url}
			src_url={project.src_url}
			technologies={project.technologies}
		/>
	{/each}
</div>

<style>
	.projectsContainer {
		grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
	}
</style>
