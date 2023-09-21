import { writable } from 'svelte/store';
import { hintFeedback } from './HintFeedback';

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

export const RevealedHints = writable(0);

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

    // count how many hints were revealed
    RevealedHints.update( () => {
      const revealedHints = Object.values(s).filter( v => v == true).length;

      if ( revealedHints == 1 ) {
        if (! hintFeedbackMessages[0].used ) {
          hintFeedback.push(hintFeedbackMessages[0].message);
          hintFeedbackMessages[0].used = true;
        }
      }
      if ( revealedHints == 3 ) {
        if (! hintFeedbackMessages[1].used ) {
          hintFeedback.push(hintFeedbackMessages[1].message);
          hintFeedbackMessages[1].used = true;
        }
      }
      if ( revealedHints == 5 ) {
        if (! hintFeedbackMessages[2].used ) {
          hintFeedback.push(hintFeedbackMessages[2].message);
          hintFeedbackMessages[2].used = true;
        }
      }

      return revealedHints;
    });

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

const hintFeedbackMessages = [
  {
    message: "Nice, you've got it!",
    used: false
  },
  {
    message: "Nice one",
    used: false
  },
  {
    message: "You are the master!",
    used: false
  }
];
