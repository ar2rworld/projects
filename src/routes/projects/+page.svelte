<script>
  import ProjectCard from './ProjectCard.svelte';
  import { projects } from '../../stores/data';
</script>

<h1>Here should be some projects:</h1>

<div class='projectsContainer'>
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
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 1%;
    align-content: center;    
  }
</style>