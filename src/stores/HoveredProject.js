import { writable } from 'svelte/store';

export const hoveredProject = writable("");

export const hoveredHint = writable(false);

export const hintVisibility = writable(-40);

export const hintShown = () => {
  hintVisibility.update((v) => {
    console.log(v);
    if (v -4 >= -56) {
      return v-4;
    }
    return v;
  });
}
