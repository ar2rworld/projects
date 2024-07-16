<script lang="ts">
  import { onDestroy } from "svelte";
	import Dots from "../../components/Dots.svelte";
	import { fade, slide } from "svelte/transition";

  let name = "";
  let email = "";
  let message = "";

  let output = "";

  const reset = () => { name = ""; email = ""; message = ""; }

  const handleSubmit = (e: SubmitEvent) => {
    e.preventDefault()
    output = "";
    console.log(name, email, message);
    fetch("/api/contact", {
      method: "POST",
      body: JSON.stringify({
        name, email, message
      }),
      headers: {
        "Access-Control-Allow-Origin": "*"
      }
    })
    .then((e:Response) => {
      console.log(e);
      if ( e.status != 200 ) {
        output = "Something went wrong("
        return;
      }
      reset();
      output = "Got it!"
    })
    .catch(() => {
      output = "Something went wrong((("
    })
  }

  let showHow = false;
  const toggle = (b: boolean) => {
    showHow = b;
  }

</script>

<div class="mx-auto text-center">
<form on:submit={handleSubmit}>
  <h1>Enter your info:</h1>
  <div class="max-w-sm mx-auto">
    <div class="my-1"><p class="text-left">Name:</p><input class="w-full bg-bg-3 dark:bg-d-bg-3" type="text" bind:value={name} required/></div>
    <div class="my-1"><p class="text-left">Email:</p><input class="w-full bg-bg-3 dark:bg-d-bg-3" type="email" bind:value={email} required/></div>
    <div class="my-1"><p class="text-left">Message:</p><textarea class="w-full bg-bg-3 dark:bg-d-bg-3 rounded-sm border-2 border-dashed border-fg-2" bind:value={message} required/></div>
    <div><input class="bg-bg-3 w-full rounded-lg border-2 border-fg-2" type="submit" value="send" /></div>
    <p class="my-2 text-sm text-right" on:mouseenter={() => toggle(true)} on:mouseleave={() => toggle(false)}>how does this form works<Dots character="?" min={2} max={3} />
      { #if showHow }
        <p class="text-right" in:slide out:fade>
          Thanks to an <a href="https://aws.amazon.com/lambda/" target="_blank">AWS Lambda function</a> written in Go, an API request is sent to <a href="https://core.telegram.org/bots/api" target="_blank">Telegram as a bot</a> to trigger the sending of a message.
        </p>
      { /if }
    </p>
  </div>
  {output ? output : ""}
</form>
</div>
