<script lang='ts'>
  export let id: string;
  export let name: string;
  export let url: string | undefined;
  export let scr_url: string;
  export let technologies: {
    name: string,
    alt: string,
    img: string
  }[] | undefined;

  let more = 'more';
  let interval : NodeJS.Timeout;

  import { hoveredProject } from "../../stores/hoveredProject";

  function addDots() {
    let c = 0;
    interval = setInterval(()=> {
      if ( c < 3 ) {
        more += '.'
        c += 1
      }
    }, 500)
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
          <a href={scr_url}>sources</a>
        </li>
      </ul>
    </div>
    <div class='column'>
      {#if technologies != undefined}
        {#each technologies as technology}
            <img src={technology.img} alt={technology.alt}/>
        {/each}
      {:else}
        <p>Some technologies</p>
      {/if}
    </div>
  </div>
  <a on:mouseenter={addDots} on:mouseleave={removeDots} class="projectView" href="/projects/{id}">
    {hoveredProjectID === id ? more : more+"?"}
  </a>
</div>

<style>
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
    grid-template-columns: 3fr 2fr;
    gap: 1%;
  }
  img {
    width: 40%;
    text-align: center;
    margin: auto;
  }
  .column {
    width: 100%;
    display: grid;
    grid-template-rows: 1fr;
    margin: 1%;
  }
</style>
