<script lang='ts'>
  import { hoveredProject } from "../../stores/HoveredProject"

  export let id: string;
  export let name: string;
  export let url: string | undefined;
  export let src_url: string;
  export let technologies: {
    name: string,
    alt: string,
    img: string
  }[] | undefined;

  let more = 'more';
  let interval : NodeJS.Timeout;

  function addDots() {
    let c = 0;
    interval = setInterval(()=> {
      if ( c < 3 ) {
        more += '.'
        c += 1
      }
    }, 200)
    hoveredProject.update(() => id)
  }
  function removeDots() {
    more = 'more'
    clearInterval(interval)
    hoveredProject.update(() => "")
  }

  let hoveredProjectID: string;
  hoveredProject.subscribe( id => hoveredProjectID = id );
  // import src from "%sveltekit.assets%/telegram.png"; // "%sveltekit.assets%/{technology}";
</script>

<div class='project'>
  <h3>{name}</h3>
  <div class='grid'>
    <div>
      <ul>
        <li>
          {#if url}
            <a href={url}>URL</a>
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
    <div class="column">
      {#if technologies != undefined}
        {#each technologies as technology}
          <div class="technologyImg">
            <img src={technology.img} alt={technology.alt}/>
          </div>
        {/each}
      {:else}
        <div>
          <p>Some technologies</p>
        </div>
      {/if}
    </div>
  </div>
  <a on:mouseenter={addDots} on:mouseleave={removeDots} class="projectView" href="/projects/{id}">
    {hoveredProjectID === id ? more : hoveredProjectID === "" ? more : more+"?"}
  </a>
</div>

<style>
  .project {
    border-radius: var(--border-radius);
    border: 1px solid var(--fg-2);
    height: max-content;
  }
  p {
    margin: 0px;
  }
  div {
    padding: 1%;
  }
  .projectView {
    display: block;
    text-align: center;
    margin: auto;
    background-color: var(--bg-2);
    border: 0;
	  border-radius: 0.317rem;
    padding: 1%;
    color: var(--fg-1);
  }
  .grid {
    display: grid;
    /* gap: 2%; */
    margin: 1% 0px;
  }
  img {
    max-height: 50px;
  }
  .technologyImg {
    text-align: center;
    margin: auto;
  }
  .column {
    width: 100%;
    display: grid;
    grid-auto-flow: column;
    margin: 1%;
  }
</style>
