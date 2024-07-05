<script>
	import { onDestroy } from "svelte";
  
  let randomNumber = Math.floor(Math.random()* 2 + 3 ); // [3..5]
  let counter = 1;
  let dots = "."
  let interval = setInterval(() => {
    counter += 1;
    dots += ".";
    if ( counter > randomNumber ) {
      counter = 1;
      dots = ".";
    }
  }, 1000);

  onDestroy(() => {
    clearInterval(interval);
  })

</script>

{ dots }
