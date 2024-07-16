<script lang="ts">
	import { onDestroy } from "svelte";

  export let character: string = ".";
  export let min: number = 2;
  export let max: number = 5;

  if (max < min) {
    console.warn("why did you enter max < min?")
  }
  
  let randomNumber = Math.round(Math.random() * min + (max - min) ); // [3..5]
  let counter = 1;
  let dots = character;
  let interval = setInterval(() => {
    counter += 1;
    dots += character;
    if ( counter > randomNumber ) {
      counter = 1;
      dots = character;
    }
  }, 1000);

  onDestroy(() => {
    clearInterval(interval);
  })

</script>

{ dots }
