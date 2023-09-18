import { writable } from 'svelte/store';

export const hoveredHint = writable(false);

const INIT_VISIBILITY = -62;

export const hintVisibility = writable(INIT_VISIBILITY);

export const resetVisibility = () => {
  return hintVisibility.set(INIT_VISIBILITY);
}

export const hintShown = () => {
  hintVisibility.update((v) => {
    if (v - 5 >= -72) {
      return v - 5;
    }
    return -72;
  });
}
