<script lang="ts">
  import { RevealedPhrases, RevealedStore } from "../stores/RevealedStore";
  const handleLifeReveal = () => {
    $RevealedStore.LIFE = true;
  }
  const handleHomeInfoReveal = () => {
    $RevealedStore.HOMEINFO = true;
  }
</script>

<div>
  <h1>Hello!</h1>
  <h2>And welcome to the Artur Linnik's projects website</h2>
  <h2>My name is Artur Linnik</h2>
  <h2>But you can call me Artur</h2>
  <br>
  <h3 on:mouseenter={handleLifeReveal}>I hope that you will have a good { $RevealedStore.LIFE ? RevealedPhrases.LIFE : "..." }</h3>
  <h3 on:mouseenter={handleHomeInfoReveal}>If you click on the "projects" tab you will find { $RevealedStore.HOMEINFO ? RevealedPhrases.HOMEINFO : "...." }</h3>
</div>

<style>
  div {
    text-align: center;
  }
</style>
