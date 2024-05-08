import { writable } from 'svelte/store';
import { hintFeedback } from './HintFeedback';

const RevealedStoreKeys = {
	LIFE: 0,
	HOMEINFO: 1,
	GITHUB: 2,
	EMAIL: 3,
	LINKEDIN: 4,
	HINTBADGE: 5
};

export enum RevealedStoreKeysE {
	LIFE = RevealedStoreKeys.LIFE,
	HOMEINFO = RevealedStoreKeys.HOMEINFO,
	GITHUB = RevealedStoreKeys.GITHUB,
	EMAIL = RevealedStoreKeys.EMAIL,
	LINKEDIN = RevealedStoreKeys.LINKEDIN,
	HINTBADGE = RevealedStoreKeys.HINTBADGE
}

export const RevealedStore = writable({
	LIFE: false,
	HOMEINFO: false,
	GITHUB: false,
	EMAIL: true,
	LINKEDIN: false,
	HINTBADGE: false
});

export const RevealedHints = writable(0);

export const handleReveal = (key: RevealedStoreKeysE) => {
	RevealedStore.update((s) => {
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
			case RevealedStoreKeysE.HINTBADGE:
				s.HINTBADGE = true;
				break;
		}

		// count how many hints were revealed
		RevealedHints.update(() => {
			const revealedHints = Object.values(s).filter((v) => v == true).length;

			switch (revealedHints) {
				case 1:
					if (!hintFeedbackMessages[0].used) {
						hintFeedback.push(hintFeedbackMessages[0].message);
						hintFeedbackMessages[0].used = true;
					}
					break;
				case 3:
					if (!hintFeedbackMessages[1].used) {
						hintFeedback.push(hintFeedbackMessages[1].message);
						hintFeedbackMessages[1].used = true;
					}
					break;
				case 5:
					if (!hintFeedbackMessages[2].used) {
						hintFeedback.push(hintFeedbackMessages[2].message);
						hintFeedbackMessages[2].used = true;
					}
					break;
				case 6:
					if (!hintFeedbackMessages[3].used) {
						hintFeedback.push(hintFeedbackMessages[3].message);
						hintFeedbackMessages[3].used = true;
					}
			}

			return revealedHints;
		});

		return s;
	});
};

export const RevealedPhrases = {
	LIFE: 'life and a special day today!',
	HOMEINFO: 'some projects that Artur developed',
	GITHUB: 'https://github.com/ar2rworld',
	EMAIL: 'artrlinnik@gmail.com',
	LINKEDIN: 'https://www.linkedin.com/in/artur-linnik-ba42b4192/'
};

const hintFeedbackMessages = [
	{
		message: 'Good, keep going...',
		used: false
	},
	{
		message: "Nice, you've got it!",
		used: false
	},
	{
		message: 'Nice one',
		used: false
	},
	{
		message: 'You are the master!',
		used: false
	}
];
