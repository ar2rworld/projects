<script lang="ts">
  import { fly } from 'svelte/transition';
  import { elasticOut } from 'svelte/easing';
	import { hoveredHint, hintVisibility, hintShown } from '../stores/Hint';

  const PopinDuration = 100;
  let radius = 60;
  let show = false;
  let frozen = false;

  const t = setTimeout(() => {
    if ( ! $hoveredHint ) {
      show = true;
    }
    clearTimeout(t);
  });
  
  const handleHide = () => {
    $hoveredHint = true;
    if (! frozen) {
      show = false;
      frozen = true;
      const t = setTimeout(() => {
        frozen = false;
        clearTimeout(t);
      }, PopinDuration);
    }
  }
  const handleShow = () => {
    hintShown();
    show = true;
  }

  const spin = (node: Node, { duration }: { duration: number }) => {
		return {
			duration,
			css: (t: number) => {
        const eased = elasticOut(t);

				return `
          transform: scale(${eased * 0.2 + 0.8});
					fill: hsl(
						${Math.trunc(t * 360)},
						${Math.min(50, 1000 * (1 - t))}%,
						${Math.min(50, 500 * (1 - t))}%
					);`;
			}
		};
	}
</script>

{#if show}
  <div
    in:fly={{duration: 1000, x: -300, opacity:0.1}}
    out:fly={{x: -300}}
  >
    <p on:mouseenter={handleHide}><i>Sometimes it is helpful to hover over "..."</i></p>
    <svg width="180" height="180">
      <circle class="outer" in:spin={{duration: PopinDuration}} cx="60" cy="70" r={radius} />
      <circle class="inner" cx="60" cy="70"  r="50" />
    </svg>
  </div>
{:else}
  <div class="hint" style="left:{$hintVisibility}px;">
    <button on:mouseenter={handleShow}>hint</button>
  </div>
{/if}

<style>
  svg {
    position: absolute;
    top: -35px;
    left: -10px;
    z-index: -1;
  }
  div {
    z-index: 1;
    font-size: smaller;
    position: absolute;
    width: 100px;
    top: 220px;
  }
  p {
    margin: 0px;
  }
  .inner {
    fill: var(--bg-1);
    opacity: 0.6;
  }
  .outer {
    fill: var(--bg-3);
    opacity: 0.5;
  }
  .hint {
    rotate: 90deg;
    text-align: center;
  }
  button {
    color: var(--fg-2);
    width: 50px;
    height: 15px;
    background-color: var(--bg-3);
    border: none;
  }
</style>