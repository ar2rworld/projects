import { writable } from 'svelte/store';

const RevealedStoreKeys = {
  LIFE     : 0,
  HOMEINFO : 1,
  GITHUB   : 2,
  EMAIL    : 3,
  LINKEDIN : 4
}

export enum RevealedStoreKeysE { 
  LIFE     = RevealedStoreKeys.LIFE,
  HOMEINFO = RevealedStoreKeys.HOMEINFO,
  GITHUB   = RevealedStoreKeys.GITHUB,
  EMAIL    = RevealedStoreKeys.EMAIL,
  LINKEDIN = RevealedStoreKeys.LINKEDIN
}

export const RevealedStore = writable({
  LIFE: false,
  HOMEINFO: false,
  GITHUB: false,
  EMAIL: false,
  LINKEDIN: false,
});

export const handleReveal = (key: RevealedStoreKeysE) => {
  RevealedStore.update(s => {
    switch (key) {
      case RevealedStoreKeysE.LIFE:
        s.LIFE = true;
        break;
      case RevealedStoreKeysE.HOMEINFO:
        s.HOMEINFO = true;
        break;
      case RevealedStoreKeysE.GITHUB:
        s.GITHUB = true;
        break;
      case RevealedStoreKeysE.EMAIL:
        s.EMAIL = true;
        break;
      case RevealedStoreKeysE.LINKEDIN:
        s.LINKEDIN = true;
        break;
    }
    return s;
  })
}

export const RevealedPhrases = {
  LIFE: "life and a special day today!",
  HOMEINFO: "some projects that Artur developed",
  GITHUB: "https://github.com/ar2rworld",
  EMAIL: "ar2r(dot)world(at)gmail(dot)com",
  LINKEDIN: "https://www.linkedin.com/in/artur-linnik-ba42b4192/"
}
