import { writable } from 'svelte/store';

export const hoveredProject = writable("");

export const hoveredHint = writable(false);

export const hintVisibility = writable(-62);

export const hintShown = () => {
  hintVisibility.update((v) => {
    if (v - 5 >= -72) {
      return v - 5;
    }
    return -72;
  });
}
