import { writable } from 'svelte/store';
import type { IMe } from '../types/me';

const me: IMe = {
  Username:      "",
  FirstName:     "",
  LastName:      "",
  FullName:      "",
  EmailVerified: false,
};

export const Me = writable(me);
