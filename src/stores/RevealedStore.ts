import { writable } from 'svelte/store';

export const RevealedStore = writable({
  LIFE: false,
  HOMEINFO: false
});

export const RevealedPhrases = {
  LIFE: "life and a special day today!",
  HOMEINFO: "some projects that Artur developed"
}
