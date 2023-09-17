<script lang="ts">
  import { RevealedStore, RevealedPhrases } from "../../stores/RevealedStore";
  const handleEmailReveal = () => {
    $RevealedStore.EMAIL = true;
  }
  const handleGithubReveal = () => {
    $RevealedStore.GITHUB = true;
  }
  const handleLinkedinReveal = () => {
    $RevealedStore.LINKEDIN = true;
  }
</script>
<h1>About page</h1>

<h3>I think that I have provided enough information on the <a href="/">home page</a>.</h3>
<h3 on:mouseenter={handleEmailReveal}>But if you are looking for my humanreadable email: 
  {#if $RevealedStore.EMAIL}
    <address>{RevealedPhrases.EMAIL}</address>
    <code>
      # Don't forget to replace:<br/>
      (dot) = .<br>
      (at) = @
    </code>
  {:else}
    ...
  {/if}
</h3>
<h3 on:mouseenter={handleGithubReveal}>Some other most recent project that I am working on could be found on
  {#if $RevealedStore.GITHUB}
    <a href={RevealedPhrases.GITHUB} target="_blank">github</a>
  {:else}
    ....
  {/if}
</h3>
<h3 on:mouseenter={handleLinkedinReveal}>
  But I also have
  {#if $RevealedStore.LINKEDIN}
    <a href={RevealedPhrases.LINKEDIN}>linkedin</a>
  {:else}
    linkedin
  {/if}
  (please hover over one more time)
</h3>

<style>
  code {
    display:block;
    border-radius: var(--border-radius);
    font-family: Consolas,"courier new";
    background-color: var(--bg-3);
    padding: 2px;
    font-size: 105%;
  }
</style>