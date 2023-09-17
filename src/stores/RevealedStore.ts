import { writable } from 'svelte/store';

export const RevealedStore = writable({
  LIFE: false,
  HOMEINFO: false,
  GITHUB: false,
  EMAIL: false,
  LINKEDIN: false
});

export const RevealedPhrases = {
  LIFE: "life and a special day today!",
  HOMEINFO: "some projects that Artur developed",
  GITHUB: "https://github.com/ar2rworld",
  EMAIL: "ar2r(dot)world(at)gmail(dot)com",
  LINKEDIN: "https://www.linkedin.com/in/artur-linnik-ba42b4192/"
}
